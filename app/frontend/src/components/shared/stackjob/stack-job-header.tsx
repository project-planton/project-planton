'use client';

import { Typography, Divider, Stack, Box, Skeleton } from '@mui/material';
import { HeaderContainer, TopSection, BottomSection, FlexCenterRow } from '@/app/stack-jobs/styled';
import { StackJob } from '@/gen/proto/stack_job_service_pb';
import { StatusChip } from '@/components/shared/status-chip';
import { TextCopy } from '@/components/shared/text-copy';

interface StackJobHeaderProps {
  stackJob: StackJob | null;
  updatedTime: string;
}

export function StackJobHeader({ stackJob, updatedTime }: StackJobHeaderProps) {
  return (
    <HeaderContainer>
      <TopSection>
        <Stack gap={1}>
          <FlexCenterRow gap={0.5}>
            {stackJob ? (
              <Typography variant="caption" color="text.secondary">
                {stackJob?.id}
              </Typography>
            ) : (
              <Skeleton variant="text" width={180} height={15} />
            )}

            {stackJob ? (
              <TextCopy text={stackJob?.id} />
            ) : (
              <Skeleton variant="rectangular" width={10} height={10} />
            )}
          </FlexCenterRow>
          <Box>
            {stackJob ? (
              <StatusChip status={stackJob?.status || 'unknown'} />
            ) : (
              <Skeleton variant="rounded" width={80} height={20} />
            )}
          </Box>
        </Stack>
      </TopSection>
      <Divider sx={{ marginY: 1.5 }} />
      <BottomSection>
        <Box />
        {stackJob ? (
          <Typography variant="subtitle2" color="text.secondary">
            {updatedTime}
          </Typography>
        ) : (
          <Skeleton variant="text" width={180} height={15} />
        )}
      </BottomSection>
    </HeaderContainer>
  );
}
