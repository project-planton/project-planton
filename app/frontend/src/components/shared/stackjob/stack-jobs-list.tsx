'use client';

import { useState, useMemo, useCallback, useEffect } from 'react';
import { Typography, Box } from '@mui/material';
import { useRouter } from 'next/navigation';
import { create } from '@bufbuild/protobuf';
import { TableComp } from '@/components/shared/table';
import { PAGINATION_MODE } from '@/models/table';
import { StatusChip } from '@/components/shared/status-chip';
import { useStackJobQuery } from '@/app/stack-jobs/_services';
import {
  ListStackJobsRequestSchema,
  StackJob,
} from '@/gen/proto/stack_job_service_pb';
import { PageInfoSchema } from '@/gen/proto/cloud_resource_service_pb';
import { formatTimestampToDate } from '@/lib';

export interface StackJobsListProps {
  cloudResourceId: string;
}

export function StackJobsList({ cloudResourceId }: StackJobsListProps) {
  const router = useRouter();
  const { query } = useStackJobQuery();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [stackJobs, setStackJobs] = useState<StackJob[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [apiLoading, setApiLoading] = useState(true);

  // Function to call the API
  const handleLoadStackJobs = useCallback(() => {
    setApiLoading(true);
    if (query) {
      query
        .listStackJobs(
          create(ListStackJobsRequestSchema, {
            cloudResourceId,
            pageInfo: create(PageInfoSchema, {
              num: page,
              size: rowsPerPage,
            }),
          })
        )
        .then((result) => {
          setStackJobs(result.jobs);
          setTotalPages(result.totalPages || 0);
          setApiLoading(false);
        })
        .catch(() => {
          setApiLoading(false);
        });
    }
  }, [query, cloudResourceId, page, rowsPerPage]);

  // Reset to first page when cloudResourceId changes
  useEffect(() => {
    setPage(0);
  }, [cloudResourceId]);

  // Auto-load on mount and when dependencies change
  useEffect(() => {
    if (query && cloudResourceId) {
      handleLoadStackJobs();
    }
  }, [query, cloudResourceId, handleLoadStackJobs]);

  // Handle page change
  const handlePageChange = useCallback((newPage: number, newRowsPerPage: number) => {
    setPage(newPage);
    setRowsPerPage(newRowsPerPage);
  }, []);

  const tableDataTemplate = useMemo(
    () => ({
      id: (val: string, row: StackJob) => {
        return <Typography variant="subtitle2">{row.id.substring(0, 12)}...</Typography>;
      },
      status: (val: string, row: StackJob) => {
        return <StatusChip status={row.status || 'unknown'} />;
      },
      createdAt: (val: string, row: StackJob) => {
        return row.createdAt ? formatTimestampToDate(row.createdAt, 'DD/MM/YYYY HH:mm') : '-';
      },
      updatedAt: (val: string, row: StackJob) => {
        return row.updatedAt ? formatTimestampToDate(row.updatedAt, 'DD/MM/YYYY HH:mm') : '-';
      },
    }),
    []
  );

  const clickableColumns = useMemo(
    () => ({
      id: (row: StackJob) => {
        router.push(`/stack-jobs/${row.id}`);
      },
    }),
    [router]
  );

  return (
    <Box>
      <TableComp
        data={stackJobs}
        loading={apiLoading}
        options={{
          headers: ['ID', 'Status', 'Created At', 'Updated At'],
          dataPath: ['id', 'status', 'createdAt', 'updatedAt'],
          dataTemplate: tableDataTemplate,
          clickableColumns,
          actions: [],
          showPagination: true,
          paginationMode: PAGINATION_MODE.SERVER,
          currentPage: page,
          rowsPerPage: rowsPerPage,
          totalPages: totalPages,
          onPageChange: handlePageChange,
          emptyMessage: 'No stack jobs found',
          border: true,
        }}
      />
    </Box>
  );
}

