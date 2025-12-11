'use client';

import { SimpleInput } from '@/components/shared/simple-input';
import { CredentialFormData } from '@/app/credentials/_components/forms/types';
import { UseFormRegister } from 'react-hook-form';

interface AwsCredentialFormProps {
  register: UseFormRegister<CredentialFormData>;
  disabled?: boolean;
}

export function AwsCredentialForm({ register, disabled }: AwsCredentialFormProps) {
  return (
    <>
      <SimpleInput
        register={register}
        path="aws.accountId"
        name="Account ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="aws.accessKeyId"
        name="Access Key ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="aws.secretAccessKey"
        name="Secret Access Key"
        type="password"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="aws.region"
        name="Region (Optional)"
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="aws.sessionToken"
        name="Session Token (Optional)"
        type="password"
        disabled={disabled}
      />
    </>
  );
}

