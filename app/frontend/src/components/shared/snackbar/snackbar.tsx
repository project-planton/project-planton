'use client';
import { Snackbar, styled, alpha } from '@mui/material';
import { forwardRef } from 'react';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { Severity } from '@/contexts/models';

const Alert = forwardRef<HTMLDivElement, AlertProps>((props, ref) => {
  return <MuiAlert elevation={0} ref={ref} variant="standard" {...props} />;
});

Alert.displayName = 'Alert';

const StyledAlert = styled(Alert)(({ theme, severity }) => {
  const colors = {
    success: theme.palette.success.main,
    warning: theme.palette.warning.main,
    error: theme.palette.error.main,
    info: theme.palette.info.main,
  };
  
  const bgColors = {
    success: theme.palette.mode === 'dark'
      ? alpha(theme.palette.success.main, 0.15)
      : alpha(theme.palette.success.main, 0.1),
    warning: theme.palette.mode === 'dark'
      ? alpha(theme.palette.warning.main, 0.15)
      : alpha(theme.palette.warning.main, 0.1),
    error: theme.palette.mode === 'dark'
      ? alpha(theme.palette.error.main, 0.15)
      : alpha(theme.palette.error.main, 0.1),
    info: theme.palette.mode === 'dark'
      ? alpha(theme.palette.info.main, 0.15)
      : alpha(theme.palette.info.main, 0.1),
  };
  
  const color = colors[severity || 'info'];
  const bgColor = bgColors[severity || 'info'];
  
  return {
    alignItems: 'center',
    borderRadius: 10,
    padding: theme.spacing(1, 2),
    backgroundColor: theme.palette.background.paper,
    border: `1px solid ${alpha(color, 0.3)}`,
    boxShadow: theme.palette.mode === 'dark'
      ? '0 8px 32px rgba(0, 0, 0, 0.4)'
      : '0 8px 32px rgba(0, 0, 0, 0.1)',
    backdropFilter: 'blur(12px)',
    
    '& .MuiAlert-icon': {
      padding: 0,
      marginRight: theme.spacing(1.5),
      color: color,
    },
    
    '& .MuiAlert-message': {
      padding: 0,
      fontSize: '0.875rem',
      fontWeight: 500,
      color: theme.palette.text.primary,
    },
    
    '& .MuiAlert-action': {
      padding: 0,
      marginRight: 0,
      marginLeft: theme.spacing(1.5),
      
      '& .MuiIconButton-root': {
        padding: 4,
        borderRadius: 6,
        color: theme.palette.text.secondary,
        
        '&:hover': {
          backgroundColor: theme.palette.mode === 'dark'
            ? alpha(theme.palette.common.white, 0.08)
            : alpha(theme.palette.common.black, 0.06),
          color: theme.palette.text.primary,
        },
      },
    },
  };
});

const StyledSnackbar = styled(Snackbar)(({ theme }) => ({
  '& .MuiSnackbar-root': {
    maxWidth: 400,
  },
}));

interface SnackBarProps {
  id?: string;
  open: boolean;
  handleClose: () => void;
  handleExited: () => void;
  severity?: Severity;
  message?: string;
  autoHideDuration?: number;
}

export const SnackBar = (props: SnackBarProps) => {
  const { id, open, handleClose, handleExited, severity, message, autoHideDuration } = props;

  return (
    <StyledSnackbar
      key={id}
      open={open}
      autoHideDuration={autoHideDuration || 5000}
      onClose={handleClose}
      slotProps={{ transition: { onExited: handleExited } }}
      anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
    >
      <StyledAlert onClose={handleClose} severity={severity}>
        {message}
      </StyledAlert>
    </StyledSnackbar>
  );
};
