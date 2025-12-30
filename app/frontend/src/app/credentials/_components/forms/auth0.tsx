'use client';

import { SimpleInput } from '@/components/shared/simple-input';
import { CredentialFormData } from '@/app/credentials/_components/forms/types';
import { UseFormRegister } from 'react-hook-form';

interface Auth0CredentialFormProps {
  register: UseFormRegister<CredentialFormData>;
  disabled?: boolean;
}

export function Auth0CredentialForm({ register, disabled }: Auth0CredentialFormProps) {
  return (
    <>
      <SimpleInput
        register={register}
        path="auth0.domain"
        name="Domain"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="auth0.clientId"
        name="Client ID"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
      <SimpleInput
        register={register}
        path="auth0.clientSecret"
        name="Client Secret"
        type="password"
        registerOptions={{ required: true }}
        disabled={disabled}
      />
    </>
  );
}

