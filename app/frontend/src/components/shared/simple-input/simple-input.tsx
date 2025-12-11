'use client';
import { Grid2, InputLabelProps, TextField, TextFieldProps } from '@mui/material';
import { HTMLInputTypeAttribute, ReactNode } from 'react';
import { FormField } from '@/components/shared/form-field';
import { InputLabelHelp } from '@/components/shared/input-label-help';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';
import { Path, RegisterOptions, UseFormRegister } from 'react-hook-form';

interface ISimpleInput<T> {
  register: UseFormRegister<T>;
  registerOptions?: RegisterOptions<T>;
  path: Path<T>;
  name?: string;
  type?: HTMLInputTypeAttribute;
  multiline?: boolean;
  fullWidth?: boolean;
  disabled?: boolean;
  rows?: string | number;
  textFieldProps?: TextFieldProps;
  inputLabelProps?: InputLabelProps;
  endLabel?: ReactNode;
}

export const SimpleInput = <T,>({
  register,
  registerOptions,
  path,
  name,
  type,
  multiline = false,
  rows = 1,
  disabled,
  textFieldProps,
  fullWidth = false,
  inputLabelProps,
  endLabel,
}: ISimpleInput<T>) => {
  return (
    <Grid2 container flexGrow={fullWidth ? 1 : 'inherit'}>
      <Grid2 size={12}>
        <FlexCenterRow justifyContent={'space-between'} gap={1}>
          {name && <InputLabelHelp label={name} {...inputLabelProps} />}
          {endLabel}
        </FlexCenterRow>
        <FormField fullWidth={fullWidth}>
          <TextField
            type={type}
            {...(register ? register(path, registerOptions) : {})}
            multiline={multiline}
            rows={rows}
            disabled={disabled}
            {...textFieldProps}
          />
        </FormField>
      </Grid2>
    </Grid2>
  );
};
