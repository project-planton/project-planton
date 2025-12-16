import { Box, Button, Dialog, DialogContent, styled } from '@mui/material';

export const StyledSecretBox = styled(Box)(({ theme }) => ({
  padding: theme.spacing(1.5),
  border: `1px solid ${theme.palette.divider}`,
  borderRadius: theme.shape.borderRadius * 3,
  wordWrap: 'break-word',
  fontSize: 12,
  fontWeight: 400,
  color: theme.palette.text.primary,
  maxHeight: 500,
  overflow: 'auto',
  whiteSpace: 'pre-line',
}));

export const StyledButton = styled(Button)(({ theme }) => ({
  minWidth: 20,
  paddingLeft: theme.spacing(0.5),
  paddingRight: theme.spacing(0.5),
  marginLeft: theme.spacing(1.25),
}));

export const StyledCopyBtn = styled(Button)(() => ({
  width: 'fit-content',
  marginLeft: 'auto',
}));

export const StyledDialog = styled(Dialog)(({ theme }) => ({
  '& .MuiDialog-paper': {
    border: `1px solid ${theme.palette.divider}`,
    borderRadius: 3 * theme.shape.borderRadius,
    margin: theme.spacing(1),
    [theme.breakpoints.up('md')]: {
      margin: theme.spacing(3),
    },
    minHeight: '150px',
    maxHeight: 'unset',
  },
  '& .MuiDialogTitle-root': {
    borderBottom: `1px solid ${theme.palette.divider}`,
    padding: theme.spacing(1.5, 2),
  },
}));

export const StyledDialogContent = styled(DialogContent)(({ theme }) => ({
  minHeight: '350px',
  paddingBottom: theme.spacing(0),
  overflow: 'auto',
  padding: theme.spacing(0, 1),
  [theme.breakpoints.up('md')]: {
    padding: theme.spacing(0, 3),
  },
}));

