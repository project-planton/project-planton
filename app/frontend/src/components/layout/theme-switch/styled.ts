import { PCThemeType } from '@/contexts/models';
import { Button, styled } from '@mui/material';

export const StyledThemeButton = styled(Button, {
  shouldForwardProp: (prop) => prop !== '$mode',
})<{ $mode: PCThemeType }>(({ theme, $mode }) => ({
  display: 'flex',
  flexDirection: 'row',
  alignItems: 'center',
  padding: theme.spacing(0.75),
  borderRadius: theme.shape.borderRadius * 50,
  border: `1px solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.grey[70],
  minWidth: 'unset',
  height: 'auto',
  '& .MuiButton-icon': {
    margin: 0,
    borderRadius: '50%',
  },
  '& .MuiButton-startIcon': {
    padding: theme.spacing(0.25),
    backgroundColor: $mode === 'light' ? theme.palette.background.default : 'unset',
    ...($mode === 'light' && {
      '& *': {
        fill: theme.palette.text.primary,
        stroke: theme.palette.text.primary,
      },
    }),
  },
  '& .MuiButton-endIcon': {
    padding: theme.spacing(0.25),
    backgroundColor: $mode === 'dark' ? theme.palette.background.default : 'unset',
    ...($mode === 'dark' && {
      '& *': {
        fill: theme.palette.text.primary,
        stroke: theme.palette.text.primary,
      },
    }),
  },
}));

