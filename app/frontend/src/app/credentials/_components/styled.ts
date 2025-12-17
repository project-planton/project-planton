import {
  Box,
  BoxProps,
  Button,
  ButtonProps,
  Card,
  CardContent,
  CardHeader,
  styled,
  TextField,
  TextFieldProps,
  Tabs,
  TabsProps,
  Typography,
} from '@mui/material';

export const StyledCard = styled(Card)(({ theme }) => ({
  padding: theme.spacing(1.5),
  border: `1px solid ${theme.palette.divider}`,
  borderRadius: theme.shape.borderRadius * 3,
  boxShadow: '1px 1px 6px 0px rgba(0, 0, 0, 0.06)',
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  cursor: 'pointer',
  '&:hover': {
    backgroundColor: theme.palette.grey[20],
  },
}));

export const StyledCardHeader = styled(CardHeader)(({ theme }) => ({
  padding: 0,
  '& .MuiCardHeader-content': {
    display: 'flex',
    flexDirection: 'column',
    gap: theme.spacing(0.75),
  },
  '& .MuiCardHeader-title': {
    fontSize: theme.spacing(1.5),
    fontWeight: 500,
  },
  '& .MuiCardHeader-subheader': {
    fontSize: theme.spacing(1.5),
    fontWeight: 400,
    color: theme.palette.text.secondary,
    display: '-webkit-box',
    '-webkit-line-clamp': '2',
    '-webkit-box-orient': 'vertical',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    lineHeight: 1.5,
  },
}));

export const StyledCardContent = styled(CardContent)(({ theme }) => ({
  display: 'flex',
  flexDirection: 'column',
  gap: theme.spacing(1.5),
  padding: '0 !important',
  height: '100%',
}));

export const StyledChip = styled(Typography)(({ theme }) => ({
  backgroundColor: theme.palette.grey[70],
  color: theme.palette.text.secondary,
  fontSize: theme.spacing(1.5),
  fontWeight: 400,
  borderRadius: theme.shape.borderRadius,
  padding: theme.spacing(0.25, 0.5),
}));

export const StyledButton = styled(Button)<ButtonProps>(({ theme }) => ({
  borderRadius: theme.shape.borderRadius * 2,
  boxShadow: '0px 1px 1px 0px rgba(0, 0, 0, 0.05)',
  minHeight: 34,
  padding: theme.spacing(0.75, 1.5),
  marginTop: 'auto',
  border: `1px solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.common.white,
  color: theme.palette.text.primary,
  fontWeight: 500,
  '&:hover': {
    backgroundColor: theme.palette.grey[20],
  },
}));

export const DrawerContainer = styled(Box)<BoxProps>(() => ({
  display: 'flex',
  flexDirection: 'column',
  height: '100%',
  position: 'relative',
}));

export const DrawerContentArea = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$hasFooter',
})<BoxProps & { $hasFooter?: boolean }>(({ theme, $hasFooter }) => ({
  flex: 1,
  overflowY: 'auto',
  paddingBottom: $hasFooter ? theme.spacing(10) : 0,
}));

export const DrawerFooter = styled(Box)<BoxProps>(({ theme }) => ({
  position: 'absolute',
  bottom: 0,
  left: 0,
  right: 0,
  padding: theme.spacing(2),
  backgroundColor: theme.palette.background.paper,
  borderTop: `1px solid ${theme.palette.divider}`,
  display: 'flex',
  justifyContent: 'flex-end',
  gap: theme.spacing(1),
}));

export const SearchContainer = styled(Box)<BoxProps>(({ theme }) => ({
  display: 'flex',
  gap: theme.spacing(2),
  alignItems: 'center',
}));

export const SearchTextField = styled(TextField)<TextFieldProps>(({ theme }) => ({
  '& .MuiOutlinedInput-root': {
    backgroundColor: theme.palette.background.paper,
  },
}));

export const EmptyStateBox = styled(Box)<BoxProps>(({ theme }) => ({
  textAlign: 'center',
  paddingTop: theme.spacing(4),
  paddingBottom: theme.spacing(4),
  color: theme.palette.text.secondary,
}));

export const StyledTabs = styled(Tabs)<TabsProps>(() => ({
  width: '100%',
}));
