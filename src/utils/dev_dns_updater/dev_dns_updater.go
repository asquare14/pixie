/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/txn2/txeh"
	"golang.org/x/sync/errgroup"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var dnsEntriesByServiceCfg = map[string][]string{
	"cloud-proxy-service": {"", "work", "segment", "docs"},
	"vzconn-service":      {"cloud"},
}

var dnsEntriesByService = map[string][]string{}

type svcInfo struct {
	SvcName string
	Addr    string
}

func setupFlags() {
	pflag.String("n", "plc-dev", "The namespace to watch (plc-dev) by default")
	pflag.String("domain-name", "dev.withpixie.dev", "The domain name to use")
	pflag.String("kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
}

func parseFlags() {
	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PL")
	viper.BindPFlags(pflag.CommandLine)
}

// getConfig gets the kubernetes rest config.
func getConfig() *rest.Config {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", viper.GetString("kubeconfig"))
	if err != nil {
		log.WithError(err).Fatal("Could not build kubeconfig")
	}

	return config
}

// getClientset gets the clientset for the current kubernetes cluster.
func getClientset(config *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.WithError(err).Fatal("Could not create k8s clientset")
	}

	return clientset
}

func homeDir() string {
	if u := os.Getenv("SUDO_USER"); u != "" {
		return fmt.Sprintf("/home/%s", u)
	}
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func generateDomainEntries() {
	for svcName, domainPrefixes := range dnsEntriesByServiceCfg {
		entries := []string{}

		for _, domainPrefix := range domainPrefixes {
			entries = append(entries, getDomainEntry(domainPrefix))
		}
		dnsEntriesByService[svcName] = entries

		log.WithField("service", svcName).
			WithField("entries", strings.Join(entries, ", ")).
			Info("DNS Entries")
	}
}

func getDomainEntry(prefix string) string {
	domainSuffix := viper.GetString("domain-name")
	if len(prefix) == 0 {
		return domainSuffix
	}
	return fmt.Sprintf("%s.%s", prefix, domainSuffix)
}

func watchForExternalIP(ch <-chan watch.Event, outCh chan<- svcInfo) error {
	for u := range ch {
		svc := u.Object.(*v1.Service)
		svcName := svc.ObjectMeta.Name

		log.WithField("service", svcName).
			Debug("Service")

		if _, ok := dnsEntriesByService[svcName]; !ok {
			continue
		}

		ing := svc.Status.LoadBalancer.Ingress

		if len(ing) > 0 {
			if ing[0].IP != "" {
				outCh <- svcInfo{
					SvcName: svc.ObjectMeta.Name,
					Addr:    ing[0].IP,
				}
			} else {
				if ing[0].Hostname != "" {
					log.WithField("ing[0].Hostname", ing[0].Hostname).
						Debug("Using Hostname")

					ip, _ := net.LookupIP(ing[0].Hostname)

					outCh <- svcInfo{
						SvcName: svc.ObjectMeta.Name,
						Addr:    ip[0].String(),
					}
				}
			}
		}
	}
	return nil
}

func k8sWatchAndUpdateHosts() error {
	kubeConfig := getConfig()
	clientset := getClientset(kubeConfig)
	namespace := viper.GetString("n")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceWatcher, err := clientset.CoreV1().Services(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.WithError(err).Fatal("failed to watch cloud proxy")
	}
	defer serviceWatcher.Stop()

	svcInfoCh := make(chan svcInfo)
	var g errgroup.Group
	g.Go(func() error {
		return watchForExternalIP(serviceWatcher.ResultChan(), svcInfoCh)
	})

	g.Go(func() error {
		return updateHostsFile(svcInfoCh)
	})

	return g.Wait()
}

func updateHostsFile(svcInfoCh <-chan svcInfo) error {
	for s := range svcInfoCh {
		log.WithField("service", s.SvcName).
			WithField("addr", s.Addr).
			Info("Update")
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			return err
		}

		if entries, ok := dnsEntriesByService[s.SvcName]; ok {
			hosts.RemoveHosts(entries)
			hosts.AddHosts(s.Addr, entries)
		}
		err = hosts.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanup() {
	log.Info("Cleaning up hosts file")
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		log.WithError(err).Fatal("Failed to get hosts file")
	}

	for _, dnsEntries := range dnsEntriesByService {
		hosts.RemoveHosts(dnsEntries)
	}
	err = hosts.Save()
	if err != nil {
		log.WithError(err).Fatal("Failed to save hosts file")
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func sudoSelfIfNotRoot() {
	uid := os.Getuid()
	if uid != 0 {
		f, _ := filepath.Abs(os.Args[0])
		args := append([]string{f}, os.Args[1:]...)

		c1 := exec.Command("sudo", args...)
		c1.Stdin = os.Stdin
		c1.Stdout = os.Stdout
		c1.Stderr = os.Stderr

		if err := c1.Run(); err != nil {
			panic(err)
		}

		err := c1.Wait()
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}
}
func main() {
	sudoSelfIfNotRoot()

	setupFlags()
	parseFlags()

	generateDomainEntries()
	err := copyFile("/etc/hosts", "/etc/hosts.bak")
	if err != nil {
		log.WithError(err).Fatal("Failed to backup /etc/hosts")
	}
	defer cleanup()

	// Also run cleanup on ctrl+c.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	var g errgroup.Group
	g.Go(k8sWatchAndUpdateHosts)

	// TODO(zasgar): Add Minikube tunnel

	err = g.Wait()
	if err != nil {
		log.WithError(err).Fatal("Error auto updating entries")
	}
}
