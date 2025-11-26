import { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { Layout } from '@/components';
import { Providers } from '@/components/providers';

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
  // Default to localhost:50051 if API_ENDPOINT is not set
  const connectHost = process.env.API_ENDPOINT || 'http://localhost:50051';

  return (
    <html lang="en" className={inter.className}>
      <body>
        <Providers connectHost={connectHost}>
          <Layout>{children}</Layout>
        </Providers>
      </body>
    </html>
  );
}

