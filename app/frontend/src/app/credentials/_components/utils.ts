import { CredentialProvider } from '@/gen/proto/credential_service_pb';
import { ICON_NAMES } from '@/components/shared/icon';

export interface ProviderConfig {
  label: string;
  description: string;
  icon?: ICON_NAMES;
}

export const providerConfig: Record<CredentialProvider, ProviderConfig> = {
  [CredentialProvider.CREDENTIAL_PROVIDER_UNSPECIFIED]: {
    label: 'Unspecified',
    description: '',
    icon: undefined,
  },
  [CredentialProvider.GCP]: {
    label: 'Google Cloud Platform',
    description: 'Link your GCP Organization to start deploying infrastructure',
    icon: ICON_NAMES.GCP,
  },
  [CredentialProvider.AWS]: {
    label: 'Amazon Web Services',
    description: 'Link your AWS Account to start deploying infrastructure',
    icon: ICON_NAMES.AWS,
  },
  [CredentialProvider.AZURE]: {
    label: 'Microsoft Azure',
    description: 'Link your Azure Subscription to start deploying infrastructure',
    icon: ICON_NAMES.AZURE,
  },
};

