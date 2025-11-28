import { NextFont } from 'next/dist/compiled/@next/font';
import { ThemeOptions } from '@mui/material';
import colors from '@/themes/colors';
import { listBase, listPaper, menuItemRoot } from '@/themes/theme';

const { primaryDark, secondaryDark, greyDark, errorDark, warningDark, infoDark, successDark, crimsonDark } =
  colors;

export const getDarkTheme = (font: NextFont): ThemeOptions => ({
  palette: {
    mode: 'dark',
    primary: {
      lighter: primaryDark[90],
      main: primaryDark[50],
      light: primaryDark[95],
      dark: primaryDark[10],
      darker: primaryDark[95],
      contrastText: primaryDark[0],
      ...primaryDark,
    },
    secondary: {
      main: greyDark[10],
      light: greyDark[95],
      dark: greyDark[10],
      ...secondaryDark,
    },
    error: {
      main: errorDark[50],
      light: errorDark[95],
      dark: errorDark[10],
      ...errorDark,
    },
    warning: {
      main: warningDark[60],
      light: warningDark[90],
      dark: warningDark[10],
      ...warningDark,
    },
    info: {
      main: infoDark[60],
      light: infoDark[90],
      dark: infoDark[10],
      ...infoDark,
    },
    success: {
      main: successDark[50],
      light: successDark[95],
      dark: successDark[10],
      ...successDark,
    },
    grey: { ...greyDark },
    crimson: { ...crimsonDark },
    background: {
      default: greyDark[100],
      paper: greyDark[100],
    },
    divider: greyDark[60],
    action: {
      disabledBackground: '#70707046',
      disabled: '#626262',
    },
    text: {
      primary: greyDark[10],
      secondary: greyDark[30],
      disabled: greyDark[50],
    },
    // neutral: { // TODO: Extend ThemeOptions type to include custom colors
    //   main: '#424242',
    //   contrastText: '#fff',
    // },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          textTransform: 'none',
          borderRadius: '6px',
          fontSize: '12px',
        },
        sizeMedium: {
          height: 34,
          padding: '10px 12px',
        },
      },
    },
    MuiMenu: {
      styleOverrides: {
        paper: ({ theme }) => listPaper(theme),
        list: ({ theme }) => listBase(theme),
      },
    },
    MuiMenuItem: {
      styleOverrides: {
        root: ({ theme }) => menuItemRoot(theme),
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          backgroundColor: greyDark[100],
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: greyDark[100],
        },
      },
    },
  },
  typography: {
    htmlFontSize: 15,
    allVariants: {
      fontFamily: font.style.fontFamily,
      letterSpacing: '0 !important',
      fontWeight: 400,
    },
    h1: { fontSize: 90, fontWeight: 700 },
    h2: { fontSize: 54, fontWeight: 700 },
    h3: { fontSize: 45, fontWeight: 700 },
    h4: { fontSize: 32, fontWeight: 700 },
    h5: { fontSize: 20, fontWeight: 700 },
    h6: { fontSize: 16, fontWeight: 600 },
    body1: { fontSize: 15, fontWeight: 600 },
    body2: { fontSize: 13, fontWeight: 600 },
    button: { fontSize: 15 },
    caption: { fontSize: 12, fontWeight: 500 },
  },
});
