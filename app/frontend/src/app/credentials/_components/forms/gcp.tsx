'use client';

import { SimpleInput } from '@/components/shared/simple-input';
import { CredentialFormData } from '@/app/credentials/_components/forms/types';
import { UseFormRegister } from 'react-hook-form';

interface GcpCredentialFormProps {
  register: UseFormRegister<CredentialFormData>;
  disabled?: boolean;
}

export function GcpCredentialForm({ register, disabled }: GcpCredentialFormProps) {
  return (
    <SimpleInput
      register={register}
      path="gcp.serviceAccountKeyBase64"
      name="Service Account Key (Base64)"
      multiline
      rows={8}
      registerOptions={{ required: true }}
      disabled={disabled}
    />
  );
}
