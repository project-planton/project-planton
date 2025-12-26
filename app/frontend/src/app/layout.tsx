import { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { Layout } from '@/components';
import { Providers } from '@/components/providers';
import { getAllCookiesParsed, getCookieThemeMode, getCookieNavbarOpen } from '@/lib/server/cookies';

export const metadata: Metadata = {
  title: 'Project Planton Web App',
  description: 'Project Planton Web Application',
};

const inter = Inter({
  weight: ['400', '500', '600', '700'],
  subsets: ['latin'],
  display: 'swap',
});

export default function RootLayout({ children }: { children: React.ReactNode }) {
  // Default to localhost:50051 if NEXT_PUBLIC_API_URL is not set
  const connectHost = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:50051';

  // Get cookies for SSR
  const allCookies = getAllCookiesParsed();
  const cookieThemeMode = getCookieThemeMode(allCookies);
  const cookieNavbarOpen = getCookieNavbarOpen(allCookies);

  return (
    <html lang="en" className={inter.className}>
      <body>
        <Providers
          connectHost={connectHost}
          cookieThemeMode={cookieThemeMode}
          cookieNavbarOpen={cookieNavbarOpen}
        >
          <Layout>{children}</Layout>
        </Providers>
      </body>
    </html>
  );
}
