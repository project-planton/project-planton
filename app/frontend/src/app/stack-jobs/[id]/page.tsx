'use client';

import { useEffect, useState, useMemo, useCallback, useContext } from 'react';
import { useParams } from 'next/navigation';
import { Box, Skeleton, Stack } from '@mui/material';
import { StackJobContainer } from '@/app/stack-jobs/styled';
import { useStackJobQuery } from '@/app/stack-jobs/_services';
import { StackJob } from '@/gen/proto/stack_job_service_pb';
import { Breadcrumb, BreadcrumbStartIcon, IBreadcrumbItem } from '@/components/shared/breadcrumb';
import { StackJobsDrawer, StackJobHeader } from '@/components/shared/stackjob';
import { ICON_NAMES } from '@/components/shared/icon';
import { formatTimestampToDate } from '@/lib';
import { JsonCode } from '@/components/shared/syntax-highlighter';
import { AppContext, THEME } from '@/contexts';

export default function StackJobDetailPage() {
  const { theme } = useContext(AppContext);
  const params = useParams();
  const { query } = useStackJobQuery();
  const [stackJob, setStackJob] = useState<StackJob | null>(null);
  const [stackJobsDrawerOpen, setStackJobsDrawerOpen] = useState(false);

  const stackJobId = params?.id as string;

  const handleCloseStackJobs = useCallback(() => {
    setStackJobsDrawerOpen(false);
  }, []);

  const handleStackJobsClick = useCallback(() => {
    if (stackJob?.cloudResourceId) {
      setStackJobsDrawerOpen(true);
    }
  }, [stackJob?.cloudResourceId]);

  const breadcrumbs: IBreadcrumbItem[] = useMemo(() => {
    const items: IBreadcrumbItem[] = [];

    // Always show the ID from params, even if stackJob is not loaded yet
    if (stackJobId) {
      items.push({
        name: stackJobId,
        handler: undefined, // Last item is not clickable
      });
    }

    return items;
  }, [stackJobId, handleStackJobsClick]);

  useEffect(() => {
    if (query && stackJobId) {
      query.getById(stackJobId).then((job) => {
        setStackJob(job);
      });
    }
  }, [query, stackJobId]);

  const updatedTime = stackJob?.updatedAt
    ? formatTimestampToDate(stackJob.updatedAt, 'DD/MM/YYYY, HH:mm:ss')
    : '-';

  return (
    <StackJobContainer>
      <Stack gap={2}>
        <Breadcrumb
          breadcrumbs={breadcrumbs}
          startBreadcrumb={
            <BreadcrumbStartIcon
              icon={ICON_NAMES.INFRA_HUB}
              iconProps={{ sx: { filter: theme.mode === THEME.DARK ? 'invert(1)' : 'none' } }}
              label="Stack Jobs"
              handler={handleStackJobsClick}
            />
          }
        />
        <StackJobHeader stackJob={stackJob} updatedTime={updatedTime} />

        <Box>
          {stackJob ? (
            <JsonCode content={stackJob?.output || {}} />
          ) : (
            <Skeleton variant="rounded" width={'100%'} height={200} />
          )}
        </Box>

        {/* Stack Jobs Drawer */}
        {stackJob?.cloudResourceId && (
          <StackJobsDrawer
            open={stackJobsDrawerOpen}
            cloudResourceId={stackJob.cloudResourceId}
            onClose={handleCloseStackJobs}
          />
        )}
      </Stack>
    </StackJobContainer>
  );
}
