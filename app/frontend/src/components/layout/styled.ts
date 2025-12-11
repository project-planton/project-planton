'use client';
import { Box, Container, styled } from '@mui/material';

export const StyledWrapperBox = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$hasWhiteBg',
})<{ $hasWhiteBg?: boolean }>(({ theme, $hasWhiteBg }) => ({
  display: 'flex',
  flex: 1,
  minHeight: `calc(100vh - ${theme.spacing(8)})`,
  backgroundColor: $hasWhiteBg ? theme.palette.grey[0] : theme.palette.grey[20],
  marginTop: theme.spacing(8),
  minWidth: 0,
  overflowX: 'hidden',
}));

export const StyledContainer = styled(Container, {
  shouldForwardProp: (prop) => prop !== '$fullWidth' && prop !== '$navbarOpen',
})<{ $fullWidth?: boolean; $navbarOpen?: boolean }>(
  ({ theme, $fullWidth, $navbarOpen }) => ({
    flexGrow: 1,
    padding: $fullWidth ? `${theme.spacing(0)} !important` : `${theme.spacing(3)} !important`,
    width: '100%',
    maxWidth: '100% !important',
    minHeight: `calc(100vh - ${theme.spacing(8)})`,
    minWidth: 0,
    overflowX: 'hidden',
    [theme.breakpoints.up('md')]: {
      maxWidth: $fullWidth
        ? `calc(100% - ${$navbarOpen ? '260px' : '60px'}) !important`
        : '100% !important',
    },
  })
);
