'use client';
import React from 'react';
import { IconButton, Typography, Divider } from '@mui/material';
import { Close as CloseIcon } from '@mui/icons-material';
import { StyledDrawer, DrawerHeader, DrawerContent } from './styled';

interface DrawerProps {
  open: boolean;
  onClose: () => void;
  title: string;
  anchor?: 'left' | 'right' | 'top' | 'bottom';
  width?: number | string;
  children: React.ReactNode;
}

export const Drawer: React.FC<DrawerProps> = ({
  open,
  onClose,
  title,
  anchor = 'right',
  width = 600,
  children,
}) => {
  return (
    <StyledDrawer
      anchor={anchor}
      open={open}
      onClose={onClose}
      slotProps={{
        paper: {
          sx: {
            width: typeof width === 'number' ? `${width}px` : width,
          },
        },
      }}
    >
      <DrawerHeader>
        <Typography variant="h6" component="h2">
          {title}
        </Typography>
        <IconButton
          onClick={onClose}
          size="small"
          sx={{
            color: 'text.secondary',
            '&:hover': {
              backgroundColor: 'action.hover',
            },
          }}
        >
          <CloseIcon />
        </IconButton>
      </DrawerHeader>
      <Divider />
      <DrawerContent>{children}</DrawerContent>
    </StyledDrawer>
  );
};

