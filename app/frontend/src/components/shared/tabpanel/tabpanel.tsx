'use client';
import React from 'react';
import { Box } from '@mui/material';

export interface TabPanelProps {
  children?: React.ReactNode;
  index: number | string;
  value: number | string;
}

export const TabPanel = (props: TabPanelProps) => {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`tab-panel-${index}`}
      style={{ height: '100%' }}
      {...other}
    >
      {value === index && <Box height={'100%'}>{children}</Box>}
    </div>
  );
};

