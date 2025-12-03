'use client';
import React, { useContext } from 'react';
import { AppContext } from '@/contexts';
// import { Badge, MenuItem, Divider, Typography } from '@mui/material';
import {
  StyledAppBar,
  StyledToolbar,
  Spacer,
  RightSection,
  StyledLinearProgress,
  // SearchBox,
  // StyledSearchIcon,
  // StyledInputBase,
  // StyledAvatarButton,
  // StyledAvatar,
  // StyledMenu,
  // StyledMenuItemIcon,
} from '@/components/layout/header/styled';
import { HeaderIcon } from '@/components/layout/header/header-icon';
import { ThemeSwitch } from '@/components/layout/theme-switch';

export const Header = () => {
  const { pageLoading } = useContext(AppContext);
  // const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  // const open = Boolean(anchorEl);

  // const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
  //   setAnchorEl(event.currentTarget);
  // };

  // const handleMenuClose = () => {
  //   setAnchorEl(null);
  // };

  return (
    <StyledAppBar>
      {pageLoading && <StyledLinearProgress color="primary" />}
      <StyledToolbar>
        <HeaderIcon />

        <Spacer />

        <RightSection>
          {/* <SearchBox>
            <StyledSearchIcon />
            <StyledInputBase placeholder="Search..." />
          </SearchBox> */}

          <ThemeSwitch />

          {/* <StyledIconButton size="small">
            <Badge badgeContent={0} color="error">
              <NotificationsOutlined />
            </Badge>
          </StyledIconButton>

          <StyledAvatarButton onClick={handleMenuOpen} size="small">
            <StyledAvatar>DU</StyledAvatar>
          </StyledAvatarButton> */}
        </RightSection>

        {/* <StyledMenu
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
        </StyledMenu> */}
      </StyledToolbar>
    </StyledAppBar>
  );
};
