'use client';
import React, { useState } from 'react';
import { ListItem, Divider } from '@mui/material';
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
} from './styled';

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
  const [open, setOpen] = useState(true);

  const handleToggle = () => {
    setOpen(!open);
  };

  const handleNavigation = (path: string) => {
    router.push(path);
  };

  return (
    <StyledDrawer variant="permanent" open={open}>
      <SidebarContainer>
        <ContentContainer>
          {menuGroups.map((group, groupIndex) => (
            <React.Fragment key={groupIndex}>
              {group.title && open && <MenuGroupTitle>{group.title}</MenuGroupTitle>}
              <StyledList>
                {group.items.map((item) => {
                  const isActive = pathname === item.path;
                  return (
                    <ListItem key={item.text} disablePadding>
                      <StyledListItemButton
                        open={open}
                        selected={isActive}
                        onClick={() => handleNavigation(item.path)}
                      >
                        <StyledListItemIcon open={open}>{item.icon}</StyledListItemIcon>
                        {open && (
                          <>
                            <StyledListItemText open={open} primary={item.text} />
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
              {groupIndex < menuGroups.length - 1 && open && <StyledDivider />}
            </React.Fragment>
          ))}
        </ContentContainer>

        <StyledBottomSection open={open}>
          <Divider />
          <StyledToggleIconButton onClick={handleToggle} size="small">
            {open ? <ChevronLeft /> : <MenuIcon />}
          </StyledToggleIconButton>
        </StyledBottomSection>
      </SidebarContainer>
    </StyledDrawer>
  );
};

