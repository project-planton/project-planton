'use client';
import { FC, ReactNode } from 'react';
import { CustomTooltip } from '@/components/shared/custom-tooltip';
import { Help } from '@mui/icons-material';

interface HelpTooltipProps {
  title: ReactNode;
}

export const HelpTooltip: FC<HelpTooltipProps> = ({ title }) => {
  if (!title) return <></>;
  return (
    <CustomTooltip title={title}>
      <Help fontSize="small" />
    </CustomTooltip>
  );
};

