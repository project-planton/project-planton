'use client';
import { FC } from 'react';
import { InputLabel, InputLabelProps } from '@mui/material';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';
import { HelpTooltip } from '@/components/shared/help-tooltip';

export interface InputLabelHelpProps extends InputLabelProps {
  label: string;
  help?: string;
}

export const InputLabelHelp: FC<InputLabelHelpProps> = ({ label, help, ...props }) => {
  return (
    <InputLabel {...props}>
      <FlexCenterRow gap={0.5}>
        {label}
        {label && help && <HelpTooltip title={help} />}
      </FlexCenterRow>
    </InputLabel>
  );
};

