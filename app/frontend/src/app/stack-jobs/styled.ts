'use client';
import { Box, Paper, Stack, styled } from '@mui/material';

export const StackJobContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
}));

export const FlexCenterRow = styled(Stack)(() => ({
  flexDirection: 'row',
  alignItems: 'center',
}));

export const HeaderContainer = styled(Paper)(({ theme }) => ({
  borderRadius: theme.shape.borderRadius * 2,
  border: `1px solid ${theme.palette.divider}`,
  padding: theme.spacing(1.5),
  boxShadow: 'none',
}));

export const TopSection = styled(FlexCenterRow)(({ theme }) => ({
  justifyContent: 'space-between',
  gap: theme.spacing(1),
  flexWrap: 'wrap',
}));

export const BottomSection = styled(FlexCenterRow)(({ theme }) => ({
  justifyContent: 'space-between',
  gap: theme.spacing(1),
  marginTop: theme.spacing(1.5),
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

export const LogEntry = styled(Box, {
  shouldForwardProp: (prop) => prop !== 'streamType',
})<{ streamType: 'stdout' | 'stderr' }>(({ theme, streamType }) => {
  let color = theme.palette.text.primary;
  let backgroundColor = 'transparent';

  if (streamType === 'stderr') {
    color = theme.palette.error.main;
    backgroundColor = theme.palette.error.light + '20';
  }

  return {
    padding: theme.spacing(0.5),
    marginBottom: theme.spacing(0.25),
    borderRadius: theme.shape.borderRadius,
    backgroundColor,
    color,
    borderLeft: streamType === 'stderr' ? `3px solid ${color}` : 'none',
    whiteSpace: 'pre-wrap',
    wordBreak: 'break-word',
    '&:last-child': {
      marginBottom: 0,
    },
  };
});
