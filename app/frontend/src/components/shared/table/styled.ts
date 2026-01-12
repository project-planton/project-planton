import Image from 'next/image';
import {
  Box,
  Button,
  IconButton,
  InputLabel,
  Link,
  LinkProps,
  Menu,
  MenuItem,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  styled,
  alpha,
} from '@mui/material';
import { MoreHoriz } from '@mui/icons-material';
import { TextCopy } from '@/components/shared/text-copy';

export const StyledTableContainer = styled(TableContainer, {
  shouldForwardProp: (prop) =>
    prop !== '$stickyHeader' &&
    prop !== '$border' &&
    prop !== '$borderColor' &&
    prop !== '$bgColor',
})<{ $stickyHeader?: boolean; $border?: boolean; $bgColor?: string }>(
  ({ theme, $stickyHeader = false, $border = false, $bgColor }) => ({
    maxHeight: $stickyHeader ? '100vh' : 'unset',
    backgroundColor: $bgColor ? $bgColor : 'transparent',
    borderRadius: 12,
    border: $border ? `1px solid ${theme.palette.divider}` : 'none',
    overflow: 'hidden',

    // Subtle scrollbar styling
    '&::-webkit-scrollbar': {
      width: 8,
      height: 8,
    },
    '&::-webkit-scrollbar-thumb': {
      borderRadius: 8,
      backgroundColor: theme.palette.mode === 'dark'
        ? alpha(theme.palette.common.white, 0.15)
        : alpha(theme.palette.common.black, 0.15),
    },
    '&::-webkit-scrollbar-track': {
      backgroundColor: 'transparent',
    },
  })
);

export const StyledTableHead = styled(TableHead, {
  shouldForwardProp: (prop) =>
    prop !== '$headerBgColor' && prop !== '$border' && prop !== '$borderColor',
})<{ $headerBgColor?: string; $border?: boolean; $borderColor?: string }>(
  ({ theme, $headerBgColor, $border, $borderColor }) => ({
    backgroundColor: $headerBgColor || (theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.03)
      : alpha(theme.palette.common.black, 0.02)),

    // Sticky header with blur
    position: 'sticky',
    top: 0,
    zIndex: 1,
    backdropFilter: 'blur(8px)',

    '& tr': {
      borderBottom: `1px solid ${$borderColor || theme.palette.divider}`,
    },
  })
);

export const StyledTextCopy = styled(TextCopy)`
  visibility: hidden;
  opacity: 0;
  transition: opacity 150ms ease;
`;

export const StyledTableCell = styled(TableCell, {
  shouldForwardProp: (prop) =>
    prop !== '$isFixed' && prop !== '$fixedBgColor' && prop !== '$verticalAlign' && prop !== '$border',
})<{
  $isFixed?: boolean;
  $fixedBgColor?: string;
  $border?: boolean;
  $verticalAlign?: string;
}>(({ theme, $isFixed, $fixedBgColor, $border, $verticalAlign = 'middle' }) => {
  let baseStyles = {
    fontSize: '0.8125rem', // 13px
    lineHeight: 1.4,
    whiteSpace: 'nowrap' as const,
    padding: theme.spacing(1.5, 2),
    borderBottom: 'none',
    color: theme.palette.text.primary,
    transition: 'background-color 150ms ease',

    [`&:hover ${StyledTextCopy}`]: {
      visibility: 'visible' as const,
      opacity: 1,
    },
    verticalAlign: $verticalAlign,
  };

  if ($isFixed) {
    baseStyles = {
      ...baseStyles,
      position: 'sticky' as const,
      right: 0,
      width: '120px',
      textAlign: 'center' as const,
      backgroundColor: $fixedBgColor || theme.palette.background.paper,
      borderLeft: `1px solid ${theme.palette.divider}`,
    };
  }
  return baseStyles;
});

export const StyledCheckboxTableCell = styled(StyledTableCell)`
  width: 48px;
  padding-left: 16px;
`;

export const StyledInputLabel = styled(InputLabel)(({ theme }) => ({
  textAlign: 'left',
  color: theme.palette.text.primary,
  lineHeight: 1.4,
  display: 'inline-block',
  width: '100%',
  fontSize: '0.8125rem',
}));

export const StyledHeaderLabel = styled(StyledInputLabel)(({ theme }) => ({
  fontSize: '0.75rem', // 12px
  fontWeight: 600,
  color: theme.palette.text.secondary,
  textTransform: 'uppercase',
  letterSpacing: '0.02em',
  width: 'auto',
}));

export const StyledTableRow = styled(TableRow, {
  shouldForwardProp: (prop) =>
    prop !== '$headerBgColor' && prop !== '$border' && prop !== '$borderColor',
})<{ $border?: boolean; $borderColor?: string }>(({ theme, $border, $borderColor }) => ({
  borderBottom: `1px solid ${$borderColor || theme.palette.divider}`,
  transition: 'background-color 150ms ease',

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.03)
      : alpha(theme.palette.common.black, 0.02),
  },

  '&:last-child': {
    borderBottom: 'none',
  },
}));

