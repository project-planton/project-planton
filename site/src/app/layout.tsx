import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  metadataBase: new URL("https://project-planton.org"),
  title: {
    default: "ProjectPlanton — Open‑Source Multi‑Cloud Infrastructure Framework",
    template: "%s | ProjectPlanton",
  },
  description:
    "Kubernetes‑style manifests for multi‑cloud infrastructure. Define once. Deploy anywhere.",
  keywords: [
    "ProjectPlanton",
    "multi‑cloud",
    "infrastructure",
    "Pulumi",
    "OpenTofu",
    "Kubernetes",
    "Protobuf",
    "Buf",
  ],
  openGraph: {
    type: "website",
    url: "/",
    title: "ProjectPlanton — Open‑Source Multi‑Cloud Infrastructure Framework",
    description:
      "Kubernetes‑style manifests for multi‑cloud infrastructure. Define once. Deploy anywhere.",
    siteName: "ProjectPlanton",
    images: [
      { url: "/icon.png", width: 512, height: 512, alt: "ProjectPlanton" },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "ProjectPlanton — Open‑Source Multi‑Cloud Infrastructure Framework",
    description:
      "Kubernetes‑style manifests for multi‑cloud infrastructure. Define once. Deploy anywhere.",
    images: ["/icon.png"],
  },
  icons: {
    icon: "/favicon.ico",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased min-h-screen bg-slate-950 text-white`}
      >
        {children}
      </body>
    </html>
  );
}
