'use client';
import React, { useContext, useCallback } from 'react';
import { ListItem, Divider } from '@mui/material';
import { AppContext } from '@/contexts';
import {
  Dashboard as DashboardIcon,
  Menu as MenuIcon,
  ChevronLeft,
  Inventory2,
  ShoppingBag,
  People,
  Article,
  Folder,
  Person,
} from '@mui/icons-material';
import { usePathname, useRouter } from 'next/navigation';
import {
  StyledDrawer,
  SidebarContainer,
  StyledBottomSection,
  StyledToggleIconButton,
  ContentContainer,
  MenuGroupTitle,
  StyledList,
  StyledListItemButton,
  StyledListItemIcon,
  StyledListItemText,
  BadgeContainer,
  StyledDivider,
} from '@/components/layout/sidebar/styled';

interface MenuItem {
  text: string;
  icon: React.ReactNode;
  path: string;
  badge?: number;
}

interface MenuGroup {
  title?: string;
  items: MenuItem[];
}

const menuGroups: MenuGroup[] = [
  {
    items: [
      {
        text: 'Dashboard',
        icon: <DashboardIcon />,
        path: '/dashboard',
      },
    ],
  },
  {
    title: 'Shop',
    items: [
      {
        text: 'Products',
        icon: <Inventory2 />,
        path: '/products',
      },
      {
        text: 'Orders',
        icon: <ShoppingBag />,
        path: '/orders',
        badge: 0,
      },
      {
        text: 'Customers',
        icon: <People />,
        path: '/customers',
      },
    ],
  },
  {
    title: 'Blog',
    items: [
      {
        text: 'Posts',
        icon: <Article />,
        path: '/posts',
      },
      {
        text: 'Categories',
        icon: <Folder />,
        path: '/categories',
      },
      {
        text: 'Authors',
        icon: <Person />,
        path: '/authors',
      },
    ],
  },
];

export const Sidebar = () => {
  const pathname = usePathname();
  const router = useRouter();
  const { navbarOpen, setNavbarOpen } = useContext(AppContext);

  const handleNavbarToggle = useCallback(() => {
    setNavbarOpen(!navbarOpen);
  }, [setNavbarOpen, navbarOpen]);

  const handleNavigation = (path: string) => {
    router.push(path);
  };

  return (
    <StyledDrawer variant="permanent" open={navbarOpen}>
      <SidebarContainer>
        <ContentContainer>
          {menuGroups.map((group, groupIndex) => (
            <React.Fragment key={groupIndex}>
              {group.title && navbarOpen && <MenuGroupTitle>{group.title}</MenuGroupTitle>}
              <StyledList>
                {group.items.map((item) => {
                  const isActive = pathname === item.path;
                  return (
                    <ListItem key={item.text} disablePadding>
                      <StyledListItemButton
                        open={navbarOpen}
                        selected={isActive}
                        onClick={() => handleNavigation(item.path)}
                      >
                        <StyledListItemIcon open={navbarOpen}>{item.icon}</StyledListItemIcon>
                        {navbarOpen && (
                          <>
                            <StyledListItemText open={navbarOpen} primary={item.text} />
                            {item.badge !== undefined && (
                              <BadgeContainer>{item.badge}</BadgeContainer>
                            )}
                          </>
                        )}
                      </StyledListItemButton>
                    </ListItem>
                  );
                })}
              </StyledList>
              {groupIndex < menuGroups.length - 1 && navbarOpen && <StyledDivider />}
            </React.Fragment>
          ))}
        </ContentContainer>

        <StyledBottomSection open={navbarOpen}>
          <Divider />
          <StyledToggleIconButton onClick={handleNavbarToggle} size="small">
            {navbarOpen ? <ChevronLeft /> : <MenuIcon />}
          </StyledToggleIconButton>
        </StyledBottomSection>
      </SidebarContainer>
    </StyledDrawer>
  );
};
