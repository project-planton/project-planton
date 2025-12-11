'use client';

import { useEffect, useCallback, useMemo } from 'react';
import { useForm, useWatch } from 'react-hook-form';
import { Drawer } from '@/components/shared/drawer';
import { Stack, Button } from '@mui/material';
import {
  DrawerContainer,
  DrawerContentArea,
  DrawerFooter,
} from '@/app/credentials/_components/styled';
import { SimpleInput } from '@/components/shared/simple-input';
import { SimpleSelect } from '@/components/shared/simple-select';
import {
  CredentialFormData,
  GcpCredentialForm,
  AwsCredentialForm,
  AzureCredentialForm,
} from '@/app/credentials/_components/forms';
import { useCredentialCommand } from '@/app/credentials/_services';
import {
  CredentialProvider,
  Credential,
  CreateCredentialRequest,
  GcpCredentialSpec,
  AwsCredentialSpec,
  AzureCredentialSpec,
  GcpCredentialSpecSchema,
  AwsCredentialSpecSchema,
  AzureCredentialSpecSchema,
} from '@/gen/proto/credential_service_pb';
import { create } from '@bufbuild/protobuf';
import { providerConfig } from '@/app/credentials/_components/utils';

export type DrawerMode = 'view' | 'edit' | 'create' | null;

interface CredentialDrawerProps {
  open: boolean;
  mode: DrawerMode;
  onClose: () => void;
  onSaveSuccess: () => void;
  selectedCredential?: Credential | null;
  initialProvider?: CredentialProvider;
}

