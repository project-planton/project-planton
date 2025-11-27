'use client';
import { Inter } from 'next/font/google';
import { AppContextProvider, PCThemeType } from '@/contexts';

interface ProvidersProps {
  children: React.ReactNode;
  connectHost: string;
  cookieThemeMode?: PCThemeType;
  cookieNavbarOpen?: boolean;
}

const inter = Inter({
  weight: ['400', '500', '600', '700'],
  subsets: ['latin'],
  display: 'swap',
});

export function Providers({
  children,
  connectHost,
  cookieThemeMode,
  cookieNavbarOpen,
}: ProvidersProps) {
  return (
    <AppContextProvider
      connectHost={connectHost}
      font={inter}
      cookieThemeMode={cookieThemeMode}
      cookieNavbarOpen={cookieNavbarOpen}
    >
      {children}
    </AppContextProvider>
  );
}
