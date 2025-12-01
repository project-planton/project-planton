import '@mui/material/styles';

// Extend MUI's Palette interface to include 'exceptions'
declare module '@mui/material/styles' {
  interface Palette {
    exceptions: PaletteColor;
    crimson: PaletteColor;
  }

  interface PaletteOptions {
    exceptions?: PaletteColorOptions;
    crimson?: PaletteColorOptions;
  }

  interface TypeText {
    link: string;
  }
}
