import { useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { useConnectRpcClient } from '@/hooks';
import { StreamingService } from '@/gen/proto/streaming_service_pb';
import { StreamDataRequestSchema, StreamDataResponse } from '@/gen/proto/streaming_service_pb';

interface QueryType {
  streamData: (
    options?: { messageCount?: number; intervalMs?: number },
    signal?: AbortSignal
  ) => AsyncIterable<StreamDataResponse>;
}

export const useStreamingQuery = () => {
  const queryClient = useConnectRpcClient(StreamingService);
  const [query, setQuery] = useState<QueryType>(null);

  const serviceQuery: QueryType = useMemo(
    () => ({
      streamData: (
        options?: { messageCount?: number; intervalMs?: number },
        signal?: AbortSignal
      ): AsyncIterable<StreamDataResponse> => {
        const request = create(StreamDataRequestSchema, {
          messageCount: options?.messageCount,
          intervalMs: options?.intervalMs,
        });

        return queryClient.streamData(request, { signal });
      },
    }),
    [queryClient]
  );

  useEffect(() => {
    if (queryClient && !query) {
      setQuery(serviceQuery);
    }
  }, [serviceQuery]);

  return { query };
};
