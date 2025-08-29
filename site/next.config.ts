import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: true,
  // Export as a fully static site (for GitHub Pages)
  output: "export",
  images: {
    // Static export requires unoptimized images
    unoptimized: true,
  },
};

export default nextConfig;
