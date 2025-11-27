import { Theme } from '@mui/material/styles';

export type PCThemeType = 'dark' | 'light';

export const THEME = {
  DARK: 'dark' as PCThemeType,
  LIGHT: 'light' as PCThemeType,
};

export const PCS_THEME_IDENTIFIER = 'pcs-theme';
export const NAVBAR_OPEN = 'navbar-open';

export type PCTheme = Theme & {
  mode: PCThemeType;
};

export type Severity = 'success' | 'error' | 'warning' | 'info';

export interface ISnack {
  id?: string;
  message?: string;
  severity?: Severity;
}

export interface AppContextType {
  connectHost: string;
  title: string;
  theme: PCTheme;
  changeTheme: (type: PCThemeType) => void;
  openSnackbar: (message?: string, severity?: Severity) => void;
  setTitle: (title: string) => void | null;
  pageLoading: boolean;
  setPageLoading: (loading: boolean) => void | null;
  navbarOpen: boolean;
  setNavbarOpen: (open: boolean) => void;
}
