// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/services/auth/proto/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	_ "pixielabs.ai/pixielabs/src/services/common/proto"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (m *LoginRequest) Reset()      { *m = LoginRequest{} }
func (*LoginRequest) ProtoMessage() {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d888daca5d8729c, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return m.Size()
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type LoginReply struct {
	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresAt int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (m *LoginReply) Reset()      { *m = LoginReply{} }
func (*LoginReply) ProtoMessage() {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d888daca5d8729c, []int{1}
}
func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return m.Size()
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReply) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

type GetAugmentedAuthTokenRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *GetAugmentedAuthTokenRequest) Reset()      { *m = GetAugmentedAuthTokenRequest{} }
func (*GetAugmentedAuthTokenRequest) ProtoMessage() {}
func (*GetAugmentedAuthTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d888daca5d8729c, []int{2}
}
func (m *GetAugmentedAuthTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAugmentedAuthTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAugmentedAuthTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAugmentedAuthTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAugmentedAuthTokenRequest.Merge(m, src)
}
func (m *GetAugmentedAuthTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAugmentedAuthTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAugmentedAuthTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAugmentedAuthTokenRequest proto.InternalMessageInfo

func (m *GetAugmentedAuthTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetAugmentedAuthTokenResponse struct {
	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresAt int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (m *GetAugmentedAuthTokenResponse) Reset()      { *m = GetAugmentedAuthTokenResponse{} }
func (*GetAugmentedAuthTokenResponse) ProtoMessage() {}
func (*GetAugmentedAuthTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d888daca5d8729c, []int{3}
}
func (m *GetAugmentedAuthTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAugmentedAuthTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAugmentedAuthTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAugmentedAuthTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAugmentedAuthTokenResponse.Merge(m, src)
}
func (m *GetAugmentedAuthTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAugmentedAuthTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAugmentedAuthTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAugmentedAuthTokenResponse proto.InternalMessageInfo

func (m *GetAugmentedAuthTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetAugmentedAuthTokenResponse) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pl.services.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pl.services.LoginReply")
	proto.RegisterType((*GetAugmentedAuthTokenRequest)(nil), "pl.services.GetAugmentedAuthTokenRequest")
	proto.RegisterType((*GetAugmentedAuthTokenResponse)(nil), "pl.services.GetAugmentedAuthTokenResponse")
}

func init() { proto.RegisterFile("src/services/auth/proto/auth.proto", fileDescriptor_2d888daca5d8729c) }

var fileDescriptor_2d888daca5d8729c = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0xb4, 0x41, 0x45, 0xaa, 0xdb, 0x05, 0x0b, 0x89, 0x12, 0x51, 0xab, 0x84, 0xa5, 0x20, 0x91,
	0x08, 0xca, 0x86, 0x18, 0xc2, 0xc2, 0xc2, 0x54, 0x3a, 0xb1, 0x54, 0x49, 0x30, 0xa9, 0x21, 0x89,
	0x43, 0xec, 0x40, 0xbb, 0xf1, 0x09, 0x7c, 0x06, 0xdf, 0xc0, 0x17, 0x30, 0x76, 0xec, 0x48, 0xd3,
	0x85, 0xb1, 0x9f, 0x80, 0x12, 0xb7, 0x28, 0x95, 0x5a, 0x84, 0xd8, 0xee, 0xbd, 0xdc, 0x5d, 0xee,
	0x9d, 0x8c, 0x74, 0x11, 0xbb, 0xa6, 0xa0, 0xf1, 0x13, 0x73, 0xa9, 0x30, 0xed, 0x44, 0xf6, 0xcc,
	0x28, 0xe6, 0x92, 0xe7, 0xd0, 0xc8, 0x21, 0xae, 0x44, 0xbe, 0x31, 0xa7, 0x68, 0x47, 0x1e, 0x93,
	0xbd, 0xc4, 0x31, 0x5c, 0x1e, 0x98, 0x1e, 0xf7, 0xb8, 0xa2, 0x3b, 0xc9, 0x5d, 0x3e, 0x29, 0x6d,
	0x86, 0x94, 0x56, 0xdb, 0x5f, 0xf0, 0x77, 0x79, 0x10, 0xf0, 0x70, 0xf6, 0x87, 0xfb, 0x67, 0xa9,
	0x48, 0xfa, 0x31, 0xaa, 0x5e, 0x71, 0x8f, 0x85, 0x6d, 0xfa, 0x98, 0x50, 0x21, 0xf1, 0x1e, 0xaa,
	0xda, 0xae, 0x4b, 0x85, 0xe8, 0x4a, 0xfe, 0x40, 0xc3, 0x1a, 0x6c, 0xc0, 0x66, 0xb9, 0x5d, 0x51,
	0xbb, 0x4e, 0xb6, 0xd2, 0x2d, 0x84, 0x66, 0x92, 0xc8, 0x1f, 0xe0, 0x2d, 0x54, 0x2a, 0x32, 0xd5,
	0x80, 0xeb, 0x08, 0xd1, 0x7e, 0xc4, 0x62, 0x2a, 0xba, 0xb6, 0xac, 0xad, 0x35, 0x60, 0x73, 0xbd,
	0x5d, 0x9e, 0x6d, 0x2c, 0xa9, 0x9f, 0xa2, 0xdd, 0x4b, 0x2a, 0xad, 0xc4, 0x0b, 0x68, 0x28, 0xe9,
	0xad, 0x95, 0xc8, 0x5e, 0xee, 0x3d, 0x4f, 0xb1, 0xd4, 0x54, 0xef, 0xa0, 0xfa, 0x0a, 0x95, 0x88,
	0x78, 0x28, 0xe8, 0xbf, 0xb2, 0x9c, 0xbc, 0x43, 0x54, 0xc9, 0xac, 0xae, 0x55, 0x53, 0xf8, 0x1c,
	0x95, 0xf2, 0xf3, 0xf0, 0x8e, 0x51, 0x28, 0xdf, 0x28, 0xb6, 0xa4, 0x6d, 0x2f, 0xfb, 0x14, 0xf9,
	0x03, 0x1d, 0x60, 0x1f, 0x6d, 0x16, 0x43, 0xe6, 0x01, 0xf1, 0xc1, 0x02, 0xff, 0xb7, 0xd3, 0xb5,
	0xc3, 0xbf, 0x50, 0xd5, 0xbd, 0x3a, 0xb8, 0x60, 0xc3, 0x31, 0x01, 0xa3, 0x31, 0x01, 0xd3, 0x31,
	0x81, 0x2f, 0x29, 0x81, 0x6f, 0x29, 0x81, 0x1f, 0x29, 0x81, 0xc3, 0x94, 0xc0, 0xcf, 0x94, 0xc0,
	0xaf, 0x94, 0x80, 0x69, 0x4a, 0xe0, 0xeb, 0x84, 0x80, 0xe1, 0x84, 0x80, 0xd1, 0x84, 0x80, 0x9b,
	0x56, 0xc4, 0xfa, 0x8c, 0xfa, 0xb6, 0x23, 0x0c, 0x9b, 0x99, 0x3f, 0x83, 0xb9, 0xe2, 0x41, 0x9e,
	0x65, 0xd0, 0xd9, 0xc8, 0x71, 0xeb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x6c, 0x34, 0x9e, 0xb7,
	0x02, 0x00, 0x00,
}

func (this *LoginRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoginRequest)
	if !ok {
		that2, ok := that.(LoginRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AccessToken != that1.AccessToken {
		return false
	}
	return true
}
func (this *LoginReply) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoginReply)
	if !ok {
		that2, ok := that.(LoginReply)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if this.ExpiresAt != that1.ExpiresAt {
		return false
	}
	return true
}
func (this *GetAugmentedAuthTokenRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetAugmentedAuthTokenRequest)
	if !ok {
		that2, ok := that.(GetAugmentedAuthTokenRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	return true
}
func (this *GetAugmentedAuthTokenResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetAugmentedAuthTokenResponse)
	if !ok {
		that2, ok := that.(GetAugmentedAuthTokenResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if this.ExpiresAt != that1.ExpiresAt {
		return false
	}
	return true
}
func (this *LoginRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&auth.LoginRequest{")
	s = append(s, "AccessToken: "+fmt.Sprintf("%#v", this.AccessToken)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LoginReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.LoginReply{")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "ExpiresAt: "+fmt.Sprintf("%#v", this.ExpiresAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetAugmentedAuthTokenRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&auth.GetAugmentedAuthTokenRequest{")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetAugmentedAuthTokenResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.GetAugmentedAuthTokenResponse{")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "ExpiresAt: "+fmt.Sprintf("%#v", this.ExpiresAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAuth(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetAugmentedToken(ctx context.Context, in *GetAugmentedAuthTokenRequest, opts ...grpc.CallOption) (*GetAugmentedAuthTokenResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pl.services.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAugmentedToken(ctx context.Context, in *GetAugmentedAuthTokenRequest, opts ...grpc.CallOption) (*GetAugmentedAuthTokenResponse, error) {
	out := new(GetAugmentedAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/pl.services.AuthService/GetAugmentedToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetAugmentedToken(context.Context, *GetAugmentedAuthTokenRequest) (*GetAugmentedAuthTokenResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.services.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAugmentedToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAugmentedAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAugmentedToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.services.AuthService/GetAugmentedToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAugmentedToken(ctx, req.(*GetAugmentedAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.services.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "GetAugmentedToken",
			Handler:    _AuthService_GetAugmentedToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/services/auth/proto/auth.proto",
}

func (m *LoginRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	return i, nil
}

func (m *LoginReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.ExpiresAt != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.ExpiresAt))
	}
	return i, nil
}

func (m *GetAugmentedAuthTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAugmentedAuthTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *GetAugmentedAuthTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAugmentedAuthTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.ExpiresAt != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.ExpiresAt))
	}
	return i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func (m *LoginReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAuth(uint64(m.ExpiresAt))
	}
	return n
}

func (m *GetAugmentedAuthTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func (m *GetAugmentedAuthTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAuth(uint64(m.ExpiresAt))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LoginRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LoginRequest{`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LoginReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LoginReply{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`ExpiresAt:` + fmt.Sprintf("%v", this.ExpiresAt) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetAugmentedAuthTokenRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetAugmentedAuthTokenRequest{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetAugmentedAuthTokenResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetAugmentedAuthTokenResponse{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`ExpiresAt:` + fmt.Sprintf("%v", this.ExpiresAt) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAuth(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LoginRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAugmentedAuthTokenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAugmentedAuthTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAugmentedAuthTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAugmentedAuthTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAugmentedAuthTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAugmentedAuthTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAuth
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)
