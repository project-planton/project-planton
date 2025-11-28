import { ThemeOptions } from '@mui/material';
import { NextFont } from 'next/dist/compiled/@next/font';
import colors from '@/themes/colors';
import { listBase, listPaper, menuItemRoot } from '@/themes/theme';

const {
  primaryLight,
  secondaryLight,
  greyLight,
  errorLight,
  warningLight,
  infoLight,
  successLight,
  crimsonLight,
} = colors;

export const getLightTheme = (font: NextFont): ThemeOptions => ({
  palette: {
    mode: 'light',
    primary: {
      lighter: primaryLight[90],
      main: primaryLight[50],
      light: primaryLight[95],
      dark: primaryLight[10],
      darker: primaryLight[95],
      contrastText: primaryLight[100],
      ...primaryLight,
    },
    secondary: {
      main: greyLight[10],
      light: greyLight[10],
      dark: greyLight[95],
      ...secondaryLight,
    },
    error: {
      main: errorLight[50],
      light: errorLight[95],
      dark: errorLight[10],
      ...errorLight,
    },
    warning: {
      main: warningLight[50],
      light: warningLight[95],
      dark: warningLight[20],
      ...warningLight,
    },
    info: {
      main: infoLight[50],
      light: infoLight[95],
      dark: infoLight[20],
      ...infoLight,
    },
    success: {
      main: successLight[50],
      light: successLight[95],
      dark: successLight[10],
      ...successLight,
    },
    grey: { ...greyLight },
    // exceptions: { ...exceptionsLight }, // TODO: Extend ThemeOptions type to include custom colors
    crimson: { ...crimsonLight },
    background: {
      default: greyLight[100],
      paper: greyLight[100],
    },
    divider: greyLight[60],
    text: {
      primary: greyLight[10],
      secondary: greyLight[30],
      disabled: greyLight[50],
      // link: primaryLight[50], // TODO: Extend TypeText type to include link property
    },
    // neutral: { // TODO: Extend ThemeOptions type to include custom colors
    //   main: '#e0e0e0',
    //   contrastText: '#000',
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
          backgroundColor: greyLight[100],
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: greyLight[100],
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
    caption: { fontSize: 12, fontWeight: 400 },
  },
});
