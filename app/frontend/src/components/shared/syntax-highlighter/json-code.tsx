'use client';
import { FC, useContext } from 'react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import darkTheme from 'react-syntax-highlighter/dist/esm/styles/prism/material-dark';
import lightTheme from 'react-syntax-highlighter/dist/esm/styles/prism/material-light';
import { AppContext } from '@/contexts';

interface IJsonCode {
  content: object | string;
}

export const JsonCode: FC<IJsonCode> = ({ content }) => {
  const {
    theme: { mode },
  } = useContext(AppContext);

  const codeString = (() => {
    if (typeof content === 'string') {
      // Try to parse and format if it's valid JSON
      try {
        const parsed = JSON.parse(content);
        return JSON.stringify(parsed, null, 2);
      } catch {
        // If parsing fails, return the string as-is
        return content;
      }
    }
    return JSON.stringify(content, null, 2);
  })();

  return (
    <SyntaxHighlighter
      language="json"
      wrapLines
      style={mode === 'dark' ? darkTheme : lightTheme}
      customStyle={{ fontSize: 12, fontWeight: 400 }}
    >
      {codeString}
    </SyntaxHighlighter>
  );
};
