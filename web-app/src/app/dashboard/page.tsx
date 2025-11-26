'use client';
import { useState, useMemo, useEffect, useCallback } from 'react';
import { Typography, Grid2, Button, Box, Alert } from '@mui/material';
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

interface SampleData {
  id: number;
  name: string;
  email: string;
  role: string;
  status: string;
  createdAt: string;
}

interface DeploymentComponentData {
  id: string;
  kind: string;
  provider: string;
  name: string;
  version: string;
  idPrefix: string;
  isServiceKind: boolean;
  createdAt?: Date;
  updatedAt?: Date;
}

export default function DashboardPage() {
  const { query } = useDashboardQuery();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [selectedRows, setSelectedRows] = useState<SampleData[]>([]);
  const [sortColumn, setSortColumn] = useState<keyof SampleData | string>('name');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc');
  const [deploymentComponents, setDeploymentComponents] = useState<DeploymentComponentData[]>([]);
  const [apiError, setApiError] = useState<string | null>(null);
  const [apiLoading, setApiLoading] = useState(false);

  // Sample data
  const sampleData: SampleData[] = useMemo(
    () => [
      {
        id: 1,
        name: 'John Doe',
        email: 'john.doe@example.com',
        role: 'Admin',
        status: 'Active',
        createdAt: '2024-01-15',
      },
      {
        id: 2,
        name: 'Jane Smith',
        email: 'jane.smith@example.com',
        role: 'User',
        status: 'Active',
        createdAt: '2024-01-20',
      },
      {
        id: 3,
        name: 'Bob Johnson',
        email: 'bob.johnson@example.com',
        role: 'Editor',
        status: 'Inactive',
        createdAt: '2024-02-01',
      },
      {
        id: 4,
        name: 'Alice Williams',
        email: 'alice.williams@example.com',
        role: 'User',
        status: 'Active',
        createdAt: '2024-02-10',
      },
      {
        id: 5,
        name: 'Charlie Brown',
        email: 'charlie.brown@example.com',
        role: 'Admin',
        status: 'Active',
        createdAt: '2024-02-15',
      },
    ],
    []
  );

  const columns: Column<SampleData>[] = useMemo(
    () => [
      { id: 'name', label: 'Name', sortable: true },
      { id: 'email', label: 'Email', sortable: true },
      { id: 'role', label: 'Role', sortable: true },
      {
        id: 'status',
        label: 'Status',
        sortable: true,
        render: (value: string) => (
          <StatusBadge $status={value === 'Active' ? 'active' : 'inactive'}>{value}</StatusBadge>
        ),
      },
      { id: 'createdAt', label: 'Created At', sortable: true },
    ],
    []
  );

  const actions: Action<SampleData>[] = useMemo(
    () => [
      {
        label: 'View',
        icon: <Visibility fontSize="small" />,
        onClick: (row) => {
          console.log('View:', row);
          // Handle view action
        },
        color: 'primary',
      },
      {
        label: 'Edit',
        icon: <Edit fontSize="small" />,
        onClick: (row) => {
          console.log('Edit:', row);
          // Handle edit action
        },
        color: 'primary',
      },
      {
        label: 'Delete',
        icon: <Delete fontSize="small" />,
        onClick: (row) => {
          console.log('Delete:', row);
          // Handle delete action
        },
        color: 'error',
      },
    ],
    []
  );

  const handleSelectAll = (selected: boolean) => {
    if (selected) {
      setSelectedRows(sampleData);
    } else {
      setSelectedRows([]);
    }
  };

  const handleSelectRow = (row: SampleData, selected: boolean) => {
    if (selected) {
      setSelectedRows([...selectedRows, row]);
    } else {
      setSelectedRows(selectedRows.filter((r) => r.id !== row.id));
    }
  };

  const handleSort = (columnId: keyof SampleData | string, order: 'asc' | 'desc') => {
    setSortColumn(columnId);
    setSortOrder(order);
    // Implement sorting logic here
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
      .listDeploymentComponents({})
      .then((result) => {
        console.log('result', result);
        setDeploymentComponents(result?.components || []);
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
      console.log('query', query);
      handleLoadDeploymentComponents();
    }
  }, [query, handleLoadDeploymentComponents]);

  // Apply sorting
  const sortedData = useMemo(() => {
    const sorted = [...sampleData];
    sorted.sort((a, b) => {
      const aValue = a[sortColumn as keyof SampleData];
      const bValue = b[sortColumn as keyof SampleData];
      if (sortOrder === 'asc') {
        return aValue > bValue ? 1 : -1;
      }
      return aValue < bValue ? 1 : -1;
    });
    return sorted;
  }, [sampleData, sortColumn, sortOrder]);

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
          <Typography variant="h5">Deployment Components (API Test)</Typography>
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
            No deployment components found. The collection is empty. Click Refresh to check again.
          </Alert>
        )}

        {deploymentComponents.length > 0 && (
          <Box sx={{ mb: 3 }}>
            <Typography variant="body2" color="text.secondary" gutterBottom>
              Found {deploymentComponents.length} deployment component(s)
            </Typography>
            <Box sx={{ mt: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
              {deploymentComponents.map((component) => (
                <Box
                  key={component.id}
                  sx={{ mb: 2, pb: 2, borderBottom: '1px solid', borderColor: 'divider' }}
                >
                  <Typography variant="subtitle1" fontWeight="bold">
                    {component.name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    <strong>ID:</strong> {component.id}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    <strong>Kind:</strong> {component.kind} | <strong>Provider:</strong>{' '}
                    {component.provider}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    <strong>Version:</strong> {component.version} | <strong>ID Prefix:</strong>{' '}
                    {component.idPrefix}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    <strong>Service Kind:</strong> {component.isServiceKind ? 'Yes' : 'No'}
                  </Typography>
                  {component.createdAt && (
                    <Typography variant="body2" color="text.secondary">
                      <strong>Created:</strong> {component.createdAt.toLocaleString()}
                    </Typography>
                  )}
                </Box>
              ))}
            </Box>
          </Box>
        )}

        <Typography variant="h5" gutterBottom sx={{ mt: 4 }}>
          Users Table (Sample Data)
        </Typography>
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
          totalRows={sampleData.length}
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
