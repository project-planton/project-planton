'use client';
import { Box, Grid2, Paper, styled, Typography } from '@mui/material';

export const DashboardContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
}));

export const StyledGrid2 = styled(Grid2)(({ theme }) => ({
  marginTop: theme.spacing(2),
  marginBottom: theme.spacing(3),
}));

export const StyledCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: theme.shape.borderRadius * 2,
  backgroundColor: theme.palette.background.paper,
  boxShadow: `0 ${theme.spacing(0.125)} ${theme.spacing(0.5)} rgba(0, 0, 0, 0.05)`,
  transition: theme.transitions.create(['box-shadow'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.short,
  }),
  '&:hover': {
    boxShadow: `0 ${theme.spacing(0.25)} ${theme.spacing(0.75)} rgba(0, 0, 0, 0.08)`,
  },
}));

export const StatCardTitle = styled(Typography)(({ theme }) => ({
  fontSize: theme.typography.body2.fontSize,
  fontWeight: theme.typography.fontWeightMedium,
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(1),
  letterSpacing: theme.spacing(0.0625), // 0.5px
}));

export const StatCardValue = styled(Typography)(({ theme }) => ({
  fontSize: theme.typography.h4.fontSize,
  fontWeight: theme.typography.fontWeightBold,
  color: theme.palette.text.primary,
  lineHeight: 1.2,
}));

export const TableSection = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(4),
}));
