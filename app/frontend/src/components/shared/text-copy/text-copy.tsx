'use client';
import React, { FC, useContext } from 'react';
import { ContentCopyRounded } from '@mui/icons-material';
import { SvgIconProps, styled } from '@mui/material';
import { AppContext } from '@/contexts';
import { copyText } from '@/lib';

const StyledContentCopy = styled(ContentCopyRounded)`
  width: 16px;
  height: 16px;
  margin-left: 5px;
  cursor: pointer;
`;

type TextCopyProps = {
  text?: string;
} & SvgIconProps;

export const useCopy = () => {
  const { openSnackbar } = useContext(AppContext);

  const handleCopyClick = (text: string) => {
    copyText(text).then(() => {
      if (openSnackbar) {
        openSnackbar('Copied!', 'success');
      }
    });
  };

  return handleCopyClick;
};

export const TextCopy: FC<TextCopyProps> = ({ text = null, ...props }) => {
  const handleCopyClick = useCopy();
  if (text === null) return null;
  return <StyledContentCopy onClick={() => handleCopyClick(text)} {...props} />;
};

