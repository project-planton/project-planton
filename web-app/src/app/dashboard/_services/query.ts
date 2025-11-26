import { useContext, useEffect, useMemo, useState } from 'react';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { DeploymentComponentService } from '@/gen/proto/deployment_component_service_connect';
import type { DescService } from '@bufbuild/protobuf';
import { ListDeploymentComponentsRequest } from '@/gen/proto/deployment_component_service_pb';

interface QueryType {
  listDeploymentComponents: (input: ListDeploymentComponentsRequest) => Promise<any>;
}

const RESOURCE_NAME = 'Deployment Components';

export const useDashboardQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(
    DeploymentComponentService as unknown as DescService
  ) as any;
  const [query, setQuery] = useState<QueryType | null>(null);

  const queryApis: QueryType = useMemo(
    () => ({
      listDeploymentComponents: (input: ListDeploymentComponentsRequest): Promise<any> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);

          queryClient
            .listDeploymentComponents(input)
            .then(resolve)
            .catch((err: any) => {
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
