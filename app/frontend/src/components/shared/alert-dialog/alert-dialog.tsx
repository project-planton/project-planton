'use client';
import { FC, ReactNode } from 'react';

import {
  Button,
  ButtonOwnProps,
  ButtonProps,
  Dialog,
  DialogContent,
  DialogProps,
  PaperProps,
  Stack,
  Typography,
  styled,
} from '@mui/material';

import { FlexCenterRow } from '@/components/shared/resource-header/styled';

export interface AlertDialogProps {
  open: DialogProps['open'];
  title: string;
  subTitle?: string | ReactNode;
  submitLabel?: string;
  submitBtnColor?: ButtonOwnProps['color'];
  cancelLabel?: string;
  onClose: DialogProps['onClose'] | ButtonProps['onClick'];
  onSubmit?: ButtonProps['onClick'] | (() => void);
  extraContent?: ReactNode;
  extraContentPosition?: 'top' | 'bottom';
  paperProps?: PaperProps;
  loading?: boolean;
}

export const defaultAlertDialogProps: AlertDialogProps = {
  open: false,
  onClose: null,
  title: 'Are you sure?',
};

const StyledDialogContent = styled(DialogContent)(({ theme }) => ({
  padding: theme.spacing(2),
  [theme.breakpoints.up('md')]: {
    padding: theme.spacing(3),
  },
}));

export const AlertDialog: FC<AlertDialogProps> = ({
  open,
  title,
  subTitle,
  submitLabel = 'Yes',
  submitBtnColor = 'primary',
  cancelLabel = 'No',
  onClose,
  onSubmit,
  extraContent = null,
  extraContentPosition = 'top',
  paperProps,
  loading,
}) => {
  return (
    <Dialog
      open={open}
      onClose={onClose as DialogProps['onClose']}
      slotProps={{
        paper: {
          sx: {
            borderRadius: '12px',
            textAlign: 'center',
            minWidth: 300,
            maxWidth: 540,
            ...paperProps?.sx,
          },
        },
      }}
    >
      <StyledDialogContent>
        <Stack gap={1}>
          {extraContent && extraContentPosition === 'top' && extraContent}
          <Typography fontSize={14} fontWeight="500">
            {title}
          </Typography>
          {subTitle && (
            <Typography variant="subtitle2" color="text.secondary">
              {subTitle}
            </Typography>
          )}
          {extraContent && extraContentPosition === 'bottom' && extraContent}
        </Stack>
        <FlexCenterRow gap={1.5} justifyContent="center" marginTop="16px">
          <Button
            variant="contained"
            color="secondary"
            onClick={onClose as ButtonProps['onClick']}
            size="medium"
          >
            {cancelLabel}
          </Button>
          <Button
            variant="contained"
            color={submitBtnColor}
            onClick={onSubmit}
            size="medium"
            loading={loading}
          >
            {submitLabel}
          </Button>
        </FlexCenterRow>
      </StyledDialogContent>
    </Dialog>
  );
};
