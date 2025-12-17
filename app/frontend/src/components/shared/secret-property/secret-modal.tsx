'use client';
import React, { ReactNode, useCallback, useContext } from 'react';
import { Stack, SvgIcon, Typography, DialogTitle, DialogActions, IconButton } from '@mui/material';
import { ContentCopy, Close } from '@mui/icons-material';
import { AppContext } from '@/contexts';
import { copyText } from '@/lib';
import {
  StyledCopyBtn,
  StyledSecretBox,
  StyledDialog,
  StyledDialogContent,
} from '@/components/shared/secret-property/styled';

export interface SecretModalProps {
  open: boolean;
  onClose: () => void;
  title?: ReactNode;
  message?: string;
  secretValue: string;
  styledContent?: boolean;
  minWidth?: number;
  maxWidth?: number;
}

export function SecretModal({
  open,
  onClose,
  title,
  message,
  secretValue,
  styledContent = false,
  minWidth = 480,
  maxWidth = 500,
}: SecretModalProps) {
  const { openSnackbar } = useContext(AppContext);

  const copySecretValue = useCallback(() => {
    if (secretValue) {
      copyText(secretValue).then(() => {
        openSnackbar('Copied!', 'success');
      });
    } else {
      openSnackbar('No value to copy', 'error');
    }
  }, [secretValue, openSnackbar]);

  return (
    <StyledDialog
      open={open}
      onClose={onClose}
      slotProps={{
        paper: {
          sx: {
            minWidth: `${minWidth}px !important`,
            maxWidth: `${maxWidth}px !important`,
          },
        },
      }}
    >
      <DialogTitle>
        <Stack direction="row" width="100%" justifyContent="space-between" alignItems="center">
          {title}
          <IconButton color="inherit" size="small" onClick={onClose} sx={{ marginLeft: 'auto' }}>
            <Close fontSize="small" />
          </IconButton>
        </Stack>
      </DialogTitle>
      <StyledDialogContent sx={{ minHeight: 0, padding: '12px 16px !important' }}>
        <Stack gap={1.5}>
          {!!message && <Typography>{message}</Typography>}
          <StyledSecretBox>
            {styledContent ? (
              <Typography color="text.primary" sx={{ textDecoration: 'underline' }}>
                {secretValue}
              </Typography>
            ) : (
              secretValue
            )}
          </StyledSecretBox>
        </Stack>
      </StyledDialogContent>
      <DialogActions>
        <StyledCopyBtn
          variant="contained"
          color="secondary"
          onClick={copySecretValue}
          startIcon={
            <SvgIcon component={ContentCopy} fontSize="small" sx={{ color: 'text.secondary' }} />
          }
        >
          Copy
        </StyledCopyBtn>
      </DialogActions>
    </StyledDialog>
  );
}
