'use client';
import { Box, Container, styled, alpha } from '@mui/material';
import { headerHeight, miniDrawerWidth, drawerWidth } from './sidebar/styled';

export const StyledWrapperBox = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$hasWhiteBg',
})<{ $hasWhiteBg?: boolean }>(({ theme, $hasWhiteBg }) => ({
  display: 'flex',
  flex: 1,
  minHeight: `calc(100vh - ${headerHeight}px)`,
  backgroundColor: $hasWhiteBg 
    ? theme.palette.background.paper 
    : theme.palette.background.default,
  marginTop: headerHeight,
  minWidth: 0,
  overflowX: 'hidden',
  
  // Subtle background pattern for dark mode
  ...(theme.palette.mode === 'dark' && !$hasWhiteBg && {
    backgroundImage: `radial-gradient(${alpha(theme.palette.common.white, 0.02)} 1px, transparent 1px)`,
    backgroundSize: '24px 24px',
  }),
}));

export const StyledContainer = styled(Container, {
  shouldForwardProp: (prop) => prop !== '$fullWidth' && prop !== '$navbarOpen',
})<{ $fullWidth?: boolean; $navbarOpen?: boolean }>(
  ({ theme, $fullWidth, $navbarOpen }) => ({
    flexGrow: 1,
    padding: $fullWidth ? `0 !important` : `${theme.spacing(3)} !important`,
    width: '100%',
    maxWidth: '100% !important',
    minHeight: `calc(100vh - ${headerHeight}px)`,
    minWidth: 0,
    overflowX: 'hidden',
    
    // Smooth transition when sidebar opens/closes
    transition: theme.transitions.create(['max-width', 'padding'], {
      easing: theme.transitions.easing.easeOut,
      duration: '200ms',
    }),
    
    [theme.breakpoints.up('md')]: {
      maxWidth: $fullWidth
        ? `calc(100% - ${$navbarOpen ? drawerWidth : miniDrawerWidth}px) !important`
        : '100% !important',
    },
  })
);

// Content wrapper with max-width for readability
export const ContentWrapper = styled(Box)(({ theme }) => ({
  width: '100%',
  maxWidth: 1400,
  margin: '0 auto',
  padding: theme.spacing(3),
}));

// Page header section
export const PageHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  flexDirection: 'column',
  gap: theme.spacing(1),
  marginBottom: theme.spacing(3),
}));

export const PageTitle = styled('h1')(({ theme }) => ({
  margin: 0,
  fontSize: '1.5rem',
  fontWeight: 600,
  lineHeight: 1.25,
  letterSpacing: '-0.01em',
  color: theme.palette.text.primary,
}));

export const PageDescription = styled('p')(({ theme }) => ({
  margin: 0,
  fontSize: '0.875rem',
  lineHeight: 1.5,
  color: theme.palette.text.secondary,
}));

// Section container
export const Section = styled(Box)(({ theme }) => ({
  marginBottom: theme.spacing(4),
}));

export const SectionHeader = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  marginBottom: theme.spacing(2),
  gap: theme.spacing(2),
}));

export const SectionTitle = styled('h2')(({ theme }) => ({
  margin: 0,
  fontSize: '1rem',
  fontWeight: 600,
  lineHeight: 1.4,
  color: theme.palette.text.primary,
}));

// Card grid layout
export const CardGrid = styled(Box)(({ theme }) => ({
  display: 'grid',
  gap: theme.spacing(2),
  gridTemplateColumns: 'repeat(1, 1fr)',
  
  [theme.breakpoints.up('sm')]: {
    gridTemplateColumns: 'repeat(2, 1fr)',
  },
  
  [theme.breakpoints.up('md')]: {
    gridTemplateColumns: 'repeat(3, 1fr)',
  },
  
  [theme.breakpoints.up('lg')]: {
    gridTemplateColumns: 'repeat(4, 1fr)',
  },
}));

// Flex utilities
export const FlexRow = styled(Box)({
  display: 'flex',
  flexDirection: 'row',
  alignItems: 'center',
});

export const FlexColumn = styled(Box)({
  display: 'flex',
  flexDirection: 'column',
});

export const FlexBetween = styled(Box)({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
});

export const FlexCenter = styled(Box)({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
});
