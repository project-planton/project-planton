'use client';
import { CloudResourceContainer } from '@/app/cloud-resources/styled';
import { CloudResourcesList } from '@/components/shared/cloud-resources-list';

export default function CloudResourcesPage() {
  return (
    <CloudResourceContainer>
      <CloudResourcesList title="Cloud Resources" showKindFilter={true} showErrorAlerts={false} />
    </CloudResourceContainer>
  );
}
