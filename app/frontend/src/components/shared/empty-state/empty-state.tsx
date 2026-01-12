'use client';
import React, { ReactNode } from 'react';
import { Box, Button, styled, alpha, Typography } from '@mui/material';
import { Add as AddIcon } from '@mui/icons-material';

interface EmptyStateProps {
  icon?: ReactNode;
  title: string;
  description?: string;
  actionLabel?: string;
  actionIcon?: ReactNode;
  onAction?: () => void;
  secondaryActionLabel?: string;
  onSecondaryAction?: () => void;
}

const Container = styled(Box)(({ theme }) => ({
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  padding: theme.spacing(6, 3),
  textAlign: 'center',
  backgroundColor: theme.palette.background.paper,
  borderRadius: 12,
  border: `1px dashed ${theme.palette.divider}`,
}));

const IconContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 64,
  height: 64,
  borderRadius: 16,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.05)
    : alpha(theme.palette.common.black, 0.04),
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(2),
  
  '& svg': {
    fontSize: 32,
  },
}));

const Title = styled(Typography)(({ theme }) => ({
  fontSize: '1rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
  marginBottom: theme.spacing(0.5),
}));

const Description = styled(Typography)(({ theme }) => ({
  fontSize: '0.875rem',
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(2.5),
  maxWidth: 400,
  lineHeight: 1.6,
}));

const ActionsContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(1.5),
  flexWrap: 'wrap',
  justifyContent: 'center',
}));

export const EmptyState: React.FC<EmptyStateProps> = ({
  icon,
  title,
  description,
  actionLabel,
  actionIcon = <AddIcon />,
  onAction,
  secondaryActionLabel,
  onSecondaryAction,
}) => {
  return (
    <Container>
      {icon && <IconContainer>{icon}</IconContainer>}
      
      <Title>{title}</Title>
      
      {description && <Description>{description}</Description>}
      
      <ActionsContainer>
        {actionLabel && onAction && (
          <Button
            variant="contained"
            color="primary"
            startIcon={actionIcon}
            onClick={onAction}
            size="medium"
          >
            {actionLabel}
          </Button>
        )}
        
        {secondaryActionLabel && onSecondaryAction && (
          <Button
            variant="outlined"
            color="secondary"
            onClick={onSecondaryAction}
            size="medium"
          >
            {secondaryActionLabel}
          </Button>
        )}
      </ActionsContainer>
    </Container>
  );
};
