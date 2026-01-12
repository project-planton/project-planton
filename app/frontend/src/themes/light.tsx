import { NextFont } from 'next/dist/compiled/@next/font';
import { alpha, buttonClasses, tabClasses, ThemeOptions } from '@mui/material';
import { UnfoldMore } from '@mui/icons-material';
import { lightColors, typography, shadows, borderRadius, transitions } from '@/themes/tokens';

// Legacy color imports for backwards compatibility
import colors from '@/themes/colors';
import { formControlLabelLabel, listBase, listPaper, menuItemRoot } from '@/themes/theme';

const {
  primaryLight,
  secondaryLight,
  greyLight,
  errorLight,
  warningLight,
  infoLight,
  successLight,
  exceptionsLight,
  crimsonLight,
} = colors;

export const getLightTheme = (font: NextFont): ThemeOptions => ({
  palette: {
    mode: 'light',
    primary: {
      lighter: lightColors.accent.subtle,
      main: lightColors.accent.primary,
      light: lightColors.accent.hover,
      dark: lightColors.accent.pressed,
      darker: lightColors.accent.pressed,
      contrastText: lightColors.text.inverse,
      ...primaryLight, // Keep legacy for backwards compatibility
    },
    secondary: {
      main: lightColors.text.primary,
      light: lightColors.text.secondary,
      dark: lightColors.text.muted,
      ...secondaryLight,
    },
    error: {
      main: lightColors.semantic.error,
      light: lightColors.semantic.errorSubtle,
      dark: lightColors.semantic.error,
      ...errorLight,
    },
    warning: {
      main: lightColors.semantic.warning,
      light: lightColors.semantic.warningSubtle,
      dark: lightColors.semantic.warning,
      ...warningLight,
    },
    info: {
      main: lightColors.semantic.info,
      light: lightColors.semantic.infoSubtle,
      dark: lightColors.semantic.info,
      ...infoLight,
    },
    success: {
      main: lightColors.semantic.success,
      light: lightColors.semantic.successSubtle,
      dark: lightColors.semantic.success,
      ...successLight,
    },
    grey: { ...greyLight },
    exceptions: { ...exceptionsLight },
    crimson: { ...crimsonLight },
    background: {
      default: lightColors.background.base,
      paper: lightColors.background.raised,
    },
    divider: lightColors.border.default,
    action: {
      disabledBackground: lightColors.interactive.active,
      disabled: lightColors.text.disabled,
      hover: lightColors.interactive.hover,
      selected: lightColors.interactive.selected,
    },
    text: {
      primary: lightColors.text.primary,
      secondary: lightColors.text.secondary,
      disabled: lightColors.text.disabled,
      link: lightColors.accent.primary,
    },
    neutral: {
      main: lightColors.background.overlay,
      contrastText: lightColors.text.primary,
    },
  },
  shape: {
    borderRadius: 8,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          scrollbarColor: `${lightColors.border.strong} ${lightColors.background.base}`,
          '&::-webkit-scrollbar, & *::-webkit-scrollbar': {
            width: 8,
            height: 8,
          },
          '&::-webkit-scrollbar-thumb, & *::-webkit-scrollbar-thumb': {
            borderRadius: 8,
            backgroundColor: lightColors.border.strong,
          },
          '&::-webkit-scrollbar-track, & *::-webkit-scrollbar-track': {
            backgroundColor: lightColors.background.subtle,
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
            backgroundColor: lightColors.background.overlay,
            color: lightColors.text.primary,
          },
        }),
        standardWarning: () => ({
          backgroundColor: lightColors.semantic.warningMuted,
          color: lightColors.semantic.warning,
          border: `1px solid ${alpha(lightColors.semantic.warning, 0.3)}`,
        }),
        standardInfo: () => ({
          backgroundColor: lightColors.semantic.infoMuted,
          color: lightColors.semantic.info,
          border: `1px solid ${alpha(lightColors.semantic.info, 0.3)}`,
        }),
        standardError: () => ({
          backgroundColor: lightColors.semantic.errorMuted,
          color: lightColors.semantic.error,
          border: `1px solid ${alpha(lightColors.semantic.error, 0.3)}`,
        }),
        standardSuccess: () => ({
          backgroundColor: lightColors.semantic.successMuted,
          color: lightColors.semantic.success,
          border: `1px solid ${alpha(lightColors.semantic.success, 0.3)}`,
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
          background: `linear-gradient(135deg, ${lightColors.accent.primary} 0%, ${lightColors.accent.pressed} 100%)`,
          color: lightColors.text.inverse,
          boxShadow: `0 1px 2px rgba(0,0,0,0.1), 0 0 0 1px ${alpha(lightColors.accent.primary, 0.1)}`,
          '&:hover': {
            background: `linear-gradient(135deg, ${lightColors.accent.hover} 0%, ${lightColors.accent.primary} 100%)`,
            boxShadow: shadows.glow.accent,
          },
        }),
        containedSecondary: () => ({
          color: lightColors.text.primary,
          backgroundColor: lightColors.background.base,
          border: `1px solid ${lightColors.border.default}`,
          '&:hover': {
            backgroundColor: lightColors.background.overlay,
            boxShadow: 'none',
          },
        }),
        outlinedPrimary: () => ({
          border: `1px solid ${lightColors.accent.primary}`,
          background: 'transparent',
          color: lightColors.accent.primary,
          '&:hover': {
            background: lightColors.accent.muted,
            borderColor: lightColors.accent.hover,
          },
        }),
        outlinedSecondary: () => ({
          background: 'transparent',
          border: `1px solid ${lightColors.border.default}`,
          color: lightColors.text.secondary,
          '&:hover': {
            backgroundColor: lightColors.interactive.hover,
            borderColor: lightColors.border.strong,
          },
        }),
        textPrimary: () => ({
          color: lightColors.accent.primary,
          '&:hover': {
            background: lightColors.accent.muted,
          },
        }),
        textSecondary: () => ({
          color: lightColors.text.secondary,
          '&:hover': {
            background: lightColors.interactive.hover,
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
            backgroundColor: lightColors.interactive.hover,
          },
          '&:focus-visible': {
            outline: 'none',
            boxShadow: shadows.focus,
          },
        },
      },
    },
    MuiMenuList: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.elevated,
          color: lightColors.text.primary,
        },
      },
    },
    MuiMenu: {
      styleOverrides: {
        paper: ({ theme }) => ({
          ...listPaper(theme),
          backgroundColor: lightColors.background.elevated,
          border: `1px solid ${lightColors.border.default}`,
          boxShadow: shadows.lg,
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
            backgroundColor: lightColors.interactive.hover,
          },
          '&.Mui-selected': {
            backgroundColor: lightColors.interactive.selected,
            '&:hover': {
              backgroundColor: lightColors.interactive.selected,
            },
          },
        }),
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          backgroundColor: alpha(lightColors.background.base, 0.9),
          backdropFilter: 'blur(12px)',
          borderBottom: `1px solid ${lightColors.border.default}`,
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.base,
          border: `1px solid ${lightColors.border.default}`,
          borderRadius: borderRadius.xl,
          boxShadow: 'none',
          transition: transitions.all.normal,
          '&:hover': {
            borderColor: lightColors.border.strong,
            boxShadow: shadows.md,
          },
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.base,
          backgroundImage: 'none',
        },
        elevation1: {
          boxShadow: shadows.sm,
        },
        elevation2: {
          boxShadow: shadows.md,
        },
        elevation3: {
          boxShadow: shadows.lg,
        },
      },
    },
    MuiFormControl: {
      styleOverrides: {
        root: ({ theme }) => ({
          '& label, & label.Mui-disabled': {
            color: lightColors.text.secondary,
            fontSize: typography.fontSize.sm,
          },
          '& label.MuiFormLabel-filled, & label.Mui-focused': {
            color: lightColors.text.primary,
          },
          '& .MuiInputLabel-asterisk': {
            color: lightColors.semantic.error,
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
          backgroundColor: lightColors.background.base,
          borderRadius: `${borderRadius.lg} !important`,
          fontSize: typography.fontSize.sm,
          transition: transitions.all.fast,
          '& fieldset': {
            borderColor: lightColors.border.default,
            transition: transitions.all.fast,
          },
          '&:hover:not(.Mui-disabled):not(.Mui-focused):not(.Mui-error)': {
            '& fieldset': {
              borderColor: lightColors.border.strong,
            },
          },
          '&.Mui-focused': {
            '&.Mui-error': {
              '& fieldset': {
                boxShadow: shadows.focusError,
                borderColor: lightColors.semantic.error,
              },
            },
            '&:not(.Mui-error)': {
              '& fieldset': {
                boxShadow: shadows.focus,
                borderColor: lightColors.accent.primary,
              },
            },
          },
          '& .MuiInputBase-input': {
            fontSize: typography.fontSize.sm,
            fontWeight: typography.fontWeight.normal,
            padding: '8px 12px',
            '&::placeholder': {
              color: lightColors.text.muted,
              opacity: 1,
            },
          },
          '& .MuiSelect-iconOutlined': {
            color: lightColors.text.secondary,
          },
          '& .MuiSvgIcon-root': {
            color: lightColors.text.secondary,
          },
        }),
      },
    },
    MuiInputLabel: {
      styleOverrides: {
        root: () => ({
          fontSize: typography.fontSize.sm,
          fontWeight: typography.fontWeight.normal,
          color: lightColors.text.secondary,
        }),
      },
    },
    MuiFormHelperText: {
      styleOverrides: {
        root: {
          color: lightColors.text.muted,
          fontSize: typography.fontSize.xs,
          marginTop: '6px',
          marginLeft: 0,
        },
      },
    },
    MuiPopover: {
      styleOverrides: {
        paper: {
          backgroundColor: lightColors.background.elevated,
          border: `1px solid ${lightColors.border.default}`,
          boxShadow: shadows.lg,
        },
      },
    },
    MuiDialog: {
      styleOverrides: {
        paper: {
          backgroundColor: lightColors.background.elevated,
          border: `1px solid ${lightColors.border.default}`,
          borderRadius: borderRadius['2xl'],
          boxShadow: shadows['2xl'],
        },
      },
    },
    MuiBackdrop: {
      styleOverrides: {
        root: {
          backgroundColor: alpha(lightColors.text.primary, 0.5),
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
          color: lightColors.background.base,
          '&.Mui-checked': {
            color: lightColors.background.base,
            transform: 'translateX(20px)',
            '& + .MuiSwitch-track': {
              backgroundColor: lightColors.accent.primary,
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
          backgroundColor: lightColors.border.strong,
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
          backgroundColor: lightColors.background.overlay,
          color: lightColors.text.primary,
          '&:hover': {
            backgroundColor: lightColors.border.default,
          },
        }),
        outlined: () => ({
          borderColor: lightColors.border.default,
          '&:hover': {
            backgroundColor: lightColors.interactive.hover,
          },
        }),
      },
    },
    MuiDrawer: {
      styleOverrides: {
        paper: {
          backgroundColor: lightColors.background.base,
          borderRight: `1px solid ${lightColors.border.default}`,
        },
      },
    },
    MuiAccordion: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.base,
          border: `1px solid ${lightColors.border.default}`,
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
          color: lightColors.text.secondary,
        },
      },
    },
    MuiAccordionDetails: {
      styleOverrides: {
        root: {
          borderTop: `1px solid ${lightColors.border.default}`,
          padding: '16px',
        },
      },
    },
    MuiBreadcrumbs: {
      styleOverrides: {
        separator: () => ({
          color: lightColors.text.muted,
        }),
        li: () => ({
          color: lightColors.text.secondary,
          '& a': {
            color: lightColors.text.secondary,
            textDecoration: 'none',
            transition: transitions.all.fast,
            '&:hover': {
              color: lightColors.text.primary,
            },
          },
        }),
      },
    },
    MuiInputAdornment: {
      styleOverrides: {
        root: {
          color: lightColors.text.muted,
          '& .MuiTypography-root': {
            color: lightColors.text.muted,
          },
        },
      },
    },
    MuiDivider: {
      styleOverrides: {
        root: () => ({
          borderColor: lightColors.border.default,
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
          backgroundColor: lightColors.background.overlay,
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
          color: lightColors.text.secondary,
          cursor: 'pointer',
          fontSize: typography.fontSize.sm,
          fontWeight: typography.fontWeight.medium,
          padding: theme.spacing(0.75, 1.5),
          borderRadius: borderRadius.md,
          textTransform: 'none',
          transition: transitions.all.fast,
          [`& .${tabClasses.icon}`]: {
            marginBottom: 0,
            color: lightColors.text.muted,
          },
          [`&.${tabClasses.selected}`]: {
            backgroundColor: lightColors.background.base,
            color: lightColors.text.primary,
            boxShadow: shadows.sm,
            [`& .${tabClasses.icon}`]: {
              color: lightColors.text.primary,
            },
          },
          '&:hover:not(.Mui-selected)': {
            backgroundColor: lightColors.interactive.hover,
            color: lightColors.text.primary,
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
          backgroundColor: lightColors.text.primary,
          color: lightColors.text.inverse,
          borderRadius: borderRadius.md,
          fontSize: typography.fontSize.xs,
          padding: '6px 10px',
          boxShadow: shadows.md,
        },
        arrow: {
          color: lightColors.text.primary,
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
          color: lightColors.border.strong,
          transition: transitions.all.fast,
          '&:hover': {
            backgroundColor: lightColors.interactive.hover,
          },
          '&.Mui-checked': {
            color: lightColors.accent.primary,
          },
        }),
      },
    },
    MuiRadio: {
      styleOverrides: {
        root: () => ({
          color: lightColors.border.strong,
          '&.Mui-checked': {
            color: lightColors.accent.primary,
          },
        }),
      },
    },
    MuiAutocomplete: {
      styleOverrides: {
        paper: ({ theme }) => ({
          ...listPaper(theme),
          backgroundColor: lightColors.background.elevated,
          border: `1px solid ${lightColors.border.default}`,
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
          backgroundColor: lightColors.background.overlay,
          '& .MuiChip-label': {
            paddingLeft: 0,
          },
          '& .MuiSvgIcon-root': {
            fontSize: 16,
            color: lightColors.text.muted,
          },
        }),
        clearIndicator: () => ({
          color: lightColors.text.secondary,
        }),
        popupIndicator: () => ({
          color: lightColors.text.secondary,
        }),
      },
      defaultProps: {
        popupIcon: <UnfoldMore sx={{ fontSize: '20px' }} />,
      },
    },
    MuiLink: {
      styleOverrides: {
        root: () => ({
          color: lightColors.accent.primary,
          textDecorationColor: 'transparent',
          transition: transitions.all.fast,
          '&:hover': {
            textDecorationColor: lightColors.accent.primary,
          },
        }),
      },
    },
    MuiToggleButton: {
      styleOverrides: {
        root: ({ ownerState }) => ({
          padding: '8px 12px',
          borderRadius: borderRadius.lg,
          color: lightColors.text.secondary,
          borderColor: lightColors.border.default,
          transition: transitions.all.fast,
          ...(ownerState.selected && {
            backgroundColor: `${lightColors.interactive.selected} !important`,
            color: lightColors.accent.primary,
            borderColor: lightColors.accent.primary,
          }),
        }),
      },
    },
    MuiSkeleton: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.overlay,
        },
      },
    },
    MuiLinearProgress: {
      styleOverrides: {
        root: {
          backgroundColor: lightColors.background.overlay,
          borderRadius: borderRadius.full,
        },
        bar: {
          backgroundColor: lightColors.accent.primary,
          borderRadius: borderRadius.full,
        },
      },
    },
    MuiCircularProgress: {
      styleOverrides: {
        root: {
          color: lightColors.accent.primary,
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
      color: lightColors.text.muted,
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
