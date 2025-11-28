import { createTheme, CSSObject, Theme } from '@mui/material/styles';
import { getDarkTheme } from './dark';
import { getLightTheme } from './light';
import { NextFont } from 'next/dist/compiled/@next/font';

declare module '@mui/material/styles' {
  interface PaletteColor {
    lighter?: string;
    darker?: string;
  }

  interface SimplePaletteColorOptions {
    lighter?: string;
    darker?: string;
  }

  interface Palette {
    crimson?: {
      0: string;
      10: string;
      20: string;
      30: string;
      40: string;
      50: string;
      60: string;
    };
  }

  interface PaletteOptions {
    crimson?: {
      0: string;
      10: string;
      20: string;
      30: string;
      40: string;
      50: string;
      60: string;
    };
  }
}

export const listPaper = (theme: Theme): CSSObject => ({
  marginTop: theme.spacing(1),
  boxShadow: '1px 1px 6px rgba(0, 0, 0, 0.06)',
  border: `1px solid ${theme.palette.divider}`,
  borderRadius: theme.shape.borderRadius * 3,
  backgroundColor: theme.palette.background.default,
});

export const listBase = (theme: Theme): CSSObject => ({
  display: 'flex',
  flexDirection: 'column',
  gap: theme.spacing(1),
  padding: `${theme.spacing(1)} !important`,
});

export const menuItemRoot = (theme: Theme): CSSObject => {
  const activeState = {
    backgroundColor: `${theme.palette.grey[80]} !important`,
  };

  return {
    padding: `${theme.spacing(1)} !important`,
    borderRadius: theme.shape.borderRadius * 2,
    fontSize: 12,
    fontWeight: 400,
    '&:hover': { ...activeState },
    '&.Mui-focusVisible, &.Mui-focused, &.Mui-selected, &.Mui-active, &[aria-selected="true"]': {
      ...activeState,
    },
  };
};

export const appTheme = (type: 'dark' | 'light', font: NextFont): Theme => {
  const theme = type === 'light' ? getLightTheme(font) : getDarkTheme(font);
  return createTheme(theme);
};

