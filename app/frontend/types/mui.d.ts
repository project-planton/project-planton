import '@mui/material/Alert';

import { PaletteOptions } from '@mui/material/styles';

declare module '@mui/material/styles' {
  interface Palette {
    neutral: Palette['primary'];
  }

  interface PaletteOptions {
    neutral?: PaletteOptions['primary'];
  }
}

declare module '@mui/material/Alert' {
  interface AlertPropsColorOverrides {
    neutral: true;
  }
}
