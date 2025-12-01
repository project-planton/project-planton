'use client';
import {
  Box,
  CSSObject,
  Divider,
  Drawer,
  IconButton,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  styled,
  Theme,
} from '@mui/material';

export const drawerWidth = 260; // Using fixed width for drawer
export const miniDrawerWidth = 64; // Using fixed width for mini drawer

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
  flexShrink: 0,
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
  '& .MuiDrawer-paper': {
    top: theme.spacing(8),
    height: `calc(100vh - ${theme.spacing(8)})`,
    borderRight: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
    backgroundColor: theme.palette.background.paper,
    position: 'fixed',
    padding: 0,
    paddingTop: 0,
    overflowY: 'auto',
    overflowX: 'hidden',
    ...(open && {
      ...openedMixin(theme),
    }),
    ...(!open && {
      ...closedMixin(theme),
    }),
  },
}));

export const SidebarContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  flexDirection: 'column',
  height: '100%',
  position: 'relative',
  padding: 0,
  margin: theme.spacing(1, 0),
}));

export const StyledBottomSection = styled(Box, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open?: boolean }>(({ theme, open }) => ({
  width: '100%',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-end',
  alignItems: open ? 'flex-end' : 'center',
  flexGrow: 1,
  padding: theme.spacing(1),
  gap: theme.spacing(1),
  position: 'relative',
}));

export const StyledToggleIconButton = styled(IconButton)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  padding: theme.spacing(1),
  border: `${theme.spacing(0.125)} solid transparent`,
  borderRadius: theme.shape.borderRadius * 0.75,
  alignSelf: 'flex-end',
  '&:hover': {
    border: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
    backgroundColor: theme.palette.action.hover,
  },
}));

export const ContentContainer = styled(Box)(({ theme }) => ({
  flexGrow: 1,
  overflow: 'auto',
  paddingTop: theme.spacing(1),
  paddingBottom: 0,
  paddingLeft: 0,
  paddingRight: 0,
  marginTop: 0,
  marginBottom: 0,
}));

export const MenuGroupTitle = styled(Box)(({ theme }) => ({
  paddingLeft: theme.spacing(2),
  paddingRight: theme.spacing(2),
  paddingTop: theme.spacing(1.5),
  paddingBottom: theme.spacing(0.5),
  marginTop: 0,
  marginBottom: 0,
  fontSize: theme.typography.caption.fontSize,
  fontWeight: theme.typography.fontWeightBold,
  textTransform: 'uppercase',
  color: theme.palette.text.secondary,
  letterSpacing: theme.spacing(0.0625), // 0.5px / 8 = 0.0625
}));

export const StyledList = styled(List)(({ theme }) => ({
  padding: theme.spacing(0, 1),
  marginTop: 0,
  marginBottom: 0,
  display: 'flex',
  flexDirection: 'column',
  gap: 0,
  '&.MuiList-root': {
    paddingTop: 0,
    paddingBottom: 0,
  },
}));

export const StyledListItemButton = styled(ListItemButton, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open?: boolean }>(({ theme, open }) => ({
  minHeight: theme.spacing(6),
  borderRadius: theme.shape.borderRadius * 2,
  padding: theme.spacing(1, 1.5),
  margin: theme.spacing(0.5, 0),
  justifyContent: open ? 'flex-start' : 'center',
}));

export const StyledListItemIcon = styled(ListItemIcon, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open?: boolean }>(({ theme, open }) => ({
  minWidth: 0,
  justifyContent: 'center',
  marginRight: open ? theme.spacing(2) : 0,
  color: theme.palette.text.secondary,
  transition: theme.transitions.create('color', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.short,
  }),
}));

export const StyledListItemText = styled(ListItemText, {
  shouldForwardProp: (prop) => prop !== 'open',
})<{ open?: boolean }>(({ theme, open }) => ({
  opacity: open ? 1 : 0,
  transition: theme.transitions.create('opacity', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
  '& .MuiTypography-root': {
    fontSize: theme.typography.body2.fontSize,
    fontWeight: theme.typography.fontWeightMedium,
  },
}));

export const BadgeContainer = styled(Box)(({ theme }) => ({
  marginLeft: 'auto',
  minWidth: theme.spacing(2.5),
  height: theme.spacing(2.5),
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  borderRadius: theme.shape.borderRadius * 2.5,
  backgroundColor: theme.palette.primary.main,
  color: theme.palette.primary.contrastText,
  fontSize: theme.typography.caption.fontSize,
  fontWeight: theme.typography.fontWeightBold,
}));

export const StyledDivider = styled(Divider)(({ theme }) => ({
  marginTop: theme.spacing(0.5),
  marginBottom: theme.spacing(0.5),
}));
