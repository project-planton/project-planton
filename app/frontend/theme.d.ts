import '@mui/material/styles';

// Extend MUI's Palette interface to include custom colors
declare module '@mui/material/styles' {
  interface Palette {
    exceptions: PaletteColor;
    crimson: PaletteColor;
    neutral: {
      main: string;
      contrastText: string;
    };
  }

  interface PaletteOptions {
    exceptions?: PaletteColorOptions;
    crimson?: PaletteColorOptions;
    neutral?: {
      main: string;
      contrastText: string;
    };
  }

  interface PaletteColor {
    lighter?: string;
    darker?: string;
  }

  interface SimplePaletteColorOptions {
    lighter?: string;
    darker?: string;
  }

  interface TypeText {
    link: string;
  }

  interface TypeBackground {
    default: string;
    paper: string;
  }
}

// Extend MUI's Alert color options
declare module '@mui/material/Alert' {
  interface AlertPropsColorOverrides {
    neutral: true;
  }
}

// Extend MUI's Button color options  
declare module '@mui/material/Button' {
  interface ButtonPropsColorOverrides {
    neutral: true;
  }
}
