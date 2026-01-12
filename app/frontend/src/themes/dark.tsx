import { NextFont } from 'next/dist/compiled/@next/font';
import { alpha, buttonClasses, tabClasses, ThemeOptions } from '@mui/material';
import { UnfoldMore } from '@mui/icons-material';
import { darkColors, typography, shadows, borderRadius, transitions } from '@/themes/tokens';

// Legacy color imports for backwards compatibility
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
    mode: 'dark',
    primary: {
      lighter: darkColors.accent.subtle,
      main: darkColors.accent.primary,
      light: darkColors.accent.hover,
      dark: darkColors.accent.pressed,
      darker: darkColors.accent.pressed,
      contrastText: darkColors.text.primary,
      ...primaryDark, // Keep legacy for backwards compatibility
    },
    secondary: {
      main: darkColors.text.primary,
      light: darkColors.text.secondary,
      dark: darkColors.text.muted,
      ...secondaryDark,
    },
    error: {
      main: darkColors.semantic.error,
      light: darkColors.semantic.errorSubtle,
      dark: darkColors.semantic.error,
      ...errorDark,
    },
    warning: {
      main: darkColors.semantic.warning,
      light: darkColors.semantic.warningSubtle,
      dark: darkColors.semantic.warning,
      ...warningDark,
    },
    info: {
      main: darkColors.semantic.info,
      light: darkColors.semantic.infoSubtle,
      dark: darkColors.semantic.info,
      ...infoDark,
    },
    success: {
      main: darkColors.semantic.success,
      light: darkColors.semantic.successSubtle,
      dark: darkColors.semantic.success,
      ...successDark,
    },
    grey: { ...greyDark },
    exceptions: { ...exceptionsDark },
    crimson: { ...crimsonDark },
    background: {
      default: darkColors.background.base,
      paper: darkColors.background.raised,
    },
    divider: darkColors.border.default,
    action: {
      disabledBackground: darkColors.interactive.active,
      disabled: darkColors.text.disabled,
      hover: darkColors.interactive.hover,
      selected: darkColors.interactive.selected,
    },
    text: {
      primary: darkColors.text.primary,
      secondary: darkColors.text.secondary,
      disabled: darkColors.text.disabled,
      link: darkColors.accent.primary,
    },
    neutral: {
      main: darkColors.background.overlay,
      contrastText: darkColors.text.primary,
    },
  },
  shape: {
    borderRadius: 8,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          scrollbarColor: `${darkColors.border.strong} ${darkColors.background.base}`,
          '&::-webkit-scrollbar, & *::-webkit-scrollbar': {
            width: 8,
            height: 8,
          },
          '&::-webkit-scrollbar-thumb, & *::-webkit-scrollbar-thumb': {
            borderRadius: 8,
            backgroundColor: darkColors.border.strong,
          },
          '&::-webkit-scrollbar-track, & *::-webkit-scrollbar-track': {
            backgroundColor: darkColors.background.base,
          },
        },
      },
    },
    MuiAlert: {
      styleOverrides: {
        root: ({ theme }) => ({
          padding: theme.spacing(1.5),
          borderRadius: borderRadius.lg,
          '& .MuiAlert-icon': {
            opacity: 1,
          },
          '&.MuiAlert-colorNeutral': {
            backgroundColor: darkColors.background.elevated,
            color: darkColors.text.primary,
          },
        }),
        standardWarning: () => ({
          backgroundColor: darkColors.semantic.warningMuted,
          color: darkColors.semantic.warning,
          border: `1px solid ${alpha(darkColors.semantic.warning, 0.3)}`,
        }),
        standardInfo: () => ({
          backgroundColor: darkColors.semantic.infoMuted,
          color: darkColors.semantic.info,
          border: `1px solid ${alpha(darkColors.semantic.info, 0.3)}`,
        }),
        standardError: () => ({
          backgroundColor: darkColors.semantic.errorMuted,
          color: darkColors.semantic.error,
          border: `1px solid ${alpha(darkColors.semantic.error, 0.3)}`,
        }),
        standardSuccess: () => ({
          backgroundColor: darkColors.semantic.successMuted,
          color: darkColors.semantic.success,
          border: `1px solid ${alpha(darkColors.semantic.success, 0.3)}`,
        }),
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          textTransform: 'none',
          borderRadius: borderRadius.lg,
          fontSize: typography.fontSize.sm,
          fontWeight: typography.fontWeight.medium,
          transition: transitions.all.fast,
          '&:focus-visible': {
            outline: 'none',
            boxShadow: shadows.focus,
          },
        },
        containedPrimary: () => ({
          background: `linear-gradient(135deg, ${darkColors.accent.primary} 0%, ${darkColors.accent.pressed} 100%)`,
          color: darkColors.text.primary,
          boxShadow: `0 1px 2px rgba(0,0,0,0.2), 0 0 0 1px ${alpha(darkColors.accent.primary, 0.2)}`,
          '&:hover': {
            background: `linear-gradient(135deg, ${darkColors.accent.hover} 0%, ${darkColors.accent.primary} 100%)`,
            boxShadow: shadows.glow.accent,
          },
        }),
        containedSecondary: () => ({
          color: darkColors.text.primary,
          backgroundColor: darkColors.background.elevated,
          border: `1px solid ${darkColors.border.default}`,
          '&:hover': {
            backgroundColor: darkColors.background.overlay,
            boxShadow: 'none',
          },
        }),
        outlinedPrimary: () => ({
          border: `1px solid ${darkColors.accent.primary}`,
          background: 'transparent',
          color: darkColors.accent.primary,
          '&:hover': {
            background: darkColors.accent.muted,
            borderColor: darkColors.accent.hover,
          },
        }),
        outlinedSecondary: () => ({
          background: 'transparent',
          border: `1px solid ${darkColors.border.default}`,
          color: darkColors.text.secondary,
          '&:hover': {
            backgroundColor: darkColors.interactive.hover,
            borderColor: darkColors.border.strong,
          },
        }),
        textPrimary: () => ({
          color: darkColors.accent.primary,
          '&:hover': {
            background: darkColors.accent.muted,
          },
        }),
        textSecondary: () => ({
          color: darkColors.text.secondary,
          '&:hover': {
            background: darkColors.interactive.hover,
          },
        }),
        sizeMedium: {
          height: 36,
          padding: '0 16px',
        },
        sizeSmall: {
          height: 32,
          padding: '0 12px',
          fontSize: typography.fontSize.xs,
        },
        sizeLarge: {
          height: 40,
          padding: '0 20px',
          fontSize: typography.fontSize.base,
        },
        startIcon: {
          marginRight: '6px',
        },
        endIcon: {
          marginLeft: '6px',
        },
      },
    },
    MuiIconButton: {
      styleOverrides: {
        root: {
          borderRadius: borderRadius.lg,
          transition: transitions.all.fast,
          '&:hover': {
            backgroundColor: darkColors.interactive.hover,
          },
          '&:focus-visible': {
            outline: 'none',
            boxShadow: shadows.focus,
          },
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
    MuiMenuList: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.elevated,
          color: darkColors.text.primary,
        },
      },
    },
    MuiMenu: {
      styleOverrides: {
        paper: ({ theme }) => ({
          ...listPaper(theme),
          backgroundColor: darkColors.background.elevated,
          border: `1px solid ${darkColors.border.default}`,
          boxShadow: shadows.dark.lg,
          backdropFilter: 'blur(8px)',
        }),
        list: ({ theme }) => listBase(theme),
      },
    },
    MuiMenuItem: {
      styleOverrides: {
        root: ({ theme }) => ({
          ...menuItemRoot(theme),
          transition: transitions.all.fast,
          '&:hover': {
            backgroundColor: darkColors.interactive.hover,
          },
          '&.Mui-selected': {
            backgroundColor: darkColors.interactive.selected,
            '&:hover': {
              backgroundColor: darkColors.interactive.selected,
            },
          },
        }),
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          backgroundColor: alpha(darkColors.background.base, 0.8),
          backdropFilter: 'blur(12px)',
          borderBottom: `1px solid ${darkColors.border.default}`,
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.raised,
          border: `1px solid ${darkColors.border.default}`,
          borderRadius: borderRadius.xl,
          boxShadow: 'none',
          transition: transitions.all.normal,
          '&:hover': {
            borderColor: darkColors.border.strong,
            boxShadow: shadows.dark.md,
          },
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.raised,
          backgroundImage: 'none',
        },
        elevation1: {
          boxShadow: shadows.dark.sm,
        },
        elevation2: {
          boxShadow: shadows.dark.md,
        },
        elevation3: {
          boxShadow: shadows.dark.lg,
        },
      },
    },
    MuiFormControl: {
      styleOverrides: {
        root: ({ theme }) => ({
          '& label, & label.Mui-disabled': {
            color: darkColors.text.secondary,
            fontSize: typography.fontSize.sm,
          },
          '& label.MuiFormLabel-filled, & label.Mui-focused': {
            color: darkColors.text.primary,
          },
          '& .MuiInputLabel-asterisk': {
            color: darkColors.semantic.error,
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
          minHeight: 36,
          backgroundColor: darkColors.background.elevated,
          borderRadius: `${borderRadius.lg} !important`,
          fontSize: typography.fontSize.sm,
          transition: transitions.all.fast,
          '& fieldset': {
            borderColor: darkColors.border.default,
            transition: transitions.all.fast,
          },
          '&:hover:not(.Mui-disabled):not(.Mui-focused):not(.Mui-error)': {
            '& fieldset': {
              borderColor: darkColors.border.strong,
            },
          },
          '&.Mui-focused': {
            '&.Mui-error': {
              '& fieldset': {
                boxShadow: shadows.focusError,
                borderColor: darkColors.semantic.error,
              },
            },
            '&:not(.Mui-error)': {
              '& fieldset': {
                boxShadow: shadows.focus,
                borderColor: darkColors.accent.primary,
              },
            },
          },
          '& .MuiInputBase-input': {
            fontSize: typography.fontSize.sm,
            fontWeight: typography.fontWeight.normal,
            padding: '8px 12px',
            '&::placeholder': {
              color: darkColors.text.muted,
              opacity: 1,
            },
          },
          '& .MuiSelect-iconOutlined': {
            color: darkColors.text.secondary,
          },
          '& .MuiSvgIcon-root': {
            color: darkColors.text.secondary,
          },
        }),
      },
    },
    MuiInputLabel: {
      styleOverrides: {
        root: () => ({
          fontSize: typography.fontSize.sm,
          fontWeight: typography.fontWeight.normal,
          color: darkColors.text.secondary,
        }),
      },
    },
    MuiFormHelperText: {
      styleOverrides: {
        root: {
          color: darkColors.text.muted,
          fontSize: typography.fontSize.xs,
          marginTop: '6px',
          marginLeft: 0,
        },
      },
    },
    MuiPopover: {
      styleOverrides: {
        paper: {
          backgroundColor: darkColors.background.elevated,
          border: `1px solid ${darkColors.border.default}`,
          boxShadow: shadows.dark.lg,
        },
      },
    },
    MuiDialog: {
      styleOverrides: {
        paper: {
          backgroundColor: darkColors.background.elevated,
          border: `1px solid ${darkColors.border.default}`,
          borderRadius: borderRadius['2xl'],
          boxShadow: shadows.dark['2xl'],
        },
      },
    },
    MuiBackdrop: {
      styleOverrides: {
        root: {
          backgroundColor: alpha(darkColors.background.base, 0.8),
          backdropFilter: 'blur(4px)',
        },
      },
    },
    MuiSwitch: {
      styleOverrides: {
        root: {
          borderRadius: '18px',
          width: '44px',
          height: '24px',
          padding: 0,
        },
        switchBase: () => ({
          padding: '2px',
          color: darkColors.text.muted,
          '&.Mui-checked': {
            color: darkColors.text.primary,
            transform: 'translateX(20px)',
            '& + .MuiSwitch-track': {
              backgroundColor: darkColors.accent.primary,
              opacity: 1,
            },
          },
        }),
        thumb: {
          width: '20px',
          height: '20px',
          boxShadow: shadows.sm,
        },
        track: {
          borderRadius: 12,
          backgroundColor: darkColors.border.strong,
          opacity: 1,
        },
      },
    },
    MuiChip: {
      styleOverrides: {
        root: {
          borderRadius: borderRadius.md,
          fontWeight: typography.fontWeight.medium,
          fontSize: typography.fontSize.xs,
        },
        filled: () => ({
          backgroundColor: darkColors.background.overlay,
          color: darkColors.text.primary,
          '&:hover': {
            backgroundColor: darkColors.border.strong,
          },
        }),
        outlined: () => ({
          borderColor: darkColors.border.default,
          '&:hover': {
            backgroundColor: darkColors.interactive.hover,
          },
        }),
      },
    },
    MuiDrawer: {
      styleOverrides: {
        paper: {
          backgroundColor: darkColors.background.elevated,
          borderRight: `1px solid ${darkColors.border.default}`,
        },
        root: {
          '& .MuiPaper-root.MuiCard-root': {
            backgroundColor: darkColors.background.elevated,
          },
        },
      },
    },
    MuiAccordion: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.raised,
          border: `1px solid ${darkColors.border.default}`,
          borderRadius: `${borderRadius.lg} !important`,
          boxShadow: 'none',
          '&:before': {
            display: 'none',
          },
          '&.Mui-expanded': {
            margin: 0,
          },
        },
      },
    },
    MuiAccordionSummary: {
      styleOverrides: {
        root: {
          backgroundColor: 'transparent',
          borderRadius: borderRadius.lg,
          minHeight: 48,
          '&.Mui-expanded': {
            minHeight: 48,
          },
        },
        expandIconWrapper: {
          color: darkColors.text.secondary,
        },
      },
    },
    MuiAccordionDetails: {
      styleOverrides: {
        root: {
          borderTop: `1px solid ${darkColors.border.default}`,
          padding: '16px',
        },
      },
    },
    MuiBreadcrumbs: {
      styleOverrides: {
        separator: () => ({
          color: darkColors.text.muted,
        }),
        li: () => ({
          color: darkColors.text.secondary,
          '& a': {
            color: darkColors.text.secondary,
            textDecoration: 'none',
            transition: transitions.all.fast,
            '&:hover': {
              color: darkColors.text.primary,
            },
          },
        }),
      },
    },
    MuiInputAdornment: {
      styleOverrides: {
        root: {
          color: darkColors.text.muted,
          '& .MuiTypography-root': {
            color: darkColors.text.muted,
          },
        },
      },
    },
    MuiDivider: {
      styleOverrides: {
        root: () => ({
          borderColor: darkColors.border.default,
        }),
        vertical: () => ({
          borderWidth: '1px',
          margin: '0 8px',
        }),
      },
    },
    MuiTabs: {
      styleOverrides: {
        root: ({ theme }) => ({
          minHeight: 40,
          width: 'fit-content',
          borderRadius: borderRadius.lg,
          backgroundColor: darkColors.background.subtle,
          padding: theme.spacing(0.5),
          alignItems: 'center',
          '& .MuiTabs-scroller': {
            overflow: 'visible !important',
          },
        }),
        flexContainer: {
          gap: '4px',
        },
        indicator: {
          display: 'none',
        },
      },
    },
    MuiTab: {
      defaultProps: {
        iconPosition: 'start',
      },
      styleOverrides: {
        root: ({ theme }) => ({
          minHeight: 32,
          backgroundColor: 'transparent',
          color: darkColors.text.secondary,
          cursor: 'pointer',
          fontSize: typography.fontSize.sm,
          fontWeight: typography.fontWeight.medium,
          padding: theme.spacing(0.75, 1.5),
          borderRadius: borderRadius.md,
          textTransform: 'none',
          transition: transitions.all.fast,
          [`& .${tabClasses.icon}`]: {
            marginBottom: 0,
            color: darkColors.text.muted,
          },
          [`&.${tabClasses.selected}`]: {
            backgroundColor: darkColors.background.overlay,
            color: darkColors.text.primary,
            [`& .${tabClasses.icon}`]: {
              color: darkColors.text.primary,
            },
          },
          '&:hover:not(.Mui-selected)': {
            backgroundColor: darkColors.interactive.hover,
            color: darkColors.text.primary,
          },
          [`&.${buttonClasses.disabled}`]: {
            opacity: 0.5,
            cursor: 'not-allowed',
          },
        }),
      },
    },
    MuiTooltip: {
      styleOverrides: {
        tooltip: {
          backgroundColor: darkColors.background.elevated,
          color: darkColors.text.primary,
          border: `1px solid ${darkColors.border.default}`,
          borderRadius: borderRadius.md,
          fontSize: typography.fontSize.xs,
          padding: '6px 10px',
          boxShadow: shadows.dark.md,
        },
        arrow: {
          color: darkColors.background.elevated,
          '&:before': {
            border: `1px solid ${darkColors.border.default}`,
          },
        },
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
        fontSizeMedium: {
          fontSize: '20px',
        },
      },
    },
    MuiCheckbox: {
      styleOverrides: {
        root: () => ({
          color: darkColors.border.strong,
          transition: transitions.all.fast,
          '&:hover': {
            backgroundColor: darkColors.interactive.hover,
          },
          '&.Mui-checked': {
            color: darkColors.accent.primary,
          },
        }),
      },
    },
    MuiRadio: {
      styleOverrides: {
        root: () => ({
          color: darkColors.border.strong,
          '&.Mui-checked': {
            color: darkColors.accent.primary,
          },
        }),
      },
    },
    MuiAutocomplete: {
      styleOverrides: {
        paper: ({ theme }) => ({
          ...listPaper(theme),
          backgroundColor: darkColors.background.elevated,
          border: `1px solid ${darkColors.border.default}`,
        }),
        listbox: ({ theme }) => listBase(theme),
        option: ({ theme }) => ({
          ...menuItemRoot(theme),
          minHeight: 'fit-content !important',
        }),
        tag: ({ theme }) => ({
          display: 'flex',
          alignItems: 'center',
          fontWeight: typography.fontWeight.normal,
          fontSize: typography.fontSize.xs,
          borderRadius: borderRadius.md,
          height: 24,
          padding: theme.spacing(0.5, 1),
          backgroundColor: darkColors.background.overlay,
          '& .MuiChip-label': {
            paddingLeft: 0,
          },
          '& .MuiSvgIcon-root': {
            fontSize: 16,
            color: darkColors.text.muted,
          },
        }),
        clearIndicator: () => ({
          color: darkColors.text.secondary,
        }),
        popupIndicator: () => ({
          color: darkColors.text.secondary,
        }),
      },
      defaultProps: {
        popupIcon: <UnfoldMore sx={{ fontSize: '20px' }} />,
      },
    },
    MuiLink: {
      styleOverrides: {
        root: () => ({
          color: darkColors.accent.primary,
          textDecorationColor: 'transparent',
          transition: transitions.all.fast,
          '&:hover': {
            textDecorationColor: darkColors.accent.primary,
          },
        }),
      },
    },
    MuiToggleButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          padding: '8px 12px',
          borderRadius: borderRadius.lg,
          color: darkColors.text.secondary,
          borderColor: darkColors.border.default,
          transition: transitions.all.fast,
          ...(ownerState.selected && {
            backgroundColor: `${darkColors.interactive.selected} !important`,
            color: darkColors.accent.primary,
            borderColor: darkColors.accent.primary,
          }),
        }),
      },
    },
    MuiSkeleton: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.overlay,
        },
      },
    },
    MuiLinearProgress: {
      styleOverrides: {
        root: {
          backgroundColor: darkColors.background.overlay,
          borderRadius: borderRadius.full,
        },
        bar: {
          backgroundColor: darkColors.accent.primary,
          borderRadius: borderRadius.full,
        },
      },
    },
    MuiCircularProgress: {
      styleOverrides: {
        root: {
          color: darkColors.accent.primary,
        },
      },
    },
  },

  typography: {
    htmlFontSize: 16,
    fontFamily: font.style.fontFamily,
    allVariants: {
      fontFamily: font.style.fontFamily,
      letterSpacing: typography.letterSpacing.normal,
    },
    h1: {
      fontSize: typography.fontSize['5xl'],
      fontWeight: typography.fontWeight.bold,
      lineHeight: typography.lineHeight.tight,
      letterSpacing: typography.letterSpacing.tighter,
    },
    h2: {
      fontSize: typography.fontSize['4xl'],
      fontWeight: typography.fontWeight.semibold,
      lineHeight: typography.lineHeight.tight,
      letterSpacing: typography.letterSpacing.tight,
    },
    h3: {
      fontSize: typography.fontSize['3xl'],
      fontWeight: typography.fontWeight.semibold,
      lineHeight: typography.lineHeight.snug,
      letterSpacing: typography.letterSpacing.tight,
    },
    h4: {
      fontSize: typography.fontSize['2xl'],
      fontWeight: typography.fontWeight.semibold,
      lineHeight: typography.lineHeight.snug,
    },
    h5: {
      fontSize: typography.fontSize.xl,
      fontWeight: typography.fontWeight.semibold,
      lineHeight: typography.lineHeight.snug,
    },
    h6: {
      fontSize: typography.fontSize.lg,
      fontWeight: typography.fontWeight.semibold,
      lineHeight: typography.lineHeight.normal,
    },
    body1: {
      fontSize: typography.fontSize.md,
      fontWeight: typography.fontWeight.normal,
      lineHeight: typography.lineHeight.relaxed,
    },
    body2: {
      fontSize: typography.fontSize.sm,
      fontWeight: typography.fontWeight.normal,
      lineHeight: typography.lineHeight.normal,
    },
    button: {
      fontSize: typography.fontSize.sm,
      fontWeight: typography.fontWeight.medium,
      textTransform: 'none',
    },
    caption: {
      fontSize: typography.fontSize.xs,
      fontWeight: typography.fontWeight.normal,
      lineHeight: typography.lineHeight.normal,
      color: darkColors.text.muted,
    },
    overline: {
      fontSize: typography.fontSize.xs,
      fontWeight: typography.fontWeight.medium,
      letterSpacing: typography.letterSpacing.wider,
      textTransform: 'uppercase',
    },
    subtitle1: {
      fontSize: typography.fontSize.md,
      fontWeight: typography.fontWeight.medium,
      lineHeight: typography.lineHeight.normal,
    },
    subtitle2: {
      fontSize: typography.fontSize.sm,
      fontWeight: typography.fontWeight.medium,
      lineHeight: typography.lineHeight.normal,
    },
  },
});
