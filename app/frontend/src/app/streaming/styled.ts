'use client';
import { Box, styled } from '@mui/material';

export const StreamingContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
}));

export const LogContainer = styled(Box)(({ theme }) => ({
  maxHeight: '600px',
  overflowY: 'auto',
  backgroundColor: theme.palette.background.paper,
  border: `1px solid ${theme.palette.divider}`,
  borderRadius: theme.shape.borderRadius,
  padding: theme.spacing(1),
  fontFamily: 'monospace',
  fontSize: '0.875rem',
}));

export const LogEntry = styled(Box)<{ type: 'info' | 'error' | 'success' }>(({ theme, type }) => {
  let color = theme.palette.text.primary;
  let backgroundColor = 'transparent';

  if (type === 'error') {
    color = theme.palette.error.main;
    backgroundColor = theme.palette.error.light + '20';
  } else if (type === 'success') {
    color = theme.palette.success.main;
    backgroundColor = theme.palette.success.light + '20';
  }

  return {
    padding: theme.spacing(1),
    marginBottom: theme.spacing(0.5),
    borderRadius: theme.shape.borderRadius,
    backgroundColor,
    color,
    borderLeft: `3px solid ${color}`,
    '&:last-child': {
      marginBottom: 0,
    },
  };
});

