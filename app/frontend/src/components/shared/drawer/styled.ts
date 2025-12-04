'use client';
import { Box, Drawer as MuiDrawer, styled } from '@mui/material';

export const StyledDrawer = styled(MuiDrawer)(({ theme }) => ({
  zIndex: theme.zIndex.drawer + 2,
  '& .MuiDrawer-paper': {
    boxShadow: theme.shadows[16],
    zIndex: theme.zIndex.drawer + 2,
  },
}));

export const DrawerHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  padding: theme.spacing(2),
  minHeight: theme.spacing(8),
  borderBottom: `1px solid ${theme.palette.divider}`,
}));

export const DrawerContent = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2),
  overflowY: 'auto',
  height: '100%',
}));
