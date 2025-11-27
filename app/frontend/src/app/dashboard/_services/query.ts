import { useContext, useEffect, useMemo, useState } from 'react';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { DeploymentComponentService } from '@/gen/proto/deployment_component_service_pb';
import {
  ListDeploymentComponentsRequest,
  ListDeploymentComponentsResponse,
} from '@/gen/proto/deployment_component_service_pb';

interface QueryType {
  listDeploymentComponents: (
    input: ListDeploymentComponentsRequest
  ) => Promise<ListDeploymentComponentsResponse>;
}

const RESOURCE_NAME = 'Deployment Components';

export const useDashboardQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(DeploymentComponentService);
  const [query, setQuery] = useState<QueryType | null>(null);

  const queryApis: QueryType = useMemo(
    () => ({
      listDeploymentComponents: (
        input: ListDeploymentComponentsRequest
      ): Promise<ListDeploymentComponentsResponse> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);

          queryClient
            .listDeploymentComponents(input)
            .then((response) => {
              // With binary format and proper schemas, Connect RPC automatically deserializes
              resolve(response);
            })
            .catch((err: any) => {
              console.error('RPC Error:', err);
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
      setQuery(queryApis);
    }
  }, [queryClient, queryApis, query]);

  return { query };
};
