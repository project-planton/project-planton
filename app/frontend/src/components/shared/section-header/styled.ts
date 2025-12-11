import { Stack, styled } from '@mui/material';

export const HeaderContainer = styled(Stack, {
  shouldForwardProp: (prop) => prop !== '$bb',
})<{ $bb?: boolean }>(({ theme, $bb }) => ({
  flexDirection: 'row',
  flexWrap: 'wrap',
  justifyContent: 'space-between',
  alignItems: 'center',
  gap: theme.spacing(1.5),
  width: '100%',
  borderBottom: $bb ? `1px solid ${theme.palette.divider}` : 0,
  backgroundColor: theme.palette.background.default,
}));

export const LeftSection = styled(Stack)(({ theme }) => ({
  flexDirection: 'row',
  justifyContent: 'space-between',
  alignItems: 'center',
  gap: theme.spacing(1),
}));

