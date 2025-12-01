'use client';
import React, { useContext } from 'react';
import { AppContext } from '@/contexts';
import { Badge, MenuItem, Divider, Typography } from '@mui/material';
import {
  Brightness4,
  Brightness7,
  NotificationsOutlined,
  AccountCircle,
  Logout,
  Settings,
} from '@mui/icons-material';
import {
  StyledAppBar,
  StyledToolbar,
  LogoSection,
  StyledLogoText,
  Spacer,
  RightSection,
  SearchBox,
  StyledSearchIcon,
  StyledInputBase,
  StyledIconButton,
  StyledAvatarButton,
  StyledAvatar,
  StyledMenu,
  StyledMenuItemIcon,
  StyledLinearProgress,
} from '@/components/layout/header/styled';

export const Header = () => {
  const { theme, changeTheme, pageLoading } = useContext(AppContext);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleThemeToggle = () => {
    const newTheme = theme.mode === 'light' ? 'dark' : 'light';
    changeTheme(newTheme);
  };

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  return (
    <StyledAppBar>
      {pageLoading && <StyledLinearProgress color="primary" />}
      <StyledToolbar>
        <LogoSection>
          <StyledLogoText variant="h6">Project Planton</StyledLogoText>
        </LogoSection>

        <Spacer />

        <RightSection>
          <SearchBox>
            <StyledSearchIcon />
            <StyledInputBase placeholder="Search..." />
          </SearchBox>

          <StyledIconButton onClick={handleThemeToggle} size="small">
            {theme.mode === 'light' ? <Brightness4 /> : <Brightness7 />}
          </StyledIconButton>

          <StyledIconButton size="small">
            <Badge badgeContent={0} color="error">
              <NotificationsOutlined />
            </Badge>
          </StyledIconButton>

          <StyledAvatarButton onClick={handleMenuOpen} size="small">
            <StyledAvatar>DU</StyledAvatar>
          </StyledAvatarButton>
        </RightSection>

        <StyledMenu
          anchorEl={anchorEl}
          open={open}
          onClose={handleMenuClose}
          onClick={handleMenuClose}
          transformOrigin={{ horizontal: 'right', vertical: 'top' }}
          anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
        >
          <MenuItem onClick={handleMenuClose}>
            <StyledMenuItemIcon>
              <AccountCircle />
            </StyledMenuItemIcon>
            <Typography variant="body2">Profile</Typography>
          </MenuItem>
          <MenuItem onClick={handleMenuClose}>
            <StyledMenuItemIcon>
              <Settings />
            </StyledMenuItemIcon>
            <Typography variant="body2">Settings</Typography>
          </MenuItem>
          <Divider />
          <MenuItem onClick={handleMenuClose}>
            <StyledMenuItemIcon>
              <Logout />
            </StyledMenuItemIcon>
            <Typography variant="body2">Sign out</Typography>
          </MenuItem>
        </StyledMenu>
      </StyledToolbar>
    </StyledAppBar>
  );
};
