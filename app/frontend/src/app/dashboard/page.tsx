'use client';
import { useState, useMemo, useEffect, useCallback } from 'react';
import { Typography, Grid2, Button, Box, Alert } from '@mui/material';
import { create } from '@bufbuild/protobuf';
import {
  DeploymentComponent,
  ListDeploymentComponentsRequestSchema,
} from '@/gen/proto/deployment_component_service_pb';
import {
  DashboardContainer,
  StyledGrid2,
  StyledPaper,
  TableSection,
  StatCardTitle,
  StatCardValue,
} from '@/app/dashboard/styled';
import { DataTable, Column, Action } from '@/components/shared/data-table';
import { Edit, Delete, Visibility, Refresh } from '@mui/icons-material';
import { StatusBadge } from '@/components/shared/data-table/styled';
import { useDashboardQuery } from '@/app/dashboard/_services/query';
import { Timestamp } from '@bufbuild/protobuf/wkt';
import { formatTimestampToDate } from '@/lib';

// Removed SampleData interface - using DeploymentComponent only

export default function DashboardPage() {
  const { query } = useDashboardQuery();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [selectedRows, setSelectedRows] = useState<DeploymentComponent[]>([]);
  const [sortColumn, setSortColumn] = useState<keyof DeploymentComponent | string>('name');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc');
  const [deploymentComponents, setDeploymentComponents] = useState<DeploymentComponent[]>([]);
  const [apiError, setApiError] = useState<string | null>(null);
  const [apiLoading, setApiLoading] = useState(false);

  // No sample data - using deployment components from API

  const columns: Column<DeploymentComponent>[] = useMemo(
    () => [
      { id: 'name', label: 'Name', sortable: true },
      { id: 'kind', label: 'Kind', sortable: true },
      { id: 'provider', label: 'Provider', sortable: true },
      { id: 'version', label: 'Version', sortable: true },
      { id: 'idPrefix', label: 'ID Prefix', sortable: true },
      {
        id: 'isServiceKind',
        label: 'Service Kind',
        sortable: true,
        render: (value: boolean) => (
          <StatusBadge $status={value ? 'active' : 'inactive'}>{value ? 'Yes' : 'No'}</StatusBadge>
        ),
      },
      {
        id: 'createdAt',
        label: 'Created At',
        sortable: true,
        render: (value: Timestamp) => formatTimestampToDate(value, 'DD/MM/YYYY'),
      },
    ],
    []
  );

  const actions: Action<DeploymentComponent>[] = useMemo(
    () => [
      {
        label: 'View',
        icon: <Visibility fontSize="small" />,
        onClick: (row) => {
          console.log('View deployment component:', row);
          // Handle view action
        },
        color: 'primary',
      },
      {
        label: 'Edit',
        icon: <Edit fontSize="small" />,
        onClick: (row) => {
          console.log('Edit deployment component:', row);
          // Handle edit action
        },
        color: 'primary',
      },
      {
        label: 'Delete',
        icon: <Delete fontSize="small" />,
        onClick: (row) => {
          console.log('Delete deployment component:', row);
          // Handle delete action
        },
        color: 'error',
      },
    ],
    []
  );

  const handleSelectAll = (selected: boolean) => {
    if (selected) {
      setSelectedRows(deploymentComponents);
    } else {
      setSelectedRows([]);
    }
  };

  const handleSelectRow = (row: DeploymentComponent, selected: boolean) => {
    if (selected) {
      setSelectedRows([...selectedRows, row]);
    } else {
      setSelectedRows(selectedRows.filter((r) => r.id !== row.id));
    }
  };

  const handleSort = (columnId: keyof DeploymentComponent | string, order: 'asc' | 'desc') => {
    setSortColumn(columnId);
    setSortOrder(order);
  };

  const handlePageChange = (newPage: number) => {
    setPage(newPage);
  };

  const handleRowsPerPageChange = (newRowsPerPage: number) => {
    setRowsPerPage(newRowsPerPage);
    setPage(0);
  };

  // Function to call the API
  const handleLoadDeploymentComponents = useCallback(() => {
    if (!query) {
      setApiError('Query service not ready');
      return;
    }

    setApiLoading(true);
    setApiError(null);
    query
      .listDeploymentComponents(create(ListDeploymentComponentsRequestSchema))
      .then((result) => {
        setDeploymentComponents(result.components);
        setApiLoading(false);
      })
      .catch((err: any) => {
        setApiError(err.message || 'Failed to load deployment components');
        console.error('API Error:', err);
        setApiLoading(false);
      });
  }, [query]);

  // Auto-load on mount
  useEffect(() => {
    if (query) {
      handleLoadDeploymentComponents();
    }
  }, [query, handleLoadDeploymentComponents]);

  // Apply sorting
  const sortedData = useMemo(() => {
    const sorted = [...deploymentComponents];
    sorted.sort((a, b) => {
      const aValue = a[sortColumn as keyof DeploymentComponent];
      const bValue = b[sortColumn as keyof DeploymentComponent];
      if (sortOrder === 'asc') {
        return aValue > bValue ? 1 : -1;
      }
      return aValue < bValue ? 1 : -1;
    });
    return sorted;
  }, [deploymentComponents, sortColumn, sortOrder]);

  // Apply pagination
  const paginatedData = useMemo(() => {
    const start = page * rowsPerPage;
    return sortedData.slice(start, start + rowsPerPage);
  }, [sortedData, page, rowsPerPage]);

  return (
    <DashboardContainer>
      <Typography variant="h4" gutterBottom>
        Dashboard
      </Typography>

      <StyledGrid2 container spacing={3}>
        <Grid2 size={{ xs: 12, sm: 6, md: 4 }}>
          <StyledPaper>
            <StatCardTitle>Total Products</StatCardTitle>
            <StatCardValue>51</StatCardValue>
          </StyledPaper>
        </Grid2>
        <Grid2 size={{ xs: 12, sm: 6, md: 4 }}>
          <StyledPaper>
            <StatCardTitle>Product Inventory</StatCardTitle>
            <StatCardValue>290</StatCardValue>
          </StyledPaper>
        </Grid2>
        <Grid2 size={{ xs: 12, sm: 6, md: 4 }}>
          <StyledPaper>
            <StatCardTitle>Average price</StatCardTitle>
            <StatCardValue>2,652.79</StatCardValue>
          </StyledPaper>
        </Grid2>
      </StyledGrid2>

      <TableSection>
        <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
          <Typography variant="h5">Deployment Components</Typography>
          <Button
            variant="contained"
            startIcon={<Refresh />}
            onClick={handleLoadDeploymentComponents}
            disabled={apiLoading || !query}
          >
            {apiLoading ? 'Loading...' : 'Refresh'}
          </Button>
        </Box>

        {apiError && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {apiError}
          </Alert>
        )}

        {deploymentComponents.length === 0 && !apiLoading && !apiError && (
          <Alert severity="info" sx={{ mb: 2 }}>
            No deployment components found. Click Refresh to load data.
          </Alert>
        )}

        {deploymentComponents.length > 0 && (
          <Typography variant="body2" color="text.secondary" gutterBottom sx={{ mb: 2 }}>
            Found {deploymentComponents.length} deployment component(s)
          </Typography>
        )}
        <DataTable
          columns={columns}
          data={paginatedData}
          selectable
          selectedRows={selectedRows}
          onSelectAll={handleSelectAll}
          onSelectRow={handleSelectRow}
          actions={actions}
          pagination
          page={page}
          rowsPerPage={rowsPerPage}
          totalRows={deploymentComponents.length}
          onPageChange={handlePageChange}
          onRowsPerPageChange={handleRowsPerPageChange}
          onSort={handleSort}
          defaultSortColumn="name"
          defaultSortOrder="asc"
        />
      </TableSection>
    </DashboardContainer>
  );
}
