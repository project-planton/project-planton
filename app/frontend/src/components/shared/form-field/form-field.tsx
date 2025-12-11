'use client';
import { FC, ReactNode, CSSProperties } from 'react';
import { FormControl, FormControlProps, FormHelperText } from '@mui/material';
import { styled } from '@mui/material';

const StyledFormControl = styled(FormControl)(({ theme }) => ({
  width: '100%',
  maxWidth: 442,
  borderColor: theme.palette.grey[80],
  borderRadius: 4,
  margin: theme.spacing(1.5, 0),
  [theme.breakpoints.up('md')]: { margin: theme.spacing(2, 0) },
}));

const FormErrorText = styled(FormHelperText)`
  font-size: 12px;
`;

type FormFieldProps = {
  errorMessage?: string;
  height?: string;
  children: ReactNode;
};

export const FormField: FC<FormFieldProps & FormControlProps> = ({
  errorMessage = '',
  height,
  children,
  fullWidth,
  style,
  ...formControlProps
}) => {
  let inlineStyles: CSSProperties = { height };
  if (fullWidth) {
    inlineStyles = { ...inlineStyles, maxWidth: 'none', width: '100%' };
  }
  inlineStyles = { ...inlineStyles, ...style };

  return (
    <StyledFormControl {...formControlProps} style={inlineStyles}>
      {children}
      {errorMessage !== '' && <FormErrorText error>{errorMessage}</FormErrorText>}
    </StyledFormControl>
  );
};

