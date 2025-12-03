'use client';
import {
  CSSObject,
  Drawer,
  List,
  ListItemButton,
  ListItemButtonProps,
  ListItemIcon,
  ListItemText,
  Stack,
  styled,
  Theme,
} from '@mui/material';

export const drawerWidth = 260;
export const miniDrawerWidth = 60;
export const headerHeight = 64; // 8 * 8 = 64px

export const openedMixin = (theme: Theme): CSSObject => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
  overflowX: 'hidden',
});

export const closedMixin = (theme: Theme): CSSObject => ({
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  overflowX: 'hidden',
  width: miniDrawerWidth,
});

export const StyledDrawer = styled(Drawer, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open?: boolean }>(({ theme, open }) => ({
  width: drawerWidth,
  height: `calc(100vh - ${headerHeight}px)`,
  flexShrink: 0,
  position: 'relative',
  ...(open && {
    ...openedMixin(theme),
  }),
  ...(!open && {
    ...closedMixin(theme),
  }),
  '& .MuiDrawer-paper': {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    padding: theme.spacing(2, 1.5),
    top: headerHeight,
    height: `calc(100vh - ${headerHeight}px)`,
    borderRight: `1px solid ${theme.palette.divider}`,
    ...(open && {
      ...openedMixin(theme),
    }),
    ...(!open && {
      ...closedMixin(theme),
    }),
  },
}));

export const StyledListItemIcon = styled(ListItemIcon)(() => ({
  minWidth: 'fit-content',
}));

export const StyledListItemButton = styled(ListItemButton, {
  shouldForwardProp: (prop) => prop !== '$isMenu' && prop !== 'open',
})<ListItemButtonProps & { $isMenu?: boolean; open?: boolean }>(({ theme, open }) => {
  const activeState = {
    backgroundColor: theme.palette.grey[80] || theme.palette.action.selected,
    '& *': {
      color: `${theme.palette.text.primary} !important`,
      fill: theme.palette.text.primary,
    },
  };
  return {
    gap: theme.spacing(1),
    padding: theme.spacing(1, 1.5),
    width: open ? '100%' : 'fit-content',
    borderRadius: theme.shape.borderRadius * 2,
    '&:hover': {
      ...activeState,
    },
    '&.Mui-selected': {
      ...activeState,
      '&:hover': {
        backgroundColor: theme.palette.grey[80] || theme.palette.action.selected,
      },
    },
  };
});

export const StyledListItemText = styled(ListItemText, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open: boolean }>(({ theme, open }) => ({
  opacity: open ? 1 : 0,
  display: open ? 'block' : 'none',
  marginTop: 0,
  marginBottom: 0,
  '& .MuiTypography-root': {
    fontSize: 13,
    lineHeight: 1,
    fontWeight: 500,
    whiteSpace: 'nowrap',
    color: theme.palette.text.secondary,
  },
  transition: theme.transitions.create('all', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
}));

export const StyledList = styled(List)(({ theme }) => ({
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  gap: theme.spacing(1),
}));

export const StyledBottomSection = styled(Stack)(({ theme }) => ({
  width: '100%',
  gap: theme.spacing(1),
  justifyContent: 'flex-end',
  alignItems: 'center',
  flexGrow: 1,
  position: 'relative',
}));
