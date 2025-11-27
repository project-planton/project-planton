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
} from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';

export const StyledAppBar = styled(AppBar)(({ theme }) => ({
  boxShadow: 'none',
  backgroundColor: theme.palette.background.paper,
  borderBottom: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  color: theme.palette.text.primary,
  position: 'fixed',
  zIndex: theme.zIndex.drawer + 1,
}));

export const StyledToolbar = styled(Toolbar)(({ theme }) => ({
  minHeight: `${theme.spacing(8)} !important`,
  height: theme.spacing(8),
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
  height: theme.spacing(0.375),
  zIndex: 1,
}));

export const LogoSection = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-start',
  height: '100%',
  gap: theme.spacing(2),
  fontWeight: theme.typography.fontWeightBold,
  fontSize: theme.typography.h5.fontSize,
  color: theme.palette.text.primary,
}));

export const StyledLogoText = styled(Typography)(({ theme }) => ({
  fontWeight: theme.typography.h6.fontWeight,
  lineHeight: 1,
  display: 'flex',
  alignItems: 'center',
  fontSize: theme.typography.h6.fontSize,
  fontFamily: theme.typography.h6.fontFamily,
  margin: 0,
  padding: 0,
}));

export const Spacer = styled(Box)({
  flexGrow: 1,
});

export const RightSection = styled(Stack)(({ theme }) => ({
  flexDirection: 'row',
  alignItems: 'center',
  gap: theme.spacing(1.5),
}));

export const SearchBox = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(1),
  padding: theme.spacing(0.75, 1.5),
  borderRadius: theme.shape.borderRadius,
  border: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.background.default,
  minWidth: theme.spacing(37.5), // 300px / 8 = 37.5
  '&:focus-within': {
    borderColor: theme.palette.primary.main,
  },
}));

export const StyledSearchIcon = styled(SearchIcon)(({ theme }) => ({
  fontSize: theme.spacing(2.5),
  color: theme.palette.text.secondary,
}));

export const StyledInputBase = styled(InputBase)(({ theme }) => ({
  flex: 1,
  fontSize: theme.typography.body2.fontSize,
  '& input': {
    padding: 0,
  },
}));

export const StyledIconButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.primary,
}));

export const StyledAvatarButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.primary,
  padding: theme.spacing(0.5),
}));

export const StyledAvatar = styled(Avatar)(({ theme }) => ({
  width: theme.spacing(4),
  height: theme.spacing(4),
  backgroundColor: theme.palette.primary.main,
}));

export const StyledMenu = styled(Menu)(({ theme }) => ({
  '& .MuiPaper-root': {
    marginTop: theme.spacing(1.5),
    minWidth: theme.spacing(25), // 200px / 8 = 25
    boxShadow: theme.shadows[3],
    backgroundColor: theme.palette.background.paper,
    color: theme.palette.text.primary,
  },
}));

export const StyledMenuItemIcon = styled('span')(({ theme }) => ({
  marginRight: theme.spacing(2),
  fontSize: theme.spacing(2.5),
  display: 'inline-flex',
  alignItems: 'center',
}));
