'use client';

import { useEffect, useState, useMemo, useCallback, useContext, useRef } from 'react';
import { useParams } from 'next/navigation';
import { Box, Skeleton, Stack, Typography, Paper, Divider, Tabs, Tab } from '@mui/material';
import { StackJobContainer, LogContainer, LogEntry } from '@/app/stack-jobs/styled';
import { useStackJobQuery } from '@/app/stack-jobs/_services';
import { StackJob } from '@/gen/proto/stack_job_service_pb';
import { Breadcrumb, BreadcrumbStartIcon, IBreadcrumbItem } from '@/components/shared/breadcrumb';
import { StackJobsDrawer, StackJobHeader } from '@/components/shared/stackjob';
import { ICON_NAMES } from '@/components/shared/icon';
import { formatTimestampToDate } from '@/lib';
import { JsonCode } from '@/components/shared/syntax-highlighter';
import { AppContext, THEME } from '@/contexts';

interface StreamingLog {
  sequenceNum: number;
  content: string;
  streamType: 'stdout' | 'stderr';
  timestamp: Date;
}

export default function StackJobDetailPage() {
  const { theme } = useContext(AppContext);
  const params = useParams();
  const { query } = useStackJobQuery();
  const [stackJob, setStackJob] = useState<StackJob | null>(null);
  const [stackJobsDrawerOpen, setStackJobsDrawerOpen] = useState(false);
  const [streamingLogs, setStreamingLogs] = useState<StreamingLog[]>([]);
  const [isStreaming, setIsStreaming] = useState(false);
  const [tabIndex, setTabIndex] = useState(0);
  const [intervalCount, setIntervalCount] = useState(0);
  const abortControllerRef = useRef<AbortController | null>(null);
  const logContainerRef = useRef<HTMLDivElement | null>(null);
  const hasStartedStreamingRef = useRef<boolean>(false);
  const streamingJobIdRef = useRef<string | null>(null);
  const isStreamingRef = useRef<boolean>(false);
  const hasLoadedLogsRef = useRef<boolean>(false);

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

  // Auto-scroll to bottom when new logs arrive
  useEffect(() => {
    if (logContainerRef.current) {
      logContainerRef.current.scrollTop = logContainerRef.current.scrollHeight;
    }
  }, [streamingLogs]);

  // Create interval - updates every 5 seconds
  useEffect(() => {
    const intervalId = setInterval(() => {
      setIntervalCount((prevCount) => prevCount + 1);
    }, 5000); // 5 seconds

    return () => clearInterval(intervalId);
  }, []);

  // Fetch stack job details when intervalCount changes
  useEffect(() => {
    if (!query || !stackJobId) {
      return;
    }

    query
      .getById(stackJobId)
      .then((job) => {
        setStackJob((prevJob) => {
          // Only update if status or output changed to avoid unnecessary re-renders
          if (!prevJob || prevJob.status !== job.status || prevJob.output !== job.output) {
            return job;
          }
          return prevJob;
        });
      })
      .catch((error) => {
        console.error('Failed to fetch stack job:', error);
      });
  }, [query, stackJobId, intervalCount]);

  // Auto-start streaming when job is in progress or completed
  useEffect(() => {
    if (!stackJob || !query || !stackJobId) {
      return;
    }

    // Reset streaming state if job ID changed
    if (streamingJobIdRef.current && streamingJobIdRef.current !== stackJobId) {
      if (abortControllerRef.current) {
        abortControllerRef.current.abort();
        abortControllerRef.current = null;
      }
      isStreamingRef.current = false;
      setIsStreaming(false);
      hasStartedStreamingRef.current = false;
      hasLoadedLogsRef.current = false;
      streamingJobIdRef.current = null;
    }

    // Don't start if already streaming this job
    if (streamingJobIdRef.current === stackJobId || hasStartedStreamingRef.current) {
      return;
    }

    // Start streaming for in-progress jobs or completed jobs (to load existing logs)
    const shouldStartStreaming =
      stackJob.status === 'in_progress' ||
      ((stackJob.status === 'success' || stackJob.status === 'failed') &&
        !hasLoadedLogsRef.current);

    if (!shouldStartStreaming) {
      return;
    }

    // Mark that we're loading logs for completed jobs
    if (stackJob.status === 'success' || stackJob.status === 'failed') {
      hasLoadedLogsRef.current = true;
    }

    // Mark as started to prevent re-triggering
    streamingJobIdRef.current = stackJobId;
    hasStartedStreamingRef.current = true;
    isStreamingRef.current = true;
    setIsStreaming(true);
    setStreamingLogs([]);

    // Create new AbortController
    const abortController = new AbortController();
    abortControllerRef.current = abortController;

    // Start streaming in async function
    const streamAsync = async () => {
      try {
        const stream = query.streamOutput(stackJobId, undefined, abortController.signal);

        for await (const response of stream) {
          if (abortController.signal.aborted) {
            break;
          }

          const log: StreamingLog = {
            sequenceNum: response.sequenceNum,
            content: response.content,
            streamType: response.streamType === 'stderr' ? 'stderr' : 'stdout',
            timestamp: response.timestamp
              ? new Date(Number(response.timestamp.seconds) * 1000)
              : new Date(),
          };

          setStreamingLogs((prev) => [...prev, log]);

          // If stream is completed, stop streaming
          if (response.status === 'completed' || response.status === 'cancelled') {
            isStreamingRef.current = false;
            setIsStreaming(false);
            hasStartedStreamingRef.current = false;
            if (streamingJobIdRef.current === stackJobId) {
              streamingJobIdRef.current = null;
            }
            break;
          }
        }
      } catch (error: any) {
        if (error.name !== 'AbortError') {
          console.error('Stream error:', error);
          isStreamingRef.current = false;
          setIsStreaming(false);
          hasStartedStreamingRef.current = false;
          if (streamingJobIdRef.current === stackJobId) {
            streamingJobIdRef.current = null;
          }
        }
      }
    };

    streamAsync();

    // Cleanup on unmount or when stackJobId changes
    return () => {
      if (streamingJobIdRef.current === stackJobId) {
        if (abortControllerRef.current) {
          abortControllerRef.current.abort();
          abortControllerRef.current = null;
        }
        isStreamingRef.current = false;
        setIsStreaming(false);
        hasStartedStreamingRef.current = false;
        streamingJobIdRef.current = null;
      }
    };
  }, [stackJob?.status, stackJobId, query]);

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

        <Paper sx={{ p: 2 }}>
          <Tabs value={tabIndex} onChange={(_, newValue) => setTabIndex(newValue)}>
            <Tab label="Streaming Output" />
            <Tab label="Final Output" />
          </Tabs>

          <Divider sx={{ my: 2 }} />

          {tabIndex === 0 ? (
            <Box>
              <Box
                sx={{
                  display: 'flex',
                  justifyContent: 'space-between',
                  alignItems: 'center',
                  mb: 2,
                }}
              >
                <Typography variant="h6">Deployment Logs</Typography>
                <Typography variant="body2" color="text.secondary">
                  {streamingLogs.length} log entries
                  {isStreaming && ' • Streaming...'}
                </Typography>
              </Box>

              <LogContainer ref={logContainerRef}>
                {streamingLogs.length === 0 ? (
                  <Typography
                    variant="body2"
                    color="text.secondary"
                    sx={{ textAlign: 'center', py: 4 }}
                  >
                    {isStreaming
                      ? 'Waiting for logs...'
                      : stackJob?.status === 'in_progress'
                        ? 'Click to start streaming'
                        : 'No logs available'}
                  </Typography>
                ) : (
                  streamingLogs.map((log, index) => (
                    <LogEntry key={index} streamType={log.streamType}>
                      <Box sx={{ display: 'flex', gap: 1, alignItems: 'flex-start' }}>
                        <Typography
                          variant="caption"
                          color="text.secondary"
                          sx={{ minWidth: 180, fontFamily: 'monospace' }}
                        >
                          {log.timestamp.toLocaleTimeString()}
                        </Typography>
                        <Typography
                          variant="caption"
                          color="text.secondary"
                          sx={{ minWidth: 80, fontFamily: 'monospace' }}
                        >
                          [{log.streamType}]
                        </Typography>
                        <Typography
                          variant="body2"
                          component="pre"
                          sx={{
                            margin: 0,
                            whiteSpace: 'pre-wrap',
                            wordBreak: 'break-word',
                            flex: 1,
                          }}
                        >
                          {log.content}
                        </Typography>
                      </Box>
                    </LogEntry>
                  ))
                )}
              </LogContainer>
            </Box>
          ) : (
            <Box>
              {stackJob ? (
                <JsonCode content={stackJob?.output || {}} />
              ) : (
                <Skeleton variant="rounded" width={'100%'} height={200} />
              )}
            </Box>
          )}
        </Paper>

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
