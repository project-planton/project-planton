'use client';
import { Box, Grid2, Paper, styled, Typography, alpha } from '@mui/material';

export const DashboardContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
  maxWidth: 1400,
  margin: '0 auto',
  width: '100%',
}));

export const DashboardHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  marginBottom: theme.spacing(3),
  gap: theme.spacing(2),
  flexWrap: 'wrap',
}));

export const DashboardTitle = styled(Typography)(({ theme }) => ({
  fontSize: '1.5rem',
  fontWeight: 600,
  lineHeight: 1.25,
  letterSpacing: '-0.01em',
  color: theme.palette.text.primary,
  margin: 0,
}));

export const DashboardSubtitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.875rem',
  color: theme.palette.text.secondary,
  marginTop: theme.spacing(0.5),
}));

export const StyledGrid2 = styled(Grid2)(({ theme }) => ({
  // Grid spacing is now handled by parent container
  marginTop: theme.spacing(4),
}));

// Modern Stat Card
export const StyledCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(2.5),
  borderRadius: 12,
  backgroundColor: theme.palette.background.paper,
  border: `1px solid ${theme.palette.divider}`,
  boxShadow: 'none',
  transition: 'all 150ms ease',
  cursor: 'default',
  position: 'relative',
  overflow: 'hidden',

  '&:hover': {
    borderColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.15)
      : alpha(theme.palette.common.black, 0.15),
    boxShadow: theme.palette.mode === 'dark'
      ? '0 4px 20px rgba(0, 0, 0, 0.3)'
      : '0 4px 20px rgba(0, 0, 0, 0.08)',
  },
}));

// Accent variant for highlighted stats
export const StyledAccentCard = styled(StyledCard)(({ theme }) => ({
  background: theme.palette.mode === 'dark'
    ? `linear-gradient(135deg, ${alpha(theme.palette.primary.main, 0.15)} 0%, ${alpha(theme.palette.primary.main, 0.05)} 100%)`
    : `linear-gradient(135deg, ${alpha(theme.palette.primary.main, 0.1)} 0%, ${alpha(theme.palette.primary.main, 0.02)} 100%)`,
  borderColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.primary.main, 0.3)
    : alpha(theme.palette.primary.main, 0.2),

  '&:hover': {
    borderColor: theme.palette.primary.main,
    boxShadow: `0 4px 20px ${alpha(theme.palette.primary.main, 0.2)}`,
  },
}));

export const StatCardIcon = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 40,
  height: 40,
  borderRadius: 10,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.primary.main, 0.15)
    : alpha(theme.palette.primary.main, 0.1),
  color: theme.palette.primary.main,
  marginBottom: theme.spacing(2),

  '& svg': {
    fontSize: 22,
  },
}));

export const StatCardTitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.8125rem', // 13px
  fontWeight: 500,
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(0.5),
  lineHeight: 1.4,
}));

export const StatCardValue = styled(Typography)(({ theme }) => ({
  fontSize: '1.75rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
  lineHeight: 1.2,
  letterSpacing: '-0.01em',
}));

export const StatCardTrend = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$positive',
})<{ $positive?: boolean }>(({ theme, $positive }) => ({
  display: 'inline-flex',
  alignItems: 'center',
  gap: theme.spacing(0.5),
  marginTop: theme.spacing(1),
  padding: theme.spacing(0.25, 0.75),
  borderRadius: 4,
  fontSize: '0.75rem',
  fontWeight: 500,
  backgroundColor: $positive
    ? alpha(theme.palette.success.main, 0.1)
    : alpha(theme.palette.error.main, 0.1),
  color: $positive
    ? theme.palette.success.main
    : theme.palette.error.main,

  '& svg': {
    fontSize: 14,
  },
}));

export const StatCardSubtext = styled(Typography)(({ theme }) => ({
  fontSize: '0.75rem',
  color: theme.palette.text.secondary,
  marginTop: theme.spacing(1),
}));

