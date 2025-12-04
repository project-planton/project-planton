'use client';

import { Chip, ChipProps } from '@mui/material';

export interface StatusChipProps extends Omit<ChipProps, 'label' | 'color'> {
  status: string;
  colorMap?: Record<string, 'success' | 'error' | 'warning' | 'info' | 'default'>;
}

const defaultColorMap: Record<string, 'success' | 'error' | 'warning' | 'info' | 'default'> = {
  success: 'success',
  failed: 'error',
  in_progress: 'warning',
  pending: 'info',
};

export function StatusChip({ status, colorMap = defaultColorMap, ...chipProps }: StatusChipProps) {
  const color = colorMap[status] || 'default';

  return (
    <Chip
      label={status}
      color={color}
      size="small"
      sx={{ textTransform: 'capitalize', ...chipProps.sx }}
      {...chipProps}
    />
  );
}

