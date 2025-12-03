'use client';
import React, { useContext, useMemo } from 'react';
import { Tooltip } from '@mui/material';
import { AppContext } from '@/contexts';
import { PCS_THEME_IDENTIFIER, THEME } from '@/contexts/models';
import { Utils } from '@/lib/utils';
import { StyledThemeButton } from '@/components/layout/theme-switch/styled';
import { Icon, ICON_NAMES } from '@/components/shared/icon';

const ThemeSwitch = () => {
  const {
    theme: { mode },
    changeTheme,
  } = useContext(AppContext);
  const nextTheme = useMemo(() => (mode === THEME.DARK ? THEME.LIGHT : THEME.DARK), [mode]);

  const toggleTheme = () => {
    Utils.setStorage(PCS_THEME_IDENTIFIER, nextTheme);
    changeTheme(nextTheme);
  };

  return (
    <Tooltip title={`Switch to ${nextTheme}`}>
      <StyledThemeButton
        $mode={mode}
        onClick={toggleTheme}
        startIcon={<Icon name={ICON_NAMES.SUN} alt={`Switch to ${nextTheme}`} />}
        endIcon={<Icon name={ICON_NAMES.MOON} alt={`Switch to ${nextTheme}`} />}
        disableRipple
      />
    </Tooltip>
  );
};

export default ThemeSwitch;
