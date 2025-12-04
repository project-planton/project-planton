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
