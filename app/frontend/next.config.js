/* eslint-disable @typescript-eslint/no-var-requires */
const { resolve } = require('path');

// this is needed for local Development;
// this adds all the variables from '.env' to node process.env
require('dotenv').config();

module.exports = {
  webpack: (config) => {
    config.cache = false;
    
    // Handle SVG imports as React components
    config.module.rules.push({
      test: /\.svg$/,
      use: ['@svgr/webpack'],
    });
    
    return config;
  },
  output: 'standalone',
  compiler: {
    emotion: {
      sourceMap: true,
      autoLabel: 'dev-only',
      labelFormat: '[local]',
      importMap: {
        '@mui/system': {
          styled: {
            canonicalImport: ['@emotion/styled', 'default'],
            styledBaseImport: ['@mui/system', 'styled'],
          },
        },
        '@mui/material/styles': {
          styled: {
            canonicalImport: ['@emotion/styled', 'default'],
            styledBaseImport: ['@mui/material/styles', 'styled'],
          },
        },
      },
    },
  },
};

