'use client';
import React, { useContext, useCallback, useMemo, FC } from 'react';
import { Box, Divider, Tooltip } from '@mui/material';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { AppContext } from '@/contexts';
import { NAVBAR_OPEN, THEME } from '@/contexts/models';
import { Utils } from '@/lib/utils';
import { setCookieNavbarOpen } from '@/lib/cookie-utils';
import {
  StyledDrawer,
  StyledList,
  StyledListItemButton,
  StyledListItemIcon,
  StyledListItemText,
  StyledBottomSection,
} from '@/components/layout/sidebar/styled';
import { CustomTooltip } from '@/components/shared/custom-tooltip';
import { Icon, ICON_NAMES } from '@/components/shared/icon';

export interface Route {
  name: string;
  path: string;
  iconName: ICON_NAMES;
  activePaths?: string[];
}

const TOP_MENU: Route[] = [
  {
    name: 'Dashboard',
    path: '/dashboard',
    iconName: ICON_NAMES.DASHBOARD,
    activePaths: ['/', '/dashboard'],
  },
  {
    name: 'Cloud Resources',
    path: '/cloud-resources',
    iconName: ICON_NAMES.INFRA_HUB,
  },
  {
    name: 'Streaming',
    path: '/streaming',
    iconName: ICON_NAMES.DASHBOARD,
  },
];

interface ILeftNavItem {
  route: Route;
  navbarOpen: boolean;
}

const LeftNavItem: FC<ILeftNavItem> = ({ route, navbarOpen }) => {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const { push: pushRoute } = useRouter();
  const { theme } = useContext(AppContext);

  const handleListItemClick = (path: string) => {
    pushRoute(`${path}?${searchParams.toString()}`);
  };

  const isActive = useMemo(
    () =>
      pathname === route.path ||
      (route?.activePaths?.length && route.activePaths.includes(pathname)),
    [pathname, route]
  );

  return (
    <StyledListItemButton
      onClick={() => handleListItemClick(route.path)}
      selected={isActive}
      open={navbarOpen}
    >
      <CustomTooltip title={!navbarOpen ? route?.name : ''} placement="right">
        <StyledListItemIcon>
          <Icon
            name={route.iconName}
            fontSize="small"
            sx={{
              filter: theme.mode === THEME.DARK ? 'invert(0.3)' : 'none',
            }}
          />
        </StyledListItemIcon>
      </CustomTooltip>
      <StyledListItemText primary={route.name} open={navbarOpen} />
    </StyledListItemButton>
  );
};

export const Sidebar = () => {
  const { navbarOpen, setNavbarOpen } = useContext(AppContext);

  const routesArray = useMemo(() => {
    return TOP_MENU;
  }, []);

  const handleNavbarToggle = useCallback(() => {
    const newValue = !navbarOpen;
    Utils.setStorage(NAVBAR_OPEN, newValue);
    setCookieNavbarOpen(newValue);
    setNavbarOpen(newValue);
  }, [navbarOpen, setNavbarOpen]);

  return (
    <StyledDrawer variant="permanent" open={navbarOpen}>
      <StyledList disablePadding>
        {routesArray.map((route, index) => (
          <LeftNavItem key={`${index}-${route.name}`} route={route} navbarOpen={navbarOpen} />
        ))}
      </StyledList>
      <StyledBottomSection>
        <Divider />
        <Tooltip
          arrow
          placement="right"
          title={`${navbarOpen ? 'Collapse' : 'Expand'} the navigation`}
        >
          <Box
            sx={{
              alignSelf: 'end',
              display: 'flex',
              cursor: 'pointer',
              padding: (theme) => theme.spacing(1),
              border: '1px solid transparent',
              ...(navbarOpen ? { rotate: '180deg' } : {}),
              borderRadius: (theme) => theme.shape.borderRadius - 2.5,
              '&:hover': {
                border: (theme) => `1px solid ${theme.palette.divider}`,
                backgroundColor: (theme) => theme.palette.background.default,
              },
            }}
            onClick={handleNavbarToggle}
          >
            <Icon
              name={ICON_NAMES.NAV}
              sx={{
                fontSize: 20,
                stroke: (theme) =>
                  navbarOpen ? theme.palette.grey[0] : theme.palette.secondary[50],
                fill: (theme) => (navbarOpen ? theme.palette.secondary[50] : 'none'),
              }}
            />
          </Box>
        </Tooltip>
      </StyledBottomSection>
    </StyledDrawer>
  );
};
