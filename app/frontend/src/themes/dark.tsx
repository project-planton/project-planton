import { NextFont } from 'next/dist/compiled/@next/font';
import { alpha, buttonClasses, tabClasses, ThemeOptions } from '@mui/material';
import { UnfoldMore } from '@mui/icons-material';
import colors from '@/themes/colors';
import { formControlLabelLabel, listBase, listPaper, menuItemRoot } from '@/themes/theme';

const {
  primaryDark,
  secondaryDark,
  greyDark,
  errorDark,
  warningDark,
  infoDark,
  successDark,
  exceptionsDark,
  crimsonDark,
} = colors;

export const getDarkTheme = (font: NextFont): ThemeOptions => ({
  palette: {
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
    exceptions: { ...exceptionsDark },
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
      link: greyDark[10],
    },
    neutral: {
      main: '#424242',
      contrastText: '#fff',
    },
  },
  components: {
    MuiAlert: {
      styleOverrides: {
        root: ({ theme }) => ({
          padding: theme.spacing(1.5),
          '& .MuiAlert-message': {
            // fontSize: 15,
            // padding: 0,
          },
          '&.MuiAlert-standardWarning .MuiAlert-icon': {
            color: theme.palette.warning.main,
          },
          '&.MuiAlert-standardInfo .MuiAlert-icon': {
            color: theme.palette.info.main,
          },
          '&.MuiAlert-standardError .MuiAlert-icon': {
            color: theme.palette.error.main,
          },
          '&.MuiAlert-colorNeutral': {
            backgroundColor: theme.palette.neutral.main,
            color: theme.palette.neutral.contrastText,
          },
        }),
        standardWarning: ({ theme }) => ({
          backgroundColor: theme.palette.warning.dark,
          color: theme.palette.warning.light,
        }),
        standardInfo: ({ theme }) => ({
          backgroundColor: theme.palette.info.dark,
          color: theme.palette.info.light,
        }),
        standardError: ({ theme }) => ({
          backgroundColor: theme.palette.error.light,
          color: theme.palette.error.dark,
        }),
        standardSuccess: () => ({
          backgroundColor: '#EBF9ED',
          color: '#2A6841',
        }),
      },
    },
    MuiButton: {
      styleOverrides: {
        containedPrimary: ({ theme }) => ({
          color: theme.palette.exceptions[20],
          backgroundColor: theme.palette.primary[75],
          '&:hover': {
            backgroundColor: theme.palette.primary[75],
            boxShadow: 'none',
          },
        }),
        containedSecondary: ({ theme }) => ({
          color: theme.palette.text.primary,
          backgroundColor: theme.palette.exceptions[20],
          border: `1px solid ${theme.palette.divider}`,
          '&:hover': {
            backgroundColor: theme.palette.exceptions[60],
            boxShadow: 'none',
          },
        }),
        outlinedPrimary: {
          border: `1px solid ${primaryDark[50]}`,
          background: primaryDark[100],
          '&:hover': {
            background: primaryDark[95],
          },
          color: primaryDark[10],
        },
        outlinedSecondary: ({ theme }) => ({
          background: exceptionsDark[20],
          border: `1px solid ${greyDark[60]}`,
          '&:hover': {
            backgroundColor: theme.palette.exceptions[60],
            border: `1px solid ${greyDark[60]}`,
          },
        }),
        textSecondary: {
          background: exceptionsDark[100],
          color: greyDark[30],
          '&:hover': {
            background: greyDark[90],
          },
        },
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
        iconSizeMedium: {
          fontSize: '24px',
        },
        startIcon: {
          marginRight: '4px',
        },
        endIcon: {
          marginLeft: '4px',
        },
      },
    },
    MuiMenuList: {
      styleOverrides: {
        root: {
          backgroundColor: exceptionsDark[100],
          color: greyDark[10],
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
    MuiIconButton: {
      styleOverrides: {
        root: {
          '&.Mui-disabled': {
            '& .MuiSvgIcon-root': {
              'g, path': {
                fill: 'currentColor',
              },
            },
          },
        },
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
    MuiFormControl: {
      styleOverrides: {
        root: ({ theme }) => ({
          '& label, & label.Mui-disabled': {
            color: theme.palette.grey[10],
            top: '-3px',
          },
          '& label.MuiFormLabel-filled, & label.Mui-focused': {
            color: theme.palette.grey[10],
            top: 0,
          },
          '& .MuiInputLabel-asterisk': {
            color: theme.palette.error.main,
            fontSize: '20px',
            lineHeight: '20px',
          },
        }),
      },
    },
    MuiFormControlLabel: {
      styleOverrides: {
        label: () => formControlLabelLabel(),
      },
    },
    MuiInputBase: {
      styleOverrides: {
        root: ({ theme }) => ({
          minHeight: 34,
          backgroundColor: theme.palette.exceptions[100],
          borderRadius: `${theme.shape.borderRadius * 2}px !important`,
          margin: '4px',
          '& fieldset': {
            borderColor: theme.palette.divider,
          },
          '&:hover:not(.Mui-disabled):not(.Mui-focused):not(.Mui-error)': {
            '& fieldset': {
              borderColor: theme.palette.text.primary,
              transition: 'border-color 0.5s ease-in-out',
            },
          },
          '&.Mui-focused': {
            '&.Mui-error': {
              '& fieldset': {
                boxShadow: `0 0 0 3px ${alpha(theme.palette.error.main, 0.2)}`,
                transition: 'box-shadow 0.5s ease-in-out',
              },
            },
            '&:not(.Mui-error)': {
              '& fieldset': {
                boxShadow: `0 0 0 3px ${alpha(theme.palette.primary.main, 0.34)}`,
                borderColor: '#959595 !important',
                transition: 'box-shadow 0.5s ease-in-out',
              },
            },
          },
          '& .MuiInputBase-input': {
            fontSize: 12,
            fontWeight: 400,
            padding: '0px 8px',
            '&.Mui-disabled': {
              color: theme.palette.text.secondary,
            },
          },
          '& .MuiSelect-iconOutlined': {
            color: theme.palette.text.secondary,
          },
          '&.MuiAutocomplete-inputRoot': {
            padding: 0,
          },
          '& .MuiSvgIcon-root': {
            color: theme.palette.secondary[50],
          },
          '& .MuiButtonBase-tag': {
            fontWeight: 400,
            fontSize: 12,
          },
          '& .MuiInputAdornment-root': {
            marginRight: 0,
          },
        }),
      },
    },
    MuiInputLabel: {
      styleOverrides: {
        root: ({ theme }) => ({
          fontSize: 12,
          fontWeight: 400,
          color: theme.palette.text.primary,
        }),
      },
    },
    MuiFormHelperText: {
      styleOverrides: {
        root: {
          color: '#626262',
          fontSize: '13px',
          lineHeight: '17px',
          margin: '5px 0 0 0',
        },
      },
    },
    MuiPopover: {
      styleOverrides: {
        paper: {
          backgroundColor: '#0F0F0F',
          color: '#ffffff',
        },
      },
    },
    MuiSwitch: {
      styleOverrides: {
        root: {
          borderRadius: '18px',
          width: '48px',
          height: '30px',
          margin: '4px 0',
          padding: 'unset',
        },
        switchBase: ({ theme }) => ({
          // Controls default (unchecked) color for the thumb
          padding: '5px',
          color: theme.palette.background.default,
        }),
        colorPrimary: ({ theme }) => ({
          '&.Mui-checked': {
            // Controls checked color for the thumb
            color: theme.palette.background.default,
          },
        }),
        thumb: ({ theme }) => ({
          width: '20px',
          height: '20px',
          boxShadow: 'unset',
          '&.Mui-checked': {
            // Controls checked color for the thumb
            color: theme.palette.background.default,
          },
        }),
        track: ({ theme }) => ({
          // Controls default (unchecked) color for the track
          borderRadius: 26 / 2,
          background: '#272727',
          opacity: 1,
          '.Mui-checked.Mui-checked + &': {
            // Controls checked color for the track
            opacity: 1,
            background: theme.palette.text.primary,
          },
        }),
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: greyDark[100],
        },
      },
    },
    MuiDrawer: {
      styleOverrides: {
        root: {
          '& .MuiPaper-root.MuiCard-root': {
            backgroundColor: exceptionsDark[70],
          },
          '& .MuiAccordionSummary-root': {
            '&.Mui-expanded': {
              minHeight: '48px',
            },
          },
        },
      },
    },
    MuiAccordionSummary: {
      styleOverrides: {
        root: {
          backgroundColor: '#0F0F0F',
          borderRadius: '4px',
        },
        expandIconWrapper: {
          color: greyDark[10],
        },
      },
    },
    MuiAccordionDetails: {
      styleOverrides: {
        root: {
          borderBottom: '1px solid #F0F0F030',
        },
      },
    },
    MuiBreadcrumbs: {
      styleOverrides: {
        separator: ({ theme }) => ({
          color: theme.palette.text.secondary,
        }),
        li: ({ theme }) => ({
          color: theme.palette.text.secondary,
        }),
      },
    },
    MuiInputAdornment: {
      styleOverrides: {
        root: {
          color: '#d6d6d6',
          '& .MuiTypography-root': {
            color: '#d6d6d6',
          },
        },
      },
    },
    MuiDivider: {
      styleOverrides: {
        root: ({ theme }) => ({
          '&:before': {
            borderColor: theme.palette.divider,
          },
          '&:after': {
            borderColor: theme.palette.divider,
          },
        }),
        vertical: ({ theme }) => ({
          borderWidth: '1px',
          margin: '0 10px',
          borderColor: theme.palette.divider,
        }),
      },
    },
    MuiTabs: {
      styleOverrides: {
        root: ({ theme }) => ({
          minHeight: 40,
          width: 'fit-content',
          borderRadius: 0,
          backgroundColor: theme.palette.grey[20],
          padding: theme.spacing(1.5),
          alignItems: 'center',
          '& .MuiTabs-scroller': {
            overflow: 'visible !important',
          },
        }),
        flexContainer: ({ theme }) => ({ gap: theme.spacing(1.5) }),
        indicator: ({ theme }) => ({
          bottom: theme.spacing(-1.5),
        }),
      },
    },
    MuiTab: {
      defaultProps: {
        iconPosition: 'start',
      },
      styleOverrides: {
        root: ({ theme }) => ({
          minHeight: 0,
          backgroundColor: 'transparent',
          color: theme.palette.text.primary,
          cursor: 'pointer',
          fontSize: theme.spacing(1.5),
          fontWeight: 500,
          padding: theme.spacing(1, 1.5),
          borderRadius: theme.shape.borderRadius + 2,
          textTransform: 'none',
          [`& .${tabClasses.icon}`]: {
            marginBottom: 0,
            color: theme.palette.text.secondary,
          },
          [`&.${tabClasses.selected}, &:hover`]: {
            backgroundColor: theme.palette.grey[70],
            color: theme.palette.text.primary,
            [`& .${tabClasses.icon}`]: {
              color: theme.palette.text.primary,
            },
          },
          [`&.${buttonClasses.disabled}`]: {
            opacity: 0.5,
            cursor: 'not-allowed',
          },
        }),
      },
    },
    MuiSelect: {
      defaultProps: { IconComponent: UnfoldMore },
    },
    MuiSvgIcon: {
      styleOverrides: {
        fontSizeSmall: {
          fontSize: '16px',
        },
      },
    },
    MuiCheckbox: {
      styleOverrides: {
        root: ({ theme }) => ({
          '& .MuiSvgIcon-root': {
            height: '1.3rem',
          },
          '&.Mui-checked .MuiSvgIcon-root path': {
            fill: theme.palette.text.primary,
          },
        }),
      },
    },
    MuiAutocomplete: {
      styleOverrides: {
        paper: ({ theme }) => listPaper(theme),
        listbox: ({ theme }) => listBase(theme),
        option: ({ theme }) => ({
          ...menuItemRoot(theme),
          minHeight: 'fit-content !important',
        }),
        tag: ({ theme }) => ({
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          fontWeight: 400,
          fontSize: 12,
          borderRadius: theme.shape.borderRadius,
          height: 24,
          padding: theme.spacing(1, 0, 1, 1),
          backgroundColor: theme.palette.grey[70],
          '& .MuiChip-label': {
            paddingLeft: 0,
          },
          '& .MuiSvgIcon-root': {
            fontSize: 16,
          },
        }),
        clearIndicator: ({ theme }) => ({
          color: theme.palette.text.secondary,
        }),
        popupIndicator: ({ theme }) => ({
          color: theme.palette.text.secondary,
        }),
      },
      defaultProps: {
        popupIcon: <UnfoldMore sx={{ fontSize: '20px' }} />,
      },
    },
    MuiLink: {
      styleOverrides: {
        root: ({ theme }) => ({
          color: theme.palette.text.primary,
          textDecorationColor: theme.palette.text.primary,
        }),
      },
    },
    MuiToggleButton: {
      styleOverrides: {
        root: ({ theme, ownerState }) => ({
          padding: theme.spacing(1, 1.5),
          borderRadius: theme.shape.borderRadius * 2,
          color: theme.palette.text.secondary,
          ...(ownerState.selected && {
            backgroundColor: `${theme.palette.grey[70]} !important`,
            color: theme.palette.text.primary,
          }),
        }),
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
    overline: { fontSize: 14 },
    subtitle1: { fontSize: 15 },
    subtitle2: { fontSize: 12 },
  },
});