// Table Section
export const TableSection = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(4),
}));

export const TableSectionHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  marginBottom: theme.spacing(2),
  gap: theme.spacing(2),
}));

export const TableSectionTitle = styled(Typography)(({ theme }) => ({
  fontSize: '1rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
}));

// Empty State
export const EmptyStateContainer = styled(Box)(({ theme }) => ({
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

export const EmptyStateIcon = styled(Box)(({ theme }) => ({
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

export const EmptyStateTitle = styled(Typography)(({ theme }) => ({
  fontSize: '1rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
  marginBottom: theme.spacing(0.5),
}));

export const EmptyStateDescription = styled(Typography)(({ theme }) => ({
  fontSize: '0.875rem',
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(2),
  maxWidth: 400,
}));

// Quick Actions Grid
export const QuickActionsGrid = styled(Box)(({ theme }) => ({
  display: 'grid',
  gap: theme.spacing(1.5),
  gridTemplateColumns: 'repeat(2, 1fr)',
  marginTop: theme.spacing(2),

  [theme.breakpoints.up('sm')]: {
    gridTemplateColumns: 'repeat(4, 1fr)',
  },
}));

export const QuickActionCard = styled(Paper)(({ theme }) => ({
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  padding: theme.spacing(2),
  borderRadius: 10,
  backgroundColor: theme.palette.background.paper,
  border: `1px solid ${theme.palette.divider}`,
  boxShadow: 'none',
  cursor: 'pointer',
  transition: 'all 150ms ease',
  textAlign: 'center',
  gap: theme.spacing(1),

  '&:hover': {
    borderColor: theme.palette.primary.main,
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.primary.main, 0.08)
      : alpha(theme.palette.primary.main, 0.05),
    transform: 'translateY(-2px)',
    boxShadow: `0 4px 12px ${alpha(theme.palette.primary.main, 0.15)}`,
  },

  '& svg': {
    fontSize: 24,
    color: theme.palette.primary.main,
  },
}));

export const QuickActionLabel = styled(Typography)(({ theme }) => ({
  fontSize: '0.8125rem',
  fontWeight: 500,
  color: theme.palette.text.primary,
}));

// Resource Card for grid view
export const ResourceCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(2),
  borderRadius: 12,
  backgroundColor: theme.palette.background.paper,
  border: `1px solid ${theme.palette.divider}`,
  boxShadow: 'none',
  transition: 'all 150ms ease',
  cursor: 'pointer',

  '&:hover': {
    borderColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.15)
      : alpha(theme.palette.common.black, 0.12),
    boxShadow: theme.palette.mode === 'dark'
      ? '0 4px 16px rgba(0, 0, 0, 0.25)'
      : '0 4px 16px rgba(0, 0, 0, 0.06)',
  },
}));

export const ResourceCardHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'flex-start',
  justifyContent: 'space-between',
  marginBottom: theme.spacing(1.5),
}));

export const ResourceCardTitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.875rem',
  fontWeight: 600,
  color: theme.palette.text.primary,
  lineHeight: 1.4,
}));

export const ResourceCardDescription = styled(Typography)(({ theme }) => ({
  fontSize: '0.8125rem',
  color: theme.palette.text.secondary,
  lineHeight: 1.5,
  marginBottom: theme.spacing(1.5),
  display: '-webkit-box',
  WebkitLineClamp: 2,
  WebkitBoxOrient: 'vertical',
  overflow: 'hidden',
}));

export const ResourceCardMeta = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(1.5),
  flexWrap: 'wrap',
}));

export const ResourceCardTag = styled(Box)(({ theme }) => ({
  display: 'inline-flex',
  alignItems: 'center',
  padding: theme.spacing(0.25, 0.75),
  borderRadius: 4,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.08)
    : alpha(theme.palette.common.black, 0.06),
  fontSize: '0.6875rem',
  fontWeight: 500,
  color: theme.palette.text.secondary,
}));
