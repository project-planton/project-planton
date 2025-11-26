'use client';
import React, { createContext, useEffect, useState, useMemo, useCallback } from 'react';
import { CssBaseline, useMediaQuery } from '@mui/material';
import { ThemeProvider } from '@mui/material/styles';
import { appTheme } from '@/themes/theme';
import {
  AppContextType,
  PCTheme,
  PCThemeType,
  ISnack,
  Severity,
  THEME,
  PCS_THEME_IDENTIFIER,
} from '@/contexts/models';
import { NextFont } from 'next/dist/compiled/@next/font';
import { Utils } from '@/lib/utils';

type Props = {
  connectHost: string;
  font: NextFont;
};

type PropsWithChildren = Props & {
  children: React.ReactNode;
};

const defaultAppContext: AppContextType = {
  connectHost: '',
  title: '',
  theme: {} as PCTheme,
  changeTheme: null,
  openSnackbar: null,
  setTitle: null,
  pageLoading: false,
  setPageLoading: null,
};

export const AppContext = createContext<AppContextType>(defaultAppContext);

export const AppContextProvider = ({ children, connectHost, font }: PropsWithChildren) => {
  const [title, setTitle] = useState('');
  const [snackPack, setSnackPack] = useState<ISnack[]>([]);
  const [snackMsg, setSnackMsg] = useState<ISnack | undefined>(undefined);
  const [open, setOpen] = useState(false);
  const [pageLoading, setPageLoading] = useState(false);
  const [connectRpcHost] = useState(connectHost);
  const isDarkThemePreferred = useMediaQuery('(prefers-color-scheme: dark)');
  const existingMode = Utils.getStorage(PCS_THEME_IDENTIFIER) as PCThemeType | undefined;

  useEffect(() => {
    if (snackPack.length && !snackMsg) {
      setSnackMsg({ ...snackPack[0] });
      setSnackPack((prev) => prev.slice(1));
      setOpen(true);
    } else if (snackPack.length && snackMsg && open) {
      setTimeout(() => {
        setOpen(false);
      }, 2000);
    }
  }, [snackPack, snackMsg, open]);

  const themeMode = useMemo<PCThemeType>(() => {
    if (existingMode === THEME.DARK || existingMode === THEME.LIGHT) {
      return existingMode;
    }
    return isDarkThemePreferred ? THEME.DARK : THEME.LIGHT;
  }, [existingMode, isDarkThemePreferred]);

  const [theme, setTheme] = useState<PCTheme>({
    mode: THEME.LIGHT,
    ...appTheme(THEME.LIGHT, font),
  });

  useEffect(() => {
    setTheme({ mode: themeMode, ...appTheme(themeMode, font) });
  }, [themeMode, font]);

  const openSnackbar = useCallback((message = '', svrty: Severity = 'success') => {
    setSnackPack((prev) => [...prev, { id: message, message, severity: svrty }]);
  }, []);

  const changeTheme = useCallback((type: PCThemeType) => {
    const newTheme = {
      mode: type,
      ...appTheme(type, font),
    };
    setTheme(newTheme);
    Utils.setStorage(PCS_THEME_IDENTIFIER, type);
  }, [font]);

  const handleSetTitle = useCallback((newTitle: string) => {
    setTitle(newTitle);
  }, []);

  const handleSetPageLoading = useCallback((loading: boolean) => {
    setPageLoading(loading);
  }, []);

  const appContext: AppContextType = useMemo(
    () => ({
      title,
      theme,
      changeTheme,
      openSnackbar,
      setTitle: handleSetTitle,
      pageLoading,
      setPageLoading: handleSetPageLoading,
      connectHost: connectRpcHost,
    }),
    [title, theme, changeTheme, openSnackbar, handleSetTitle, pageLoading, handleSetPageLoading, connectRpcHost]
  );

  return (
    <AppContext.Provider value={appContext}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        {children}
      </ThemeProvider>
    </AppContext.Provider>
  );
};

