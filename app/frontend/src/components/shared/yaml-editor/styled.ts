'use client';
import { Box, styled } from '@mui/material';

export const EditorWrapper = styled(Box)(({ theme }) => ({
  height: '100%',
  '.ace-crimson-editor': {
    backgroundColor: `${theme.palette.crimson?.[0] || (theme.palette.mode === 'dark' ? '#2F2F2F' : '#FAFAFA')} !important`,
    color: `${theme.palette.crimson?.[20] || theme.palette.text.primary} !important`,
    border: `1px solid ${theme.palette.divider}`,
    '.ace_indent-guide': {
      background: 'none !important',
    },
    '.ace_gutter': {
      backgroundColor: `${theme.palette.crimson?.[0] || (theme.palette.mode === 'dark' ? '#2F2F2F' : '#FAFAFA')} !important`,
      color: `${theme.palette.crimson?.[20] || theme.palette.text.secondary} !important`,
    },
    '.ace_active-line, .ace_selection, .ace_gutter-active-line': {
      backgroundColor: `${theme.palette.crimson?.[60] || theme.palette.action.hover} !important`,
    },
    '.ace_meta.ace_tag': {
      color: `${theme.palette.crimson?.[10] || theme.palette.primary.main} !important`,
    },
    '.ace_keyword, .ace_paren, .ace_list, .ace_markup': {
      color: `${theme.palette.crimson?.[30] || theme.palette.info.main} !important`,
    },
    '.ace_constant , .ace_numeric': {
      color: `${theme.palette.crimson?.[40] || theme.palette.warning.main} !important`,
    },
    '.ace_string': {
      color: `${theme.palette.crimson?.[50] || theme.palette.success.main} !important`,
    },
  },
}));