export const StyledLink = styled(Link)<LinkProps>(({ theme }) => ({
  fontSize: '0.8125rem',
  fontWeight: 500,
  cursor: 'pointer',
  overflowWrap: 'anywhere',
  textDecoration: 'none',
  color: theme.palette.primary.main,
  transition: 'color 150ms ease',

  '&:hover': {
    textDecoration: 'underline',
  },
}));

export const StyledMenuIcon = styled(MoreHoriz)(({ theme }) => ({
  fontSize: 20,
  color: theme.palette.text.secondary,
}));

export const StyledIconButton = styled(IconButton)(({ theme }) => ({
  padding: 6,
  borderRadius: 6,
  transition: 'all 150ms ease',

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.black, 0.06),

    '& svg': {
      color: theme.palette.text.primary,
    },
  },
}));

export const StyledMenu = styled(Menu)(({ theme }) => ({
  '& .MuiPaper-root': {
    marginTop: theme.spacing(0.5),
    minWidth: 180,
    borderRadius: 10,
    border: `1px solid ${theme.palette.divider}`,
    boxShadow: theme.palette.mode === 'dark'
      ? '0 8px 32px rgba(0, 0, 0, 0.4)'
      : '0 8px 32px rgba(0, 0, 0, 0.1)',
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(0.5),
  },
}));

export const StyledMenuItem = styled(MenuItem)(({ theme }) => ({
  borderRadius: 6,
  margin: 0,
  padding: theme.spacing(1, 1.5),
  fontSize: '0.8125rem',
  fontWeight: 500,
  transition: 'background-color 150ms ease',

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.05)
      : alpha(theme.palette.common.black, 0.04),
  },

  '&:not(:last-child)': {
    marginBottom: theme.spacing(0.25),
  },
}));

export const StyledPaginationContainer = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$border' && prop !== '$borderColor' && prop !== '$bgColor',
})<{
  $border?: boolean;
  $borderColor?: string;
  $bgColor?: string;
}>(({ theme, $borderColor, $bgColor, $border }) => ({
  backgroundColor: $bgColor ? $bgColor : 'transparent',
  borderTop: `1px solid ${$borderColor || theme.palette.divider}`,
  borderRadius: 0,
  padding: theme.spacing(1, 0),
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  gap: theme.spacing(1),

  '& .MuiTablePagination-selectIcon': {
    color: theme.palette.text.secondary,
  },
}));

export const StyledPaginationActionButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.secondary,
  padding: 6,
  borderRadius: 6,

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.black, 0.06),
    color: theme.palette.text.primary,
  },

  '&.Mui-disabled': {
    color: theme.palette.text.disabled,
  },
}));

export const StyledNextImage = styled(Image)`
  cursor: pointer;
  border-radius: 50%;
`;

export const StyledActionButton = styled(Button)(({ theme }) => ({
  minWidth: 'fit-content',
  padding: theme.spacing(0.5, 1),
  fontSize: '0.8125rem',
  fontWeight: 500,
}));

export const StyledPageNumBtn = styled(Button)(({ theme }) => ({
  minWidth: 32,
  height: 32,
  padding: 0,
  fontWeight: 500,
  fontSize: '0.8125rem',
  borderRadius: 6,
  color: theme.palette.text.secondary,

  '&.MuiButton-contained': {
    color: theme.palette.primary.contrastText,
  },

  '&.active': {
    backgroundColor: theme.palette.primary.main,
    color: theme.palette.primary.contrastText,
  },

  '&:hover:not(.active):not(.MuiButton-contained)': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.black, 0.06),
  },
}));

export const StyledNextPrevBtn = styled(Button)(({ theme }) => ({
  minWidth: 32,
  height: 32,
  padding: 0,
  fontWeight: 500,
  fontSize: '0.8125rem',
  borderRadius: 6,
  color: theme.palette.text.secondary,

  '&:hover': {
    backgroundColor: theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.black, 0.06),
  },

  '&.Mui-disabled': {
    color: theme.palette.text.disabled,
  },
}));

// Status badge for table cells
export const StatusBadge = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$status',
})<{ $status?: 'success' | 'warning' | 'error' | 'info' | 'default' }>(
  ({ theme, $status = 'default' }) => {
    const colors = {
      success: theme.palette.success.main,
      warning: theme.palette.warning.main,
      error: theme.palette.error.main,
      info: theme.palette.info.main,
      default: theme.palette.text.secondary,
    };

    const bgColors = {
      success: alpha(theme.palette.success.main, 0.1),
      warning: alpha(theme.palette.warning.main, 0.1),
      error: alpha(theme.palette.error.main, 0.1),
      info: alpha(theme.palette.info.main, 0.1),
      default: theme.palette.mode === 'dark'
        ? alpha(theme.palette.common.white, 0.08)
        : alpha(theme.palette.common.black, 0.06),
    };

    return {
      display: 'inline-flex',
      alignItems: 'center',
      gap: theme.spacing(0.5),
      padding: theme.spacing(0.25, 1),
      borderRadius: 4,
      fontSize: '0.75rem',
      fontWeight: 500,
      backgroundColor: bgColors[$status],
      color: colors[$status],

      '&::before': {
        content: '""',
        width: 6,
        height: 6,
        borderRadius: '50%',
        backgroundColor: 'currentColor',
      },
    };
  }
);