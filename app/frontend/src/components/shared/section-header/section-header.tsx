'use client';
import React, { FC, ReactNode } from 'react';
import { Stack, StackProps, Typography } from '@mui/material';
import { HeaderContainer, LeftSection } from '@/components/shared/section-header/styled';

type SectionHeaderProps = {
  title?: string;
  children?: ReactNode;
  containerProps?: StackProps;
  borderBottom?: boolean;
};

export const SectionHeader: FC<SectionHeaderProps> = ({
  title,
  children,
  containerProps,
  borderBottom,
}) => {
  return (
    <HeaderContainer $bb={borderBottom} {...containerProps}>
      {!!title && (
        <LeftSection>
          <Typography fontSize={20} fontWeight={600}>
            {title}
          </Typography>
        </LeftSection>
      )}

      <Stack flexDirection="row" gap={{ xs: 1, md: 2 }}>
        {children}
      </Stack>
    </HeaderContainer>
  );
};
