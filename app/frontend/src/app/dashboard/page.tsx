'use client';
import { useState, useEffect, useCallback } from 'react';
import { create } from '@bufbuild/protobuf';
import { Typography, Grid2, Skeleton } from '@mui/material';
import {
  DashboardContainer,
  StyledGrid2,
  StyledCard,
  StatCardTitle,
  StatCardValue,
} from '@/app/dashboard/styled';
import { CloudResourcesList } from '@/components/shared/cloud-resources-list';
import { useCloudResourceQuery } from '@/app/cloud-resources/_services';
import { ListCloudResourcesRequestSchema } from '@/gen/proto/cloud_resource_service_pb';

export default function DashboardPage() {
  const { query } = useCloudResourceQuery();
  const [cloudResourceCount, setCloudResourceCount] = useState<number | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const getCloudResourceCount = useCallback(() => {
    if (query) {
      setIsLoading(true);
      query
        .listCloudResources(create(ListCloudResourcesRequestSchema, {}))
        .then((response) => {
          setCloudResourceCount(response.resources.length);
          setIsLoading(false);
        })
        .catch(() => {
          setIsLoading(false);
        });
    }
  }, [query]);

  useEffect(() => {
    getCloudResourceCount();
  }, [getCloudResourceCount]);

  return (
    <DashboardContainer>
      <Typography variant="h4" gutterBottom>
        Dashboard
      </Typography>

      <StyledGrid2 container spacing={3}>
        <Grid2 size={{ xs: 12, sm: 6, md: 4 }}>
          <StyledCard>
            <StatCardTitle>Cloud Resources</StatCardTitle>
            <StatCardValue>
              {isLoading ? (
                <Skeleton variant="text" width={50} height={24} />
              ) : cloudResourceCount !== null ? (
                cloudResourceCount
              ) : (
                0
              )}
            </StatCardValue>
          </StyledCard>
        </Grid2>
      </StyledGrid2>

      <CloudResourcesList
        title="Cloud Resources"
        showErrorAlerts={true}
        onChange={getCloudResourceCount}
      />
    </DashboardContainer>
  );
}
