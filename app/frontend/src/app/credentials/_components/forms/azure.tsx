'use client';

import { SimpleInput } from '@/components/shared/simple-input';
import { CredentialFormData } from '@/app/credentials/_components/forms/types';
import { UseFormRegister } from 'react-hook-form';

interface AzureCredentialFormProps {
  register: UseFormRegister<CredentialFormData>;
  disabled?: boolean;
}

export function AzureCredentialForm({ register, disabled }: AzureCredentialFormProps) {
  return (
    <>
      <SimpleInput
        register={register}
        path="azure.clientId"
        name="Client ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="azure.clientSecret"
        name="Client Secret"
        type="password"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="azure.tenantId"
        name="Tenant ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="azure.subscriptionId"
        name="Subscription ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
    </>
  );
}

