'use client';

import { Credential_CredentialProvider } from '@/gen/org/project_planton/app/credential/v1/api_pb';
import { GcpProviderConfig } from '@/gen/org/project_planton/provider/gcp/provider_pb';
import { AwsProviderConfig } from '@/gen/org/project_planton/provider/aws/provider_pb';
import { AzureProviderConfig } from '@/gen/org/project_planton/provider/azure/provider_pb';

// Form-friendly type based on CreateCredentialRequest fields (without the protobuf Message wrapper)
export type CredentialFormData = {
  name: string;
  provider: Credential_CredentialProvider;
  gcp?: Partial<GcpProviderConfig>;
  aws?: Partial<AwsProviderConfig>;
  azure?: Partial<AzureProviderConfig>;
};

