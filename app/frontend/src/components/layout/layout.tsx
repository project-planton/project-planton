'use client';
import React, { useContext, useMemo } from 'react';
import { usePathname } from 'next/navigation';
import { StyledEngineProvider } from '@mui/material';
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

  const fullWidth = useMemo(() => pathname.startsWith('/credentials'), [pathname]);

  const hasWhiteBg = useMemo(() => pathname.startsWith('/credentials'), [pathname]);

  return (
    <StyledEngineProvider injectFirst>
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
