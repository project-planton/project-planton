'use client';
import { FC, CSSProperties } from 'react';
import AceEditor, { IAceOptions } from 'react-ace';
import 'ace-builds/src-noconflict/theme-crimson_editor';
import 'ace-builds/src-noconflict/mode-yaml';
import 'ace-builds/src-noconflict/worker-yaml';
import 'ace-builds/src-noconflict/snippets/yaml';
import { EditorWrapper } from '@/components/shared/yaml-editor/styled';

interface YamlEditorProps {
  value?: string;
  onChange?: (value: string) => void;
  readOnly?: boolean;
  style?: CSSProperties;
  height?: string;
  aceOptions?: IAceOptions;
}

export const YamlEditor: FC<YamlEditorProps> = ({
  value = '',
  onChange,
  readOnly = false,
  style,
  height = '400px',
  aceOptions,
}) => {
  return (
    <EditorWrapper>
      <AceEditor
        mode="yaml"
        theme="crimson_editor"
        name="yaml_editor"
        onChange={onChange}
        value={value}
        readOnly={readOnly}
        editorProps={{ $blockScrolling: true }}
        setOptions={{
          enableBasicAutocompletion: true,
          enableLiveAutocompletion: true,
          enableSnippets: true,
          showLineNumbers: true,
          showGutter: true,
          showPrintMargin: false,
          tabSize: 2,
          useWorker: true,
          ...aceOptions,
        }}
        onLoad={(editor) => {
          editor.renderer.setPadding(15);
          editor.renderer.setScrollMargin(15, 15);
        }}
        style={{
          width: '100%',
          height: height,
          borderRadius: '4px',
          ...style,
        }}
      />
    </EditorWrapper>
  );
};
