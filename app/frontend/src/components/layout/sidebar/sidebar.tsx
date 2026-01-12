'use client';
import React, { useContext, useCallback, useMemo, FC, ReactNode } from 'react';
import { Badge, Box, Divider, Tooltip } from '@mui/material';
import {
  GridViewRounded,
  CloudQueue,
  VpnKeyRounded,
} from '@mui/icons-material';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { AppContext } from '@/contexts';
import { NAVBAR_OPEN } from '@/contexts/models';
import { Utils } from '@/lib/utils';
import { setCookieNavbarOpen } from '@/lib/cookie-utils';
import {
  StyledDrawer,
  StyledList,
  StyledListItemButton,
  StyledListItemIcon,
  StyledListItemText,
  StyledBottomSection,
  StyledSectionLabel,
} from '@/components/layout/sidebar/styled';
import { CustomTooltip } from '@/components/shared/custom-tooltip';
import { Icon, ICON_NAMES } from '@/components/shared/icon';

export interface Route {
  name: string;
  path: string;
  icon: ReactNode;
  activePaths?: string[];
  badge?: number | string;
}

export interface MenuSection {
  label?: string;
  routes: Route[];
}

const MENU_SECTIONS: MenuSection[] = [
  {
    // Main section - no label for cleaner look at top
    routes: [
      {
        name: 'Dashboard',
        path: '/dashboard',
        icon: <GridViewRounded fontSize="small" />,
        activePaths: ['/', '/dashboard'],
      },
      {
        name: 'Cloud Resources',
        path: '/cloud-resources',
        icon: <CloudQueue fontSize="small" />,
      },
      {
        name: 'Credentials',
        path: '/credentials',
        icon: <VpnKeyRounded fontSize="small" />,
      },
    ],
  },
  // Example of future sections (commented out for now):
  // {
  //   label: 'SETTINGS',
  //   routes: [
  //     { name: 'Preferences', path: '/settings', icon: <SettingsRounded fontSize="small" /> },
  //   ],
  // },
];

interface ILeftNavItem {
  route: Route;
  navbarOpen: boolean;
}

const LeftNavItem: FC<ILeftNavItem> = ({ route, navbarOpen }) => {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const { push: pushRoute } = useRouter();

  const handleListItemClick = (path: string) => {
    pushRoute(`${path}?${searchParams.toString()}`);
  };

  const isActive = useMemo(
    () =>
      pathname === route.path ||
      (route?.activePaths?.length && route.activePaths.includes(pathname)),
    [pathname, route]
  );

  const iconElement = route.badge ? (
    <Badge
      badgeContent={route.badge}
      color="primary"
      sx={{
        '& .MuiBadge-badge': {
          fontSize: '0.625rem',
          minWidth: 16,
          height: 16,
          padding: '0 4px',
        },
      }}
    >
      {route.icon}
    </Badge>
  ) : (
    route.icon
  );

  return (
    <StyledListItemButton
      onClick={() => handleListItemClick(route.path)}
      selected={isActive}
      open={navbarOpen}
    >
      <CustomTooltip title={!navbarOpen ? route?.name : ''} placement="right">
        <StyledListItemIcon>{iconElement}</StyledListItemIcon>
      </CustomTooltip>
      <StyledListItemText primary={route.name} open={navbarOpen} />
    </StyledListItemButton>
  );
};

export const Sidebar = () => {
  const { navbarOpen, setNavbarOpen } = useContext(AppContext);

  const handleNavbarToggle = useCallback(() => {
    const newValue = !navbarOpen;
    Utils.setStorage(NAVBAR_OPEN, newValue);
    setCookieNavbarOpen(newValue);
    setNavbarOpen(newValue);
  }, [navbarOpen, setNavbarOpen]);

  return (
    <StyledDrawer variant="permanent" open={navbarOpen}>
      <StyledList disablePadding>
        {MENU_SECTIONS.map((section, sectionIndex) => (
          <Box key={`section-${sectionIndex}`} sx={{ width: '100%' }}>
            {section.label && (
              <StyledSectionLabel open={navbarOpen}>
                {section.label}
              </StyledSectionLabel>
            )}
            {section.routes.map((route, routeIndex) => (
              <LeftNavItem
                key={`${sectionIndex}-${routeIndex}-${route.name}`}
                route={route}
                navbarOpen={navbarOpen}
              />
            ))}
          </Box>
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
