import { Breadcrumbs, Stack, styled, Typography } from '@mui/material';

export const StyledBreadcrumbs = styled(Breadcrumbs)(({ theme }) => ({
  minHeight: 28,
  padding: theme.spacing(1.5, 3),
  display: 'flex',
  alignItems: 'center',
  backgroundColor: theme.palette.background.default,
  borderBottom: `1px solid ${theme.palette.divider}`,
  '& .MuiBreadcrumbs-li': {
    display: 'flex',
  },
  '& .MuiBreadcrumbs-separator': {
    margin: theme.spacing(0, 2),
  },
}));

export const BreadcrumbLabel = styled(Typography, {
  shouldForwardProp: (prop) => prop !== '$hasLink',
})<{ $hasLink?: boolean }>(({ theme, $hasLink }) => ({
  fontSize: theme.spacing(1.5),
  fontWeight: 400,
  color: $hasLink ? theme.palette.text.secondary : theme.palette.text.primary,
  cursor: $hasLink ? 'pointer' : 'auto',
}));

export const StyledBreadcrumbStartIcon = styled(Stack)(({ theme }) => ({
  flexDirection: 'row',
  alignItems: 'center',
  cursor: 'pointer',
  padding: theme.spacing(0.5),
  borderRadius: theme.shape.borderRadius,
  '& > *:not(:first-child)': {
    opacity: 0,
    overflow: 'hidden',
    textWrap: 'nowrap',
    maxWidth: '0px',
    transition: 'all 0.5s ease',
  },
  '&:hover': {
    gap: theme.spacing(1),
    '& > *': {
      color: theme.palette.primary.main,
    },
    '& > *:not(:first-child)': {
      opacity: 1,
      maxWidth: '500px',
    },
  },
}));

