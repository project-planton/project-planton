'use client';

import { Grid2, InputLabelProps, MenuItem, TextField, TextFieldProps } from '@mui/material';
import { ReactNode } from 'react';
import { FormField } from '@/components/shared/form-field';
import { InputLabelHelp } from '@/components/shared/input-label-help';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';

export interface SimpleSelectOption {
  label: string;
  value: string | number;
}

export interface SimpleSelectProps extends Omit<TextFieldProps, 'select' | 'onChange'> {
  name?: string;
  value: string | number;
  onChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
  options: SimpleSelectOption[];
  required?: boolean;
  disabled?: boolean;
  fullWidth?: boolean;
  placeholder?: string;
  help?: string;
  inputLabelProps?: InputLabelProps;
  endLabel?: ReactNode;
}

export const SimpleSelect = ({
  name,
  value,
  onChange,
  options,
  required = false,
  disabled = false,
  fullWidth = false,
  placeholder,
  help,
  inputLabelProps,
  endLabel,
  size = 'medium',
  ...textFieldProps
}: SimpleSelectProps) => {
  const textField = (
    <TextField
      select
      value={value}
      onChange={onChange}
      required={required}
      disabled={disabled}
      fullWidth={fullWidth}
      placeholder={placeholder}
      size={size}
      {...textFieldProps}
    >
      {options.map((opt) => (
        <MenuItem key={opt.value} value={opt.value}>
          {opt.label}
        </MenuItem>
      ))}
    </TextField>
  );
  return (
    <Grid2 container spacing={1.5} flexGrow={fullWidth ? 1 : 'inherit'}>
      <Grid2 size={12}>
        {name || endLabel ? (
          <>
            <FlexCenterRow justifyContent={'space-between'} gap={1}>
              {name && <InputLabelHelp label={name} help={help} {...inputLabelProps} />}
              {endLabel}
            </FlexCenterRow>
            <FormField fullWidth={fullWidth}>{textField}</FormField>
          </>
        ) : (
          textField
        )}
      </Grid2>
    </Grid2>
  );
};