export function CredentialDrawer({
  open,
  mode,
  onClose,
  onSaveSuccess,
  selectedCredential,
  initialProvider,
}: CredentialDrawerProps) {
  const { command } = useCredentialCommand();
  const isView = mode === 'view';
  const submitLabel = mode === 'edit' ? 'Update' : 'Create';

  const { register, handleSubmit, reset, setValue, control } = useForm<CredentialFormData>({
    defaultValues: {
      name: '',
      provider: CredentialProvider.CREDENTIAL_PROVIDER_UNSPECIFIED,
      gcp: {},
      aws: {},
      azure: {},
    },
  });

  useEffect(() => {
    if (initialProvider) {
      console.log('initialProvider', initialProvider);
      setValue('provider', initialProvider);
    }
  }, [initialProvider, setValue]);

  const formProvider = useWatch({ control, name: 'provider' });

  const providerOptions = useMemo(() => {
    return (Object.keys(providerConfig) as unknown as Array<CredentialProvider>)
      .filter((provider) => {
        // Filter out UNSPECIFIED (value 0) by comparing numeric enum values
        return Number(provider) !== CredentialProvider.CREDENTIAL_PROVIDER_UNSPECIFIED;
      })
      .map((provider) => ({
        label: providerConfig[provider].label,
        value: provider,
      }));
  }, []);

  // Populate form when selectedCredential changes
  useEffect(() => {
    if (selectedCredential && (mode === 'view' || mode === 'edit')) {
      const credentialData = selectedCredential.credentialData;
      const formData: CredentialFormData = {
        name: selectedCredential.name,
        provider: selectedCredential.provider,
        gcp: {},
        aws: {},
        azure: {},
      };
      if (credentialData?.case === 'gcp') {
        formData.gcp = {
          serviceAccountKeyBase64: credentialData.value.serviceAccountKeyBase64,
        };
      } else if (credentialData?.case === 'aws') {
        formData.aws = {
          accountId: credentialData.value.accountId,
          accessKeyId: credentialData.value.accessKeyId,
          secretAccessKey: credentialData.value.secretAccessKey,
          region: credentialData.value.region,
          sessionToken: credentialData.value.sessionToken,
        };
      } else if (credentialData?.case === 'azure') {
        formData.azure = {
          clientId: credentialData.value.clientId,
          clientSecret: credentialData.value.clientSecret,
          tenantId: credentialData.value.tenantId,
          subscriptionId: credentialData.value.subscriptionId,
        };
      }
      reset(formData);
    } else if (mode === 'create') {
      reset({
        name: '',
        provider: initialProvider || CredentialProvider.CREDENTIAL_PROVIDER_UNSPECIFIED,
        gcp: {},
        aws: {},
        azure: {},
      });
    }
  }, [selectedCredential, mode, initialProvider, reset]);

  const handleSave = useCallback(
    (formData: CredentialFormData) => {
      if (!command) return;

      let credentialData: CreateCredentialRequest['credentialData'];

      if (formData.provider == CredentialProvider.GCP && formData.gcp?.serviceAccountKeyBase64) {
        credentialData = {
          case: 'gcp',
          value: create(GcpCredentialSpecSchema, formData.gcp as GcpCredentialSpec),
        };
      } else if (
        formData.provider == CredentialProvider.AWS &&
        formData.aws?.accountId &&
        formData.aws?.accessKeyId &&
        formData.aws?.secretAccessKey
      ) {
        credentialData = {
          case: 'aws',
          value: create(AwsCredentialSpecSchema, formData.aws as AwsCredentialSpec),
        };
      } else if (
        formData.provider == CredentialProvider.AZURE &&
        formData.azure?.clientId &&
        formData.azure?.clientSecret &&
        formData.azure?.tenantId &&
        formData.azure?.subscriptionId
      ) {
        credentialData = {
          case: 'azure',
          value: create(AzureCredentialSpecSchema, formData.azure as AzureCredentialSpec),
        };
      } else {
        return;
      }

      if (mode === 'create') {
        command.create(formData.name, formData.provider, credentialData).then(() => {
          onSaveSuccess();
        });
      } else if (mode === 'edit' && selectedCredential) {
        command
          .update(selectedCredential.id, formData.name, formData.provider, credentialData)
          .then(() => {
            onSaveSuccess();
          });
      }
    },
    [command, mode, selectedCredential, onSaveSuccess]
  );

  const handleClose = () => {
    reset({
      name: '',
      provider: initialProvider || CredentialProvider.CREDENTIAL_PROVIDER_UNSPECIFIED,
      gcp: {},
      aws: {},
      azure: {},
    });
    onClose();
  };

  const onProviderChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
      if (isView || initialProvider) return;
      const newProvider = parseInt(e.target.value, 10) as CredentialProvider;
      setValue('provider', newProvider);
      setValue('gcp', {});
      setValue('aws', {});
      setValue('azure', {});
    },
    [setValue, isView, initialProvider]
  );

  const title =
    mode === 'view'
      ? 'View Credential'
      : mode === 'edit'
        ? 'Edit Credential'
        : initialProvider
          ? `Create ${providerConfig[initialProvider].label} Credential`
          : 'Create Credential';

  return (
    <Drawer open={open} onClose={handleClose} title={title} width={600}>
      <DrawerContainer>
        <DrawerContentArea $hasFooter={!isView}>
          <Stack spacing={3}>
            <Stack>
              <SimpleSelect
                name="Provider"
                value={formProvider}
                required
                disabled={isView || !!initialProvider}
                onChange={onProviderChange}
                options={providerOptions}
                sx={{ minWidth: 250 }}
              />
              {!!formProvider && (
                <SimpleInput
                  register={register}
                  path="name"
                  name="Name"
                  registerOptions={{ required: true }}
                  disabled={isView}
                />
              )}
              {formProvider == CredentialProvider.GCP && (
                <GcpCredentialForm register={register} disabled={isView} />
              )}
              {formProvider == CredentialProvider.AWS && (
                <AwsCredentialForm register={register} disabled={isView} />
              )}
              {formProvider == CredentialProvider.AZURE && (
                <AzureCredentialForm register={register} disabled={isView} />
              )}
            </Stack>
          </Stack>
        </DrawerContentArea>
        {!isView && (
          <DrawerFooter>
            <Button variant="contained" color="secondary" onClick={handleClose}>
              Cancel
            </Button>
            <Button variant="contained" color="primary" onClick={handleSubmit(handleSave)}>
              {submitLabel}
            </Button>
          </DrawerFooter>
        )}
      </DrawerContainer>
    </Drawer>
  );
}
