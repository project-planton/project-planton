'use client';
import {
  AppBar,
  Avatar,
  Box,
  IconButton,
  InputBase,
  LinearProgress,
  Menu,
  Stack,
  styled,
  Toolbar,
  Typography,
  alpha,
} from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';

export const StyledAppBar = styled(AppBar)(({ theme }) => ({
  boxShadow: 'none',
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.background.default, 0.8)
    : alpha(theme.palette.background.default, 0.9),
  backdropFilter: 'blur(12px)',
  WebkitBackdropFilter: 'blur(12px)',
  borderBottom: `1px solid ${theme.palette.divider}`,
  color: theme.palette.text.primary,
  position: 'fixed',
  zIndex: theme.zIndex.drawer + 1,
  minHeight: 'unset',
}));

export const StyledToolbar = styled(Toolbar)(({ theme }) => ({
  minHeight: '56px !important',
  height: 56,
  padding: theme.spacing(0, 3),
  gap: theme.spacing(2),
  justifyContent: 'space-between',
  alignItems: 'center',
  display: 'flex',
}));

export const StyledLinearProgress = styled(LinearProgress)(({ theme }) => ({
  position: 'absolute',
  top: 0,
  left: 0,
  right: 0,
  height: 2,
  zIndex: 1,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.primary.main, 0.2)
    : alpha(theme.palette.primary.main, 0.15),
  '& .MuiLinearProgress-bar': {
    backgroundColor: theme.palette.primary.main,
  },
}));

export const LogoSection = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-start',
  height: '100%',
  gap: theme.spacing(1.5),
}));

export const StyledLogoText = styled(Typography)(({ theme }) => ({
  fontWeight: 600,
  lineHeight: 1,
  display: 'flex',
  alignItems: 'center',
  fontSize: '1rem',
  fontFamily: theme.typography.fontFamily,
  margin: 0,
  padding: 0,
  color: theme.palette.text.primary,
  letterSpacing: '-0.01em',
}));

export const Spacer = styled(Box)({
  flexGrow: 1,
});

export const RightSection = styled(Stack)(({ theme }) => ({
  flexDirection: 'row',
  alignItems: 'center',
  gap: theme.spacing(1),
}));

export const SearchBox = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(1),
  padding: theme.spacing(0.75, 1.5),
  borderRadius: 8,
  border: `1px solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.03)
    : alpha(theme.palette.common.black, 0.02),
  minWidth: 240,
  transition: 'all 150ms ease',
  
  '&:hover': {
    borderColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.15)
      : alpha(theme.palette.common.black, 0.15),
  },
  
  '&:focus-within': {
    borderColor: theme.palette.primary.main,
    boxShadow: `0 0 0 2px ${alpha(theme.palette.primary.main, 0.2)}`,
    backgroundColor: theme.palette.background.paper,
  },
}));

export const SearchShortcut = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  padding: theme.spacing(0.25, 0.75),
  borderRadius: 4,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.08)
    : alpha(theme.palette.common.black, 0.06),
  fontSize: '0.6875rem', // 11px
  fontWeight: 500,
  color: theme.palette.text.secondary,
  fontFamily: 'inherit',
  lineHeight: 1.4,
}));

export const StyledSearchIcon = styled(SearchIcon)(({ theme }) => ({
  fontSize: 18,
  color: theme.palette.text.secondary,
}));

export const StyledInputBase = styled(InputBase)(({ theme }) => ({
  minHeight: 'unset',
  flex: 1,
  fontSize: '0.8125rem', // 13px
  backgroundColor: 'inherit',
  '& input': {
    padding: 0,
    '&::placeholder': {
      color: theme.palette.text.secondary,
      opacity: 0.8,
    },
  },
}));

export const StyledIconButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.secondary,
  padding: 8,
  borderRadius: 8,
  transition: 'all 150ms ease',
  
  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.05)
      : alpha(theme.palette.common.black, 0.04),
    color: theme.palette.text.primary,
  },
  
  '&:focus-visible': {
    outline: 'none',
    boxShadow: `0 0 0 2px ${alpha(theme.palette.primary.main, 0.5)}`,
  },
}));

export const StyledAvatarButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.primary,
  padding: 4,
  borderRadius: 8,
  
  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.05)
      : alpha(theme.palette.common.black, 0.04),
  },
}));

export const StyledAvatar = styled(Avatar)(({ theme }) => ({
  width: 32,
  height: 32,
  backgroundColor: theme.palette.primary.main,
  fontSize: '0.8125rem',
  fontWeight: 600,
}));

export const StyledMenu = styled(Menu)(({ theme }) => ({
  '& .MuiPaper-root': {
    marginTop: theme.spacing(1),
    minWidth: 200,
    borderRadius: 12,
    border: `1px solid ${theme.palette.divider}`,
    boxShadow: theme.palette.mode === 'dark'
      ? '0 10px 40px rgba(0, 0, 0, 0.4)'
      : '0 10px 40px rgba(0, 0, 0, 0.1)',
    backgroundColor: theme.palette.background.paper,
    color: theme.palette.text.primary,
    padding: theme.spacing(0.5),
    
    '& .MuiMenuItem-root': {
      borderRadius: 8,
      padding: theme.spacing(1, 1.5),
      fontSize: '0.8125rem',
      gap: theme.spacing(1.5),
      transition: 'background-color 150ms ease',
      
      '&:hover': {
        backgroundColor: theme.palette.mode === 'dark'
          ? alpha(theme.palette.common.white, 0.05)
          : alpha(theme.palette.common.black, 0.04),
      },
    },
  },
}));

export const StyledMenuItemIcon = styled('span')(({ theme }) => ({
  display: 'inline-flex',
  alignItems: 'center',
  color: theme.palette.text.secondary,
  fontSize: 18,
}));

// Divider styled for menu
export const MenuDivider = styled('div')(({ theme }) => ({
  height: 1,
  backgroundColor: theme.palette.divider,
  margin: theme.spacing(0.5, 0),
}));
