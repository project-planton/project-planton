'use client';
import { Snackbar, styled } from '@mui/material';
import { forwardRef } from 'react';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { Severity } from '@/contexts/models';

const Alert = forwardRef<HTMLDivElement, AlertProps>((props, ref) => {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

Alert.displayName = 'Alert';

const StyledAlert = styled(Alert)(() => ({
  alignItems: 'center',
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
    <Snackbar
      key={id}
      open={open}
      autoHideDuration={autoHideDuration || 5000}
      onClose={handleClose}
      slotProps={{ transition: { onExited: handleExited } }}
      anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      sx={{ alignItems: 'center' }} // Center alignment
    >
      <StyledAlert onClose={handleClose} severity={severity}>
        {message}
      </StyledAlert>
    </Snackbar>
  );
};
