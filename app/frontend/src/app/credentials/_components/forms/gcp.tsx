'use client';

import { CredentialFormData } from '@/app/credentials/_components/forms/types';
import { UseFormSetValue, UseFormWatch } from 'react-hook-form';
import { FileUploadWithClear } from '@/components/shared/file-upload';
import { FormField } from '@/components/shared/form-field';
import { SecretProperty } from '@/components/shared/secret-property';
import { Stack, Typography } from '@mui/material';
import { resolveSecretKey } from '@/lib';

interface GcpCredentialFormProps {
  setValue: UseFormSetValue<CredentialFormData>;
  watch?: UseFormWatch<CredentialFormData>;
  disabled?: boolean;
  credentialName?: string;
}

export function GcpCredentialForm({ setValue, watch, disabled, credentialName }: GcpCredentialFormProps) {
  const currentValue = watch ? watch('gcp.serviceAccountKeyBase64') : undefined;
  const isViewMode = disabled && currentValue;

  return (
    <Stack>
      <FormField fullWidth>
        {isViewMode ? (
          <Stack gap={1}>
            <Typography variant="subtitle2" color="text.secondary">
              Service Account Key
            </Typography>
            <SecretProperty
              property="Service Account Key"
              value="*****************"
              getSecretValue={async () => {
                // currentValue is the base64 encoded string, resolveSecretKey will decode it
                return await resolveSecretKey(currentValue);
              }}
              enableDownload={true}
              downloadFileName={credentialName || 'service-account-key'}
            />
          </Stack>
        ) : (
          <FileUploadWithClear
            buttonText="Upload Key File"
            maxSizeBytes={102400}
            setValue={setValue}
            path="gcp.serviceAccountKeyBase64"
            watch={watch}
            disabled={disabled}
            downloadFileName={credentialName || 'service-account-key'}
          />
        )}
      </FormField>
    </Stack>
  );
}
