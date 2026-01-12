'use client';
import React, { useContext, useMemo } from 'react';
import { usePathname } from 'next/navigation';
import { StyledEngineProvider, CssBaseline } from '@mui/material';
import { Header } from '@/components/layout/header';
import { Sidebar } from '@/components/layout/sidebar';
import { StyledContainer, StyledWrapperBox } from '@/components/layout/styled';
import { AppContext } from '@/contexts';

interface LayoutProps {
  children: React.ReactNode;
}

export const Layout = ({ children }: LayoutProps) => {
  const pathname = usePathname();
  const { navbarOpen } = useContext(AppContext);

  const fullWidth = useMemo(
    () =>
      pathname.startsWith('/credentials') ||
      pathname.startsWith('/cloud-resources') ||
      pathname === '/' ||
      pathname.startsWith('/dashboard'),
    [pathname]
  );

  const hasWhiteBg = useMemo(
    () =>
      pathname.startsWith('/credentials') ||
      pathname.startsWith('/cloud-resources') ||
      pathname === '/' ||
      pathname.startsWith('/dashboard'),
    [pathname]
  );

  return (
    <StyledEngineProvider injectFirst>
      <CssBaseline />
      <Header />
      <StyledWrapperBox $hasWhiteBg={hasWhiteBg}>
        <Sidebar />
        <StyledContainer $fullWidth={fullWidth} $navbarOpen={navbarOpen}>
          {children}
        </StyledContainer>
      </StyledWrapperBox>
    </StyledEngineProvider>
  );
};
