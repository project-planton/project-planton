import { useContext, useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { StackUpdateService } from '@/gen/proto/stack_job_service_pb';
import {
  ListStackUpdatesRequest,
  ListStackUpdatesResponse,
  GetStackUpdateRequestSchema,
  StreamStackUpdateOutputRequestSchema,
  StreamStackUpdateOutputResponse,
  StackUpdate,
} from '@/gen/proto/stack_job_service_pb';

interface QueryType {
  listStackUpdates: (input: ListStackUpdatesRequest) => Promise<ListStackUpdatesResponse>;
  getById: (id: string) => Promise<StackUpdate>;
  streamOutput: (
    jobId: string,
    lastSequenceNum?: number,
    signal?: AbortSignal
  ) => AsyncIterable<StreamStackUpdateOutputResponse>;
}

const RESOURCE_NAME = 'Stack Jobs';

export const useStackUpdateQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(StackUpdateService);
  const [query, setQuery] = useState<QueryType>(null);

  const stackUpdateQuery: QueryType = useMemo(
    () => ({
      listStackUpdates: (input: ListStackUpdatesRequest): Promise<ListStackUpdatesResponse> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .listStackUpdates(input)
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
      getById: (id: string): Promise<StackUpdate> => {
        return new Promise((resolve, reject) => {
          queryClient
            .getStackUpdate(create(GetStackUpdateRequestSchema, { id }))
            .then((response) => {
              resolve(response.job);
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not get ${RESOURCE_NAME}!`, 'error');
              reject(err);
            });
        });
      },
      streamOutput: (
        jobId: string,
        lastSequenceNum?: number,
        signal?: AbortSignal
      ): AsyncIterable<StreamStackUpdateOutputResponse> => {
        const request = create(StreamStackUpdateOutputRequestSchema, {
          jobId,
          lastSequenceNum: lastSequenceNum !== undefined ? lastSequenceNum : undefined,
        });

        return queryClient.streamStackUpdateOutput(request, { signal });
      },
    }),
    [queryClient, setPageLoading, openSnackbar]
  );

  useEffect(() => {
    if (queryClient && !query) {
      setQuery(stackUpdateQuery);
    }
  }, [queryClient, stackUpdateQuery, query]);

  return { query };
};
