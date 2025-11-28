'use client';
import { useState, useMemo, useCallback, useEffect } from 'react';
import { Typography, Button, Box, Stack, Alert } from '@mui/material';
import { create } from '@bufbuild/protobuf';
import { CloudResourceContainer, TableSection } from '@/app/dashboard/styled';
import { DataTable, Column, Action } from '@/components/shared/data-table';
import { Drawer } from '@/components/shared/drawer';
import { YamlEditor } from '@/components/shared/yaml-editor';
import { Edit, Delete, Visibility, Refresh, Add } from '@mui/icons-material';
import { useCloudResourceQuery, useCloudResourceCommand } from '@/app/dashboard/_services';
import {
  ListCloudResourcesRequestSchema,
  CloudResource,
} from '@/gen/proto/cloud_resource_service_pb';
import { Timestamp } from '@bufbuild/protobuf/wkt';
import { formatTimestampToDate } from '@/lib';

type DrawerMode = 'view' | 'edit' | 'create' | null;

export default function CloudResourcesPage() {
  const { query } = useCloudResourceQuery();
  const { command } = useCloudResourceCommand();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [selectedRows, setSelectedRows] = useState<CloudResource[]>([]);
  const [sortColumn, setSortColumn] = useState<keyof CloudResource | string>('name');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc');
  const [cloudResources, setCloudResources] = useState<CloudResource[]>([]);
  const [apiError, setApiError] = useState<string | null>(null);
  const [apiLoading, setApiLoading] = useState(false);

  // Drawer state
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [drawerMode, setDrawerMode] = useState<DrawerMode>(null);
  const [selectedResource, setSelectedResource] = useState<CloudResource | null>(null);
  const [formData, setFormData] = useState<{ manifest: string }>({ manifest: '' });


  // Function to call the API
  const handleLoadCloudResources = useCallback(() => {
    if (!query) {
      setApiError('Query service not ready');
      return;
    }

    setApiLoading(true);
    setApiError(null);
    query
      .listCloudResources(create(ListCloudResourcesRequestSchema))
      .then((result) => {
        setCloudResources(result.resources);
        setApiLoading(false);
      })
      .catch((err: any) => {
        setApiError(err.message || 'Failed to load cloud resources');
        console.error('API Error:', err);
        setApiLoading(false);
      });
  }, [query]);

  // Auto-load on mount
  useEffect(() => {
    if (query) {
      handleLoadCloudResources();
    }
  }, [query, handleLoadCloudResources]);

  const columns: Column<CloudResource>[] = useMemo(
    () => [
      { id: 'name', label: 'Name', sortable: true },
      { id: 'kind', label: 'Kind', sortable: true },
      {
        id: 'createdAt',
        label: 'Created At',
        sortable: true,
        render: (value: Timestamp | undefined) => value ? formatTimestampToDate(value, 'DD/MM/YYYY') : '-',
      },
      {
        id: 'updatedAt',
        label: 'Updated At',
        sortable: true,
        render: (value: Timestamp | undefined) => value ? formatTimestampToDate(value, 'DD/MM/YYYY') : '-',
      },
    ],
    []
  );

  // Drawer handlers
  const handleOpenDrawer = (mode: DrawerMode, resource?: CloudResource) => {
    setDrawerMode(mode);
    if (resource) {
      setSelectedResource(resource);
      setFormData({
        manifest: resource.manifest || '',
      });
    } else {
      setSelectedResource(null);
      setFormData({ manifest: '' });
    }
    setDrawerOpen(true);
  };

  const handleCloseDrawer = () => {
    setDrawerOpen(false);
    setDrawerMode(null);
    setSelectedResource(null);
    setFormData({ manifest: '' });
  };

  const handleSave = async () => {
    if (!command) {
      setApiError('Command service not ready');
      return;
    }

    if (!formData.manifest || formData.manifest.trim() === '') {
      setApiError('Manifest is required');
      return;
    }

    try {
      if (drawerMode === 'create') {
        await command.create(formData.manifest);
        // Refresh the list after creation
        handleLoadCloudResources();
        handleCloseDrawer();
      } else if (drawerMode === 'edit' && selectedResource) {
        await command.update(selectedResource.id, formData.manifest);
        // Refresh the list after update
        handleLoadCloudResources();
        handleCloseDrawer();
      }
    } catch (err: any) {
      // Error is already handled by command service with snackbar
      console.error('Save error:', err);
    }
  };

  const actions: Action<CloudResource>[] = useMemo(
    () => [
      {
        label: 'View',
        icon: <Visibility fontSize="small" />,
        onClick: (row) => {
          handleOpenDrawer('view', row);
        },
        color: 'primary',
      },
      {
        label: 'Edit',
        icon: <Edit fontSize="small" />,
        onClick: (row) => {
          handleOpenDrawer('edit', row);
        },
        color: 'primary',
      },
      {
        label: 'Delete',
        icon: <Delete fontSize="small" />,
        onClick: async (row) => {
          if (!command) {
            setApiError('Command service not ready');
            return;
          }
          if (window.confirm(`Are you sure you want to delete "${row.name}"?`)) {
            try {
              await command.delete(row.id);
              // Refresh the list after deletion
              handleLoadCloudResources();
            } catch (err: any) {
              // Error is already handled by command service with snackbar
              console.error('Delete error:', err);
            }
          }
        },
        color: 'error',
      },
    ],
    []
  );

  const handleSelectAll = (selected: boolean) => {
    if (selected) {
      setSelectedRows(cloudResources);
    } else {
      setSelectedRows([]);
    }
  };

  const handleSelectRow = (row: CloudResource, selected: boolean) => {
    if (selected) {
      setSelectedRows([...selectedRows, row]);
    } else {
      setSelectedRows(selectedRows.filter((r) => r.id !== row.id));
    }
  };

  const handleSort = (columnId: keyof CloudResource | string, order: 'asc' | 'desc') => {
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

  // Function to reload data from API
  const handleRefresh = () => {
    handleLoadCloudResources();
  };

  // Apply sorting
  const sortedData = useMemo(() => {
    const sorted = [...cloudResources];
    sorted.sort((a, b) => {
      const aValue = a[sortColumn as keyof CloudResource];
      const bValue = b[sortColumn as keyof CloudResource];
      if (sortOrder === 'asc') {
        return aValue > bValue ? 1 : -1;
      }
      return aValue < bValue ? 1 : -1;
    });
    return sorted;
  }, [cloudResources, sortColumn, sortOrder]);

  // Apply pagination
  const paginatedData = useMemo(() => {
    const start = page * rowsPerPage;
    return sortedData.slice(start, start + rowsPerPage);
  }, [sortedData, page, rowsPerPage]);

  return (
    <CloudResourceContainer>
      <Typography variant="h4" gutterBottom>
        Cloud Resources
      </Typography>

      <TableSection>
        <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
          <Typography variant="h5">Cloud Resources</Typography>
          <Box display="flex" gap={2}>
            <Button
              variant="contained"
              startIcon={<Add />}
              onClick={() => handleOpenDrawer('create')}
            >
              Create
            </Button>
            <Button
              variant="outlined"
              startIcon={<Refresh />}
              onClick={handleRefresh}
            >
              Refresh
            </Button>
          </Box>
        </Box>

        {apiError && (
          <Alert severity="error" sx={{ mb: 2 }} onClose={() => setApiError(null)}>
            {apiError}
          </Alert>
        )}
        {!apiLoading && cloudResources.length > 0 && (
          <Typography variant="body2" color="text.secondary" gutterBottom sx={{ mb: 2 }}>
            Found {cloudResources.length} cloud resource(s)
          </Typography>
        )}
        {apiLoading && (
          <Typography variant="body2" color="text.secondary" gutterBottom sx={{ mb: 2 }}>
            Loading cloud resources...
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
          totalRows={cloudResources.length}
          onPageChange={handlePageChange}
          onRowsPerPageChange={handleRowsPerPageChange}
          onSort={handleSort}
          defaultSortColumn="name"
          defaultSortOrder="asc"
        />
      </TableSection>

      {/* Drawer for View/Edit/Create */}
      <Drawer
        open={drawerOpen}
        onClose={handleCloseDrawer}
        title={
          drawerMode === 'view'
            ? 'View Cloud Resource'
            : drawerMode === 'edit'
            ? 'Edit Cloud Resource'
            : 'Create Cloud Resource'
        }
        width={800}
      >
        <Stack spacing={3}>
          <Box>
            <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
              <Typography variant="subtitle2">
                Manifest (YAML)
              </Typography>
              {drawerMode !== 'view' && (
                <Box display="flex" gap={1}>
                  <Button variant="contained" color="secondary" onClick={handleCloseDrawer}>
                    Cancel
                  </Button>
                  <Button variant="contained" color="primary" onClick={handleSave}>
                    {drawerMode === 'edit' ? 'Update' : 'Create'}
                  </Button>
                </Box>
              )}
            </Box>
            <YamlEditor
              value={drawerMode === 'view' ? (selectedResource?.manifest || '') : formData.manifest}
              onChange={(value) => {
                if (drawerMode !== 'view') {
                  setFormData({ manifest: value });
                }
              }}
              readOnly={drawerMode === 'view'}
              height="500px"
            />
          </Box>
        </Stack>
      </Drawer>
    </CloudResourceContainer>
  );
}

