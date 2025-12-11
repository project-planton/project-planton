import { useContext, useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { CredentialService } from '@/gen/proto/credential_service_pb';
import {
  ListCredentialsRequest,
  ListCredentialsResponse,
  GetCredentialRequestSchema,
  Credential,
} from '@/gen/proto/credential_service_pb';

interface QueryType {
  listCredentials: (input: ListCredentialsRequest) => Promise<ListCredentialsResponse>;
  getById: (id: string) => Promise<Credential>;
}

const RESOURCE_NAME = 'Credentials';

export const useCredentialQuery = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const queryClient = useConnectRpcClient(CredentialService);
  const [query, setQuery] = useState<QueryType | null>(null);

  const credentialQuery: QueryType = useMemo(
    () => ({
      listCredentials: (input: ListCredentialsRequest): Promise<ListCredentialsResponse> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .listCredentials(input)
            .then(resolve)
            .catch((err) => {
              openSnackbar(err.message || `Could not get ${RESOURCE_NAME}!`, 'error');
              reject(err);
            })
            .finally(() => {
              setPageLoading(false);
            });
        });
      },
      getById: (id: string): Promise<Credential> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          queryClient
            .getCredential(create(GetCredentialRequestSchema, { id }))
            .then((response) => {
              resolve(response.credential!);
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
    [queryClient]
  );

  useEffect(() => {
    if (queryClient && !query) {
      setQuery(credentialQuery);
    }
  }, [credentialQuery]);

  return { query };
};

