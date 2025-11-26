'use client';
import { Box, Container, styled } from '@mui/material';

export const StyledWrapperBox = styled(Box)(({ theme }) => ({
  display: 'flex',
  flex: 1,
  minHeight: `calc(100vh - ${theme.spacing(8)})`,
  marginTop: theme.spacing(8),
}));

export const StyledContainer = styled(Container)(({ theme }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  width: '100%',
  maxWidth: '100% !important',
  backgroundColor: theme.palette.background.default,
  minHeight: `calc(100vh - ${theme.spacing(8)})`,
}));

