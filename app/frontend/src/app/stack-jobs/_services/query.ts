import { useContext, useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { StackJobService } from '@/gen/proto/stack_job_service_pb';
import {
  ListStackJobsRequest,
  ListStackJobsResponse,
  GetStackJobRequestSchema,
  StackJob,
} from '@/gen/proto/stack_job_service_pb';

interface QueryType {
  listStackJobs: (input: ListStackJobsRequest) => Promise<ListStackJobsResponse>;
  getById: (id: string) => Promise<StackJob>;
}

const RESOURCE_NAME = 'Stack Jobs';

export const useStackJobQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(StackJobService);
  const [query, setQuery] = useState<QueryType>(null);

  const stackJobQuery: QueryType = useMemo(
    () => ({
      listStackJobs: (input: ListStackJobsRequest): Promise<ListStackJobsResponse> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .listStackJobs(input)
            .then((response) => {
              resolve(response);
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not get ${RESOURCE_NAME}!`, 'error');
              reject(err);
            })
            .finally(() => {
              setPageLoading(false);
            });
        });
      },
      getById: (id: string): Promise<StackJob> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .getStackJob(create(GetStackJobRequestSchema, { id }))
            .then((response) => {
              resolve(response.job);
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not get ${RESOURCE_NAME}!`, 'error');
              reject(err);
            })
            .finally(() => {
              setPageLoading(false);
            });
        });
      },
    }),
    [queryClient, setPageLoading, openSnackbar]
  );

  useEffect(() => {
    if (queryClient && !query) {
      setQuery(stackJobQuery);
    }
  }, [queryClient, stackJobQuery, query]);

  return { query };
};
