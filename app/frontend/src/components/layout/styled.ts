'use client';
import { Box, Container, styled } from '@mui/material';

export const StyledWrapperBox = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$hasWhiteBg',
})<{ $hasWhiteBg?: boolean }>(({ theme, $hasWhiteBg }) => ({
  display: 'flex',
  flex: 1,
  minHeight: `calc(100vh - ${theme.spacing(8)})`,
  backgroundColor: $hasWhiteBg ? theme.palette.grey[0] : theme.palette.grey[20],
  marginTop: theme.spacing(8),
}));

export const StyledContainer = styled(Container)(({ theme }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  width: '100%',
  maxWidth: '100% !important',
  minHeight: `calc(100vh - ${theme.spacing(8)})`,
}));
