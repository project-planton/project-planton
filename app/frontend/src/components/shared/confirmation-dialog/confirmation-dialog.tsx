'use client';
import React, { FC, useState } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  Typography,
} from '@mui/material';
import { ActionType } from '@/models/table';

export interface ConfirmationDialogProps {
  open: boolean;
  message?: string;
  onClose: () => void;
  closeLabel?: string;
  onSubmit?: (reason: string, force?: boolean) => void;
  submitLabel?: string;
  reasonPlaceholder?: string;
  id?: string;
  actionType?: ActionType;
}

export const defaultConfirmationDialogProps: ConfirmationDialogProps = {
  open: false,
  onClose: null,
  closeLabel: 'Cancel',
  submitLabel: 'Confirm',
  message: 'Are you absolutely Sure?',
  actionType: ActionType.UNSPECIFIED,
};

export const ConfirmationDialog: FC<ConfirmationDialogProps> = ({
  open,
  message,
  onClose,
  closeLabel = 'Cancel',
  onSubmit,
  submitLabel = 'Confirm',
  reasonPlaceholder,
}) => {
  const [reason, setReason] = useState('');

  const handleSubmit = () => {
    if (onSubmit) {
      onSubmit(reason);
    }
    setReason('');
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>Confirmation</DialogTitle>
      <DialogContent>
        <Typography variant="body1" sx={{ mb: 2 }}>
          {message || 'Are you absolutely Sure?'}
        </Typography>
        {reasonPlaceholder && (
          <TextField
            fullWidth
            multiline
            rows={3}
            placeholder={reasonPlaceholder}
            value={reason}
            onChange={(e) => setReason(e.target.value)}
            sx={{ mt: 1 }}
          />
        )}
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>{closeLabel}</Button>
        <Button onClick={handleSubmit} variant="contained" color="primary">
          {submitLabel}
        </Button>
      </DialogActions>
    </Dialog>
  );
};
