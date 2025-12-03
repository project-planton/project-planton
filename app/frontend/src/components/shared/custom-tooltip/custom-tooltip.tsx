'use client';
import React from 'react';
import { Tooltip, TooltipProps, tooltipClasses } from '@mui/material';
import { styled } from '@mui/material/styles';

export const CustomTooltip = styled(({ className, ...props }: TooltipProps) => (
  <Tooltip {...props} arrow classes={{ popper: className }} />
))(() => ({
  [`& .${tooltipClasses.arrow}`]: {
    color: '#0E131E',
  },
  [`& .${tooltipClasses.tooltip}`]: {
    maxWidth: '500px',
    borderRadius: '5px',
    border: '1px solid #0E131E',
    background: '#0E131E',
    boxShadow: '0px 4px 12px 0px rgba(0, 0, 0, 0.12)',
    padding: '5px 10px',
    fontSize: '12px',
    whiteSpace: 'pre-line',
  },
}));

