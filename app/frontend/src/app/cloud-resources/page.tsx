'use client';
import { Typography } from '@mui/material';
import { CloudResourceContainer } from '@/app/cloud-resources/styled';
import { CloudResourcesList } from '@/components/shared/cloud-resources-list';

export default function CloudResourcesPage() {
  return (
    <CloudResourceContainer>
      <Typography variant="h4" gutterBottom>
        Cloud Resources
      </Typography>

      <CloudResourcesList
        title="Cloud Resources"
        showKindFilter={true}
        showErrorAlerts={false}
      />
    </CloudResourceContainer>
  );
}
