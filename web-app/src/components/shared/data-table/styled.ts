'use client';
import {
  Box,
  Table,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TableSortLabel,
  styled,
} from '@mui/material';

export const StyledTableContainer = styled(TableContainer)(({ theme }) => ({
  borderRadius: theme.shape.borderRadius * 2,
  border: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.text.primary,
  overflow: 'hidden',
}));

export const StyledTable = styled(Table)(({ theme }) => ({
  minWidth: theme.spacing(81.25), // 650px / 8 = 81.25
}));

export const StyledTableHead = styled(TableHead)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' 
    ? theme.palette.grey[800] || theme.palette.background.default
    : theme.palette.grey[100] || theme.palette.background.default,
  '& .MuiTableCell-root': {
    fontWeight: theme.typography.fontWeightBold,
    fontSize: theme.typography.body2.fontSize,
    color: theme.palette.text.primary,
    borderBottom: `${theme.spacing(0.25)} solid ${theme.palette.divider}`,
    backgroundColor: 'transparent',
  },
}));

export const StyledTableRow = styled(TableRow)(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.text.primary,
  '&:hover': {
    backgroundColor: theme.palette.action.hover,
  },
  '&.Mui-selected': {
    backgroundColor: theme.palette.action.selected,
    '&:hover': {
      backgroundColor: theme.palette.action.selected,
    },
  },
}));

export const StyledTableCell = styled(TableCell)(({ theme }) => ({
  padding: theme.spacing(1.5, 2),
  fontSize: theme.typography.body2.fontSize,
  borderBottom: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  color: theme.palette.text.primary,
}));

export const StyledTableSortLabel = styled(TableSortLabel)(({ theme }) => ({
  color: theme.palette.text.secondary,
  '&:hover': {
    color: theme.palette.text.primary,
  },
  '&.Mui-active': {
    color: theme.palette.text.primary,
    '& .MuiTableSortLabel-icon': {
      color: theme.palette.primary.main,
    },
  },
}));

export const StyledCheckboxTableCell = styled(TableCell)(({ theme }) => ({
  padding: theme.spacing(1),
  width: theme.spacing(6),
  borderBottom: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  '&.MuiTableCell-head': {
    borderBottom: `${theme.spacing(0.25)} solid ${theme.palette.divider}`,
  },
}));

export const EmptyStateContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(4),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));

export const LoadingContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3),
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
}));

export const ActionCell = styled(Box)(({ theme }) => ({
  display: 'flex',
  gap: theme.spacing(1),
  alignItems: 'center',
  justifyContent: 'flex-end',
}));

export const MenuItemIconContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(1),
  marginRight: theme.spacing(1),
}));

export const StatusBadge = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$status',
})<{ $status?: 'active' | 'inactive' }>(({ theme, $status }) => ({
  padding: theme.spacing(0.5, 1.5),
  borderRadius: theme.shape.borderRadius * 3,
  fontSize: theme.typography.caption.fontSize,
  fontWeight: theme.typography.fontWeightBold,
  display: 'inline-block',
  backgroundColor:
    $status === 'active'
      ? theme.palette.mode === 'dark'
        ? theme.palette.success.dark
        : theme.palette.success.light
      : theme.palette.mode === 'dark'
      ? theme.palette.error.dark
      : theme.palette.error.light,
  color:
    $status === 'active'
      ? theme.palette.mode === 'dark'
        ? theme.palette.success.light
        : theme.palette.success.dark
      : theme.palette.mode === 'dark'
      ? theme.palette.error.light
      : theme.palette.error.dark,
}));

export const StyledTablePagination = styled(Box)(({ theme }) => ({
  borderTop: `${theme.spacing(0.125)} solid ${theme.palette.divider}`,
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.text.primary,
  '& .MuiTablePagination-root': {
    color: theme.palette.text.primary,
  },
  '& .MuiTablePagination-selectLabel, & .MuiTablePagination-displayedRows': {
    color: theme.palette.text.primary,
  },
  '& .MuiIconButton-root': {
    color: theme.palette.text.primary,
    '&:hover': {
      backgroundColor: theme.palette.action.hover,
    },
    '&.Mui-disabled': {
      color: theme.palette.text.disabled,
    },
  },
  '& .MuiSelect-root': {
    color: theme.palette.text.primary,
  },
}));

export const ActionsTableCell = styled(StyledTableCell)(({ theme }) => ({
  minWidth: theme.spacing(12.5), // 100px / 8 = 12.5
}));
