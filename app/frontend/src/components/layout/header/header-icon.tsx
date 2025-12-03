'use client';
import { Stack } from '@mui/material';
import { useRouter, useSearchParams } from 'next/navigation';
import { FC } from 'react';
import { Icon, ICON_NAMES } from '@/components/shared/icon';

export const HeaderIcon: FC = () => {
  const router = useRouter();
  const searchParams = useSearchParams();

  const clickHandler = () => {
    router.push(`/dashboard?${searchParams.toString()}`);
  };

  return (
    <Stack flexDirection={'row'} gap={0.5} alignItems={'flex-start'} id="header-icon">
      <Icon
        name={ICON_NAMES.PLANTON_LOGO}
        sx={{ width: 32, height: 32 }}
        onClick={clickHandler}
      />
    </Stack>
  );
};

