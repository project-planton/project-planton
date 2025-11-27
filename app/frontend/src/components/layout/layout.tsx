'use client';
import React from 'react';
import { StyledEngineProvider, Box } from '@mui/material';
import { Header } from './header';
import { Sidebar } from './sidebar';
import { StyledContainer, StyledWrapperBox } from '@/components/layout/styled';

interface LayoutProps {
  children: React.ReactNode;
}

export const Layout = ({ children }: LayoutProps) => {
  return (
    <StyledEngineProvider injectFirst>
      <Header />
      <StyledWrapperBox>
        <Sidebar />
        <StyledContainer>
          {children}
        </StyledContainer>
      </StyledWrapperBox>
    </StyledEngineProvider>
  );
};

