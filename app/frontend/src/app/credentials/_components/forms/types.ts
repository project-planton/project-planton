'use client';

import { CredentialProvider, AwsCredentialSpec, AzureCredentialSpec, GcpCredentialSpec } from '@/gen/proto/credential_service_pb';

// Form-friendly type based on CreateCredentialRequest fields (without the protobuf Message wrapper)
export type CredentialFormData = {
  name: string;
  provider: CredentialProvider;
  gcp?: Partial<GcpCredentialSpec>;
  aws?: Partial<AwsCredentialSpec>;
  azure?: Partial<AzureCredentialSpec>;
};

