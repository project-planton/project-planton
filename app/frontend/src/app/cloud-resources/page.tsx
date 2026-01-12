'use client';
import { Box, Stack } from '@mui/material';
import { CloudResourcesList } from '@/components/shared/cloud-resources-list';
import { SectionHeader } from '@/components/shared/section-header';

export default function CloudResourcesPage() {
  return (
    <Stack height={'100%'}>
      <SectionHeader
        title="Cloud Resources"
        borderBottom
        containerProps={{ paddingX: 4, paddingY: 3 }}
      />
      <Box bgcolor="grey.20" height={'100%'} px={4} mt={4}>
        <CloudResourcesList title="" showKindFilter={true} showErrorAlerts={false} />
      </Box>
    </Stack>
  );
}
