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

import * as React from 'react';

import { ChevronRight } from '@mui/icons-material';
import { Tab, Tabs } from '@mui/material';
import { Theme, useTheme, styled } from '@mui/material/styles';
import { createStyles, makeStyles } from '@mui/styles';

import {
  CodeEditor,
  EDITOR_THEME_MAP,
  LazyPanel,
  ResizableDrawer,
} from 'app/components';
import { getKeyMap } from 'app/containers/live/shortcuts';
import { EditorContext } from 'app/context/editor-context';
import { LayoutContext } from 'app/context/layout-context';
import { ScriptContext } from 'app/context/script-context';

const useStyles = makeStyles((theme: Theme) => createStyles({
  root: {
    height: '100%',
    color: theme.palette.foreground.one,
    minWidth: 0,
    overflow: 'hidden',
    width: '100%',
  },
  rootPanel: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  tabs: {
    display: 'flex',
    flexDirection: 'row',
    backgroundColor: theme.palette.background.three,
  },
  panel: {
    flex: 1,
    minHeight: 0,
  },
  editor: {
    height: '100%',
  },
  spinner: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
  },
  closer: {
    cursor: 'pointer',
    display: 'flex',
    alignItems: 'center',
  },
}), { name: 'Editor' });

const shortcutKeys = Object.values(getKeyMap()).map((keyBinding) => keyBinding.sequence);

const VisEditor = React.memo<{ visible: boolean }>(({ visible }) => {
  const classes = useStyles();
  const { script } = React.useContext(ScriptContext);
  const { setVisEditorText } = React.useContext(EditorContext);

  const editorRef = React.createRef<CodeEditor>();
  // We useEffect instead of relying on the prop because of an issue where a cursor
  // in the field causes onChange to be triggered partway through, leading to a
  // partial state being set.
  React.useEffect(() => {
    if (!editorRef.current && !script) {
      return;
    }
    editorRef.current.changeEditorValue(JSON.stringify(script.vis, undefined, 4));
    // Don't use editorRef because as a dep because the object is updated on each
    // key typed.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [script]);

  return (
    <CodeEditor
      ref={editorRef}
      onChange={setVisEditorText}
      className={classes.editor}
      visible={visible}
      spinnerClass={classes.spinner}
      shortcutKeys={shortcutKeys}
      language='json'
    />
  );
});
VisEditor.displayName = 'VisEditor';

const PxLEditor = React.memo<{ visible: boolean }>(({ visible }) => {
  const classes = useStyles();
  const { script } = React.useContext(ScriptContext);
  const { setPxlEditorText } = React.useContext(EditorContext);
  const editorRef = React.createRef<CodeEditor>();
  const theme = useTheme();

  // We useEffect instead of relying on the prop because of an issue where a cursor
  // in the field causes onChange to be triggered partway through, leading to a
  // partial state being set.
  React.useEffect(() => {
    if (!editorRef.current || !script) {
      return;
    }

    editorRef.current.changeEditorValue(script.code);
    // Don't use editorRef because as a dep because the object is updated on each
    // key typed.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [script]);

  return (
    <CodeEditor
      ref={editorRef}
      onChange={setPxlEditorText}
      className={classes.editor}
      visible={visible}
      spinnerClass={classes.spinner}
      shortcutKeys={shortcutKeys}
      language='python'
      theme={EDITOR_THEME_MAP[theme.palette.mode]}
    />
  );
});
PxLEditor.displayName = 'PxLEditor';

// eslint-disable-next-line react-memo/require-memo
const StyledTabs = styled(Tabs)(({ theme }) => ({
  flex: 1,
  '& .MuiTabs-indicator': {
    backgroundColor: theme.palette.foreground.one,
  },
}));

// eslint-disable-next-line react-memo/require-memo
const StyledTab = styled(Tab)(({ theme }) => ({
  textTransform: 'none',
  fontSize: theme.typography.subtitle2.fontSize,
  minWidth: theme.spacing(20),
  '&:focus': {
    color: theme.palette.foreground.two,
  },
}));

const LiveViewEditor = React.memo<{ visible: boolean }>(({ visible }) => {
  const classes = useStyles();
  const [tab, setTab] = React.useState('pixie');
  const { setEditorPanelOpen } = React.useContext(LayoutContext);
  const closeEditor = () => setEditorPanelOpen(false);

  return (
    <div className={classes.root}>
      <LazyPanel show={visible} className={classes.rootPanel}>
        <div className={classes.tabs}>
          <StyledTabs
            value={tab}
            // eslint-disable-next-line react-memo/require-usememo
            onChange={(event, newTab) => setTab(newTab)}
          >
            <StyledTab value='pixie' label='PxL Script' />
            <StyledTab value='vis' label='Vis Spec' />
          </StyledTabs>
          <div className={classes.closer} onClick={closeEditor}>
            <ChevronRight />
          </div>
        </div>
        <LazyPanel className={classes.panel} show={tab === 'pixie'}>
          <PxLEditor visible={visible && tab === 'pixie'} />
        </LazyPanel>
        <LazyPanel className={classes.panel} show={tab === 'vis'}>
          <VisEditor visible={visible && tab === 'vis'} />
        </LazyPanel>
      </LazyPanel>
    </div>
  );
});
LiveViewEditor.displayName = 'LiveViewEditor';

export const EditorSplitPanel: React.FC = React.memo(({ children }) => {
  const { editorPanelOpen } = React.useContext(LayoutContext);

  return (
    <ResizableDrawer
      drawerDirection='right'
      initialSize={850}
      minSize={25}
      open={editorPanelOpen}
      otherContent={children}
      overlay
    >
      <LiveViewEditor visible={editorPanelOpen} />
    </ResizableDrawer>
  );
});
EditorSplitPanel.displayName = 'EditorSplitPanel';
