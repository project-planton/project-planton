import { useContext, useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { CloudResourceService } from '@/gen/proto/cloud_resource_service_pb';
import {
  ListCloudResourcesRequest,
  ListCloudResourcesResponse,
  GetCloudResourceRequestSchema,
  CountCloudResourcesRequestSchema,
  CountCloudResourcesResponse,
  CloudResource,
} from '@/gen/proto/cloud_resource_service_pb';

interface QueryType {
  listCloudResources: (input: ListCloudResourcesRequest) => Promise<ListCloudResourcesResponse>;
  getById: (id: string) => Promise<CloudResource>;
  count: (kind?: string) => Promise<number>;
}

const RESOURCE_NAME = 'Cloud Resources';

export const useCloudResourceQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(CloudResourceService);
  const [query, setQuery] = useState<QueryType>(null);

  const cloudResourceQuery: QueryType = useMemo(
    () => ({
      listCloudResources: (
        input: ListCloudResourcesRequest
      ): Promise<ListCloudResourcesResponse> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          console.log('input', input);
          queryClient
            .listCloudResources(input)
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
      getById: (id: string): Promise<CloudResource> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .getCloudResource(create(GetCloudResourceRequestSchema, { id }))
            .then((response) => {
              resolve(response.resource);
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
      count: (kind?: string): Promise<number> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .countCloudResources(
              create(CountCloudResourcesRequestSchema, {
                kind: kind || undefined,
              })
            )
            .then((response: CountCloudResourcesResponse) => {
              resolve(Number(response.count));
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not count ${RESOURCE_NAME}!`, 'error');
              reject(err);
            })
            .finally(() => {
              setPageLoading(false);
            });
        });
      },
    }),
    [queryClient]
  );

  useEffect(() => {
    if (queryClient && !query) {
      setQuery(cloudResourceQuery);
    }
  }, [cloudResourceQuery]);

  return { query };
};
