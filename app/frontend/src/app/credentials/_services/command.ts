import { useContext, useEffect, useMemo, useState } from 'react';
import { create } from '@bufbuild/protobuf';
import { AppContext } from '@/contexts';
import { useConnectRpcClient } from '@/hooks';
import { CredentialService } from '@/gen/proto/credential_service_pb';
import {
  CreateCredentialRequest,
  CreateCredentialRequestSchema,
  CreateCredentialResponse,
  UpdateCredentialRequest,
  UpdateCredentialRequestSchema,
  UpdateCredentialResponse,
  DeleteCredentialRequestSchema,
  DeleteCredentialResponse,
  Credential,
  CredentialProvider,
} from '@/gen/proto/credential_service_pb';

interface CommandType {
  create: (
    name: string,
    provider: CredentialProvider,
    credentialData: CreateCredentialRequest['credentialData']
  ) => Promise<Credential>;
  update: (
    id: string,
    name: string,
    provider: CredentialProvider,
    credentialData: UpdateCredentialRequest['credentialData']
  ) => Promise<Credential>;
  delete: (id: string) => Promise<void>;
}

const RESOURCE_NAME = 'Credential';

export const useCredentialCommand = () => {
  const { setPageLoading, openSnackbar } = useContext(AppContext);
  const commandClient = useConnectRpcClient(CredentialService);
  const [command, setCommand] = useState<CommandType | null>(null);

  const commandApis: CommandType = useMemo(
    () => ({
      create: (
        name: string,
        provider: CredentialProvider,
        credentialData: CreateCredentialRequest['credentialData']
      ): Promise<Credential> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          commandClient
            .createCredential(
              create(CreateCredentialRequestSchema, {
                name,
                provider,
                credentialData,
              })
            )
            .then((response: CreateCredentialResponse) => {
              if (response?.credential) {
                openSnackbar(
                  `${RESOURCE_NAME} ${response.credential.name} created successfully`,
                  'success'
                );
                resolve(response.credential);
              }
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not create ${RESOURCE_NAME}`, 'error');
              reject(err);
            })
            .finally(() => setPageLoading(false));
        });
      },
      update: (
        id: string,
        name: string,
        provider: CredentialProvider,
        credentialData: UpdateCredentialRequest['credentialData']
      ): Promise<Credential> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          commandClient
            .updateCredential(
              create(UpdateCredentialRequestSchema, {
                id,
                name,
                provider,
                credentialData,
              })
            )
            .then((response: UpdateCredentialResponse) => {
              if (response?.credential) {
                openSnackbar(
                  `${RESOURCE_NAME} ${response.credential.name} updated successfully`,
                  'success'
                );
              }
              resolve(response.credential!);
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not update ${RESOURCE_NAME}`, 'error');
              reject(err);
            })
            .finally(() => setPageLoading(false));
        });
      },
      delete: (id: string): Promise<void> => {
        return new Promise((resolve, reject) => {
          setPageLoading(true);
          commandClient
            .deleteCredential(create(DeleteCredentialRequestSchema, { id }))
            .then((response: DeleteCredentialResponse) => {
              openSnackbar(
                response.message || `${RESOURCE_NAME} deleted successfully`,
                'success'
              );
              resolve();
            })
            .catch((err) => {
              openSnackbar(err.message || `Could not delete ${RESOURCE_NAME}`, 'error');
              reject(err);
            })
            .finally(() => setPageLoading(false));
        });
      },
    }),
    [commandClient]
  );

  useEffect(() => {
    if (commandClient && !command) {
      setCommand(commandApis);
    }
  }, [commandApis]);

  return { command };
};

