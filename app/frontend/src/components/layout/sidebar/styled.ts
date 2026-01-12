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
  alpha,
} from '@mui/material';

// Layout dimensions
export const drawerWidth = 240;
export const miniDrawerWidth = 64;
export const headerHeight = 56;

export const openedMixin = (theme: Theme): CSSObject => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.easeOut,
    duration: '200ms',
  }),
  overflowX: 'hidden',
});

export const closedMixin = (theme: Theme): CSSObject => ({
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.easeIn,
    duration: '150ms',
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
    alignItems: 'flex-start',
    padding: theme.spacing(2, 1.5),
    top: headerHeight,
    height: `calc(100vh - ${headerHeight}px)`,
    border: 'none',
    borderRight: `1px solid ${theme.palette.divider}`,
    backgroundColor: 'transparent',
    backgroundImage: 'none',
    ...(open && {
      ...openedMixin(theme),
    }),
    ...(!open && {
      ...closedMixin(theme),
    }),
  },
}));

export const StyledListItemIcon = styled(ListItemIcon)(({ theme }) => ({
  minWidth: 'unset',
  width: 20,
  height: 20,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: theme.palette.text.secondary,
  transition: 'color 200ms cubic-bezier(0.4, 0, 0.2, 1), transform 200ms cubic-bezier(0.4, 0, 0.2, 1)',
}));

export const StyledListItemButton = styled(ListItemButton, {
  shouldForwardProp: (prop) => prop !== '$isMenu' && prop !== 'open',
})<ListItemButtonProps & { $isMenu?: boolean; open?: boolean }>(({ theme, open }) => ({
  gap: theme.spacing(1.5),
  padding: theme.spacing(1, 1.5),
  width: open ? '100%' : 'fit-content',
  minHeight: 40,
  borderRadius: 8,
  position: 'relative',
  transition: 'all 200ms cubic-bezier(0.4, 0, 0.2, 1)',

  // Default state
  backgroundColor: 'transparent',
  borderLeft: '3px solid transparent',
  marginLeft: -3,
  paddingLeft: `calc(${theme.spacing(1.5)} + 3px)`,

  // Hover state
  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.05)
      : alpha(theme.palette.common.black, 0.04),
    transform: 'translateX(2px)',
    '& .MuiListItemIcon-root': {
      color: theme.palette.primary.main,
      transform: 'scale(1.1)',
    },
    '& .MuiListItemText-root .MuiTypography-root': {
      color: theme.palette.text.primary,
    },
  },

  // Selected/Active state
  '&.Mui-selected': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.primary.main, 0.15)
      : alpha(theme.palette.primary.main, 0.1),
    borderLeftColor: theme.palette.primary.main,
    '& .MuiListItemIcon-root': {
      color: theme.palette.primary.main,
    },
    '& .MuiListItemText-root .MuiTypography-root': {
      color: theme.palette.primary.main,
      fontWeight: 600,
    },
    '&:hover': {
      backgroundColor: theme.palette.mode === 'dark'
        ? alpha(theme.palette.primary.main, 0.2)
        : alpha(theme.palette.primary.main, 0.15),
      transform: 'none',
      '& .MuiListItemIcon-root': {
        transform: 'none',
      },
    },
  },
}));

export const StyledListItemText = styled(ListItemText, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open: boolean }>(({ theme, open }) => ({
  opacity: open ? 1 : 0,
  display: open ? 'block' : 'none',
  marginTop: 0,
  marginBottom: 0,
  '& .MuiTypography-root': {
    fontSize: '0.8125rem', // 13px
    lineHeight: 1.4,
    fontWeight: 500,
    whiteSpace: 'nowrap',
    color: theme.palette.text.secondary,
    transition: 'color 150ms ease, font-weight 150ms ease',
  },
  transition: 'opacity 200ms ease',
}));

export const StyledList = styled(List)(({ theme }) => ({
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'stretch',
  gap: theme.spacing(0.5),
  padding: 0,
}));

export const StyledSectionLabel = styled('div', {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open: boolean }>(({ theme, open }) => ({
  fontSize: '0.6875rem', // 11px
  fontWeight: 600,
  textTransform: 'uppercase',
  letterSpacing: '0.05em',
  color: theme.palette.text.disabled,
  padding: theme.spacing(1.5, 1.5, 0.5, 1.5),
  marginTop: theme.spacing(1),
  opacity: open ? 1 : 0,
  height: open ? 'auto' : 0,
  overflow: 'hidden',
  transition: 'opacity 200ms ease, height 150ms ease',
  whiteSpace: 'nowrap',
}));

export const StyledBottomSection = styled(Stack)(({ theme }) => ({
  width: '100%',
  gap: theme.spacing(1),
  justifyContent: 'flex-end',
  alignItems: 'center',
  flexGrow: 1,
  position: 'relative',
  paddingTop: theme.spacing(2),
  marginTop: 'auto',

  // Subtle divider at top
  '& > hr': {
    width: '100%',
    borderColor: theme.palette.divider,
    margin: 0,
  },
}));

// Toggle button for sidebar
export const StyledToggleButton = styled('button')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 32,
  height: 32,
  padding: 0,
  border: `1px solid transparent`,
  borderRadius: 6,
  backgroundColor: 'transparent',
  color: theme.palette.text.secondary,
  cursor: 'pointer',
  transition: 'all 150ms ease',

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.05)
      : alpha(theme.palette.common.black, 0.04),
    borderColor: theme.palette.divider,
    color: theme.palette.text.primary,
  },

  '&:focus-visible': {
    outline: 'none',
    boxShadow: `0 0 0 2px ${alpha(theme.palette.primary.main, 0.5)}`,
  },
}));
