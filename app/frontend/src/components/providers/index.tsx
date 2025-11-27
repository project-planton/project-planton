'use client';
import { Inter } from 'next/font/google';
import { AppContextProvider } from '@/contexts';

interface ProvidersProps {
  children: React.ReactNode;
  connectHost: string;
}

const inter = Inter({
  weight: ['400', '500', '600', '700'],
  subsets: ['latin'],
  display: 'swap',
});

export function Providers({ children, connectHost }: ProvidersProps) {
  return (
    <AppContextProvider connectHost={connectHost} font={inter}>
      {children}
    </AppContextProvider>
  );
}

