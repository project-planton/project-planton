'use client';
import { Box, Drawer as MuiDrawer, styled, alpha, IconButton, Typography } from '@mui/material';

export const StyledDrawer = styled(MuiDrawer)(({ theme }) => ({
  zIndex: theme.zIndex.drawer + 2,
  '& .MuiDrawer-paper': {
    backgroundColor: theme.palette.background.paper,
    border: 'none',
    borderLeft: `1px solid ${theme.palette.divider}`,
    boxShadow: theme.palette.mode === 'dark'
      ? '-20px 0 60px rgba(0, 0, 0, 0.5)'
      : '-20px 0 60px rgba(0, 0, 0, 0.1)',
    zIndex: theme.zIndex.drawer + 2,
    maxWidth: '100vw',
    
    // Default width for most drawers
    width: 480,
    
    [theme.breakpoints.down('sm')]: {
      width: '100vw',
    },
  },
  
  '& .MuiBackdrop-root': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.black, 0.7)
      : alpha(theme.palette.common.black, 0.5),
    backdropFilter: 'blur(4px)',
  },
}));

export const DrawerHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  padding: theme.spacing(2, 3),
  minHeight: 64,
  borderBottom: `1px solid ${theme.palette.divider}`,
  position: 'sticky',
  top: 0,
  backgroundColor: theme.palette.background.paper,
  zIndex: 1,
}));

export const DrawerTitle = styled(Typography)(({ theme }) => ({
  fontSize: '1.125rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
  lineHeight: 1.4,
}));

export const DrawerSubtitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.8125rem',
  color: theme.palette.text.secondary,
  marginTop: theme.spacing(0.25),
}));

export const DrawerCloseButton = styled(IconButton)(({ theme }) => ({
  padding: 8,
  borderRadius: 8,
  color: theme.palette.text.secondary,
  transition: 'all 150ms ease',
  
  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.black, 0.06),
    color: theme.palette.text.primary,
  },
}));

export const DrawerContent = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
  overflowY: 'auto',
  flex: 1,
  
  // Subtle scrollbar
  '&::-webkit-scrollbar': {
    width: 6,
  },
  '&::-webkit-scrollbar-thumb': {
    borderRadius: 6,
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.15)
      : alpha(theme.palette.common.black, 0.15),
  },
  '&::-webkit-scrollbar-track': {
    backgroundColor: 'transparent',
  },
}));

export const DrawerFooter = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  gap: theme.spacing(1.5),
  padding: theme.spacing(2, 3),
  borderTop: `1px solid ${theme.palette.divider}`,
  position: 'sticky',
  bottom: 0,
  backgroundColor: theme.palette.background.paper,
}));

export const DrawerSection = styled(Box)(({ theme }) => ({
  marginBottom: theme.spacing(3),
  
  '&:last-child': {
    marginBottom: 0,
  },
}));

export const DrawerSectionTitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.75rem',
  fontWeight: 600,
  color: theme.palette.text.secondary,
  textTransform: 'uppercase',
  letterSpacing: '0.02em',
  marginBottom: theme.spacing(1.5),
}));
