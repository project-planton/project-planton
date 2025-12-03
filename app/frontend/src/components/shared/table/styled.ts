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
    backgroundColor: $bgColor ? $bgColor : theme.palette.background.default,
    padding: theme.spacing(!$border ? 1 : 0),
  })
);

export const StyledTableHead = styled(TableHead, {
  shouldForwardProp: (prop) => prop !== '$headerBgColor',
})<{ $headerBgColor?: string; $border?: boolean; $borderColor?: string }>(
  ({ theme, $headerBgColor, $border, $borderColor }) => ({
    backgroundColor: $headerBgColor || theme.palette.grey[90],
    '& tr': {
      borderBottom: $border
        ? `1px solid ${$borderColor ? $borderColor : theme.palette.divider}`
        : 'none',
    },
  })
);

export const StyledTextCopy = styled(TextCopy)`
  visibility: hidden;
`;

export const StyledTableCell = styled(TableCell, {
  shouldForwardProp: (prop) =>
    prop !== '$isFixed' && prop !== '$fixedBgColor' && prop !== '$verticalAlign',
})<{
  $isFixed?: boolean;
  $fixedBgColor?: string;
  $border?: boolean;
  $verticalAlign?: string;
}>(({ theme, $isFixed, $fixedBgColor, $border, $verticalAlign = 'middle' }) => {
  let baseStyles = {
    fontSize: theme.spacing(1.5),
    lineHeight: 1,
    whiteSpace: 'nowrap',
    padding: $border ? theme.spacing(1.5) : theme.spacing(0.75, 1.25),
    borderBottom: 'unset',
    [`&:hover ${StyledTextCopy}`]: {
      visibility: 'visible',
    },
    verticalAlign: $verticalAlign,
  };
  if ($isFixed) {
    baseStyles = {
      ...baseStyles,
      position: 'sticky',
      right: 0,
      width: '150px',
      textAlign: 'center',
      backgroundColor: $fixedBgColor || 'initial',
    };
  }
  return baseStyles;
});

export const StyledCheckboxTableCell = styled(StyledTableCell)`
  width: 50px;
`;

export const StyledInputLabel = styled(InputLabel)(({ theme }) => ({
  textAlign: 'left',
  color: theme.palette.text.primary,
  lineHeight: 1.2,
  display: 'inline-block',
  width: '100%',
}));

export const StyledHeaderLabel = styled(StyledInputLabel)(({ theme }) => ({
  fontSize: '12px',
  fontWeight: 500,
  color: theme.palette.text.secondary,
  width: 'auto',
}));

export const StyledTableRow = styled(TableRow, {
  shouldForwardProp: (prop) => prop !== '$headerBgColor',
})<{ $border?: boolean; $borderColor?: string }>(({ theme, $border, $borderColor }) => ({
  borderBottom: $border
    ? `1px solid ${$borderColor ? $borderColor : theme.palette.divider}`
    : 'none',
}));

export const StyledLink = styled(Link)<LinkProps>(({ theme }) => ({
  fontSize: theme.spacing(1.5),
  fontWeight: 400,
  cursor: 'pointer',
  overflowWrap: 'anywhere',
  textDecoration: 'none',
  verticalAlign: 'super',
}));

export const StyledMenuIcon = styled(MoreHoriz)`
  font-size: 18px;
`;

export const StyledIconButton = styled(IconButton)`
  padding: 0;
`;

export const StyledMenu = styled(Menu)(({ theme }) => ({
  '& .MuiPaper-root': {
    marginTop: theme.spacing(0.5),
    minWidth: '190px',
  },
}));

export const StyledMenuItem = styled(MenuItem)(({ theme }) => ({
  borderRadius: 0,
  margin: 0,
  '&:hover': {
    backgroundColor: 'none',
  },
  '&:not(:last-child)': {
    borderBottom: `1px solid ${theme.palette.grey[20]}`,
  },
}));

export const StyledPaginationContainer = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$border' && prop !== '$borderColor' && prop !== '$bgColor',
})<{
  $border?: boolean;
  $borderColor?: string;
  $bgColor?: string;
}>(({ theme, $borderColor, $bgColor, $border }) => ({
  backgroundColor: $bgColor ? $bgColor : theme.palette.background.default,
  border: $border ? `1px solid ${$borderColor ? $borderColor : theme.palette.divider}` : 'none',
  borderRadius: theme.spacing(1),
  '&.MuiTablePagination-selectIcon': {
    color: theme.palette.text.secondary,
  },
}));

export const StyledPaginationActionButton = styled(IconButton)`
  color: ${({ theme }) => theme.palette.text.secondary};
`;

export const StyledNextImage = styled(Image)`
  cursor: pointer;
  border-radius: 50%;
`;

export const StyledActionButton = styled(Button)(() => ({
  minWidth: 'fit-content',
  padding: 0,
}));

export const StyledPageNumBtn = styled(Button)(() => ({
  minWidth: 32,
  fontWeight: 500,
}));

export const StyledNextPrevBtn = styled(Button)(() => ({
  minWidth: 32,
  fontWeight: 500,
}));

