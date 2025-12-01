'use client';
import { useState, useMemo, useCallback, useEffect } from 'react';
import { Typography, Button, Box, Stack, TextField } from '@mui/material';
import { create } from '@bufbuild/protobuf';
import { CloudResourceContainer, TableSection } from '@/app/cloud-resources/styled';
import { DataTable, Column, Action } from '@/components/shared/data-table';
import { Drawer } from '@/components/shared/drawer';
import { YamlEditor } from '@/components/shared/yaml-editor';
import { Edit, Delete, Visibility, Refresh, Add } from '@mui/icons-material';
import { useCloudResourceQuery, useCloudResourceCommand } from '@/app/cloud-resources/_services';
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
  const [apiLoading, setApiLoading] = useState(false);
  const [kindFilter, setKindFilter] = useState<string>('');

  // Drawer state
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [drawerMode, setDrawerMode] = useState<DrawerMode>(null);
  const [selectedResource, setSelectedResource] = useState<CloudResource | null>(null);
  const [formData, setFormData] = useState<{ manifest: string }>({ manifest: '' });

  // Function to call the API
  const handleLoadCloudResources = useCallback(() => {
    if (!query) {
      return;
    }

    setApiLoading(true);

    const request = create(ListCloudResourcesRequestSchema, {
      kind: kindFilter.trim() || undefined,
    });

    query
      .listCloudResources(request)
      .then((result) => {
        setCloudResources(result.resources);
        setApiLoading(false);
      })
      .catch((err: any) => {
        console.error('API Error:', err);
        setApiLoading(false);
      });
  }, [query, kindFilter]);

  // Auto-load on mount and when kind filter changes
  useEffect(() => {
    if (query) {
      handleLoadCloudResources();
    }
  }, [query, handleLoadCloudResources]);

  // Handle kind filter change
  const handleKindFilterChange = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    const newKind = event.target.value;
    setKindFilter(newKind);
    setPage(0); // Reset to first page when filter changes
    setSelectedRows([]); // Clear selected rows when filter changes
  }, []);

  const columns: Column<CloudResource>[] = useMemo(
    () => [
      { id: 'name', label: 'Name', sortable: true },
      { id: 'kind', label: 'Kind', sortable: true },
      {
        id: 'createdAt',
        label: 'Created At',
        sortable: true,
        render: (value: Timestamp | undefined) =>
          value ? formatTimestampToDate(value, 'DD/MM/YYYY') : '-',
      },
      {
        id: 'updatedAt',
        label: 'Updated At',
        sortable: true,
        render: (value: Timestamp | undefined) =>
          value ? formatTimestampToDate(value, 'DD/MM/YYYY') : '-',
      },
    ],
    []
  );

  // Drawer handlers
  const handleOpenDrawer = useCallback((mode: DrawerMode, resource?: CloudResource) => {
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
  }, []);

  const handleCloseDrawer = useCallback(() => {
    setDrawerOpen(false);
    setDrawerMode(null);
    setSelectedResource(null);
    setFormData({ manifest: '' });
  }, []);

  const handleSave = useCallback(() => {
    if (!command) {
      return;
    }

    if (drawerMode === 'create') {
      command.create(formData.manifest).then(() => {
        handleLoadCloudResources();
        handleCloseDrawer();
      });
    } else if (drawerMode === 'edit' && selectedResource) {
      command.update(selectedResource.id, formData.manifest).then(() => {
        handleLoadCloudResources();
        handleCloseDrawer();
      });
    }
  }, [
    command,
    drawerMode,
    formData.manifest,
    selectedResource,
    handleLoadCloudResources,
    handleCloseDrawer,
  ]);

  const handleDelete = useCallback(
    (row: CloudResource) => {
      if (!command) {
        return;
      }
      if (window.confirm(`Are you sure you want to delete "${row.name}"?`)) {
        command
          .delete(row.id)
          .then(() => {
            // Refresh the list after deletion
            handleLoadCloudResources();
          })
          .catch((err: any) => {
            // Error is already handled by command service with snackbar
            console.error('Delete error:', err);
          });
      }
    },
    [command, handleLoadCloudResources]
  );

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
        onClick: handleDelete,
        color: 'error',
      },
    ],
    [handleOpenDrawer, handleDelete]
  );

  const handleSelectAll = useCallback(
    (selected: boolean) => {
      if (selected) {
        setSelectedRows(cloudResources);
      } else {
        setSelectedRows([]);
      }
    },
    [cloudResources]
  );

  const handleSelectRow = useCallback((row: CloudResource, selected: boolean) => {
    setSelectedRows((prevSelected) => {
      if (selected) {
        return [...prevSelected, row];
      } else {
        return prevSelected.filter((r) => r.id !== row.id);
      }
    });
  }, []);

  const handleSort = useCallback(
    (columnId: keyof CloudResource | string, order: 'asc' | 'desc') => {
      setSortColumn(columnId);
      setSortOrder(order);
    },
    []
  );

  const handlePageChange = useCallback((newPage: number) => {
    setPage(newPage);
  }, []);

  const handleRowsPerPageChange = useCallback((newRowsPerPage: number) => {
    setRowsPerPage(newRowsPerPage);
    setPage(0);
  }, []);

  // Function to reload data from API
  const handleRefresh = useCallback(() => {
    handleLoadCloudResources();
  }, [handleLoadCloudResources]);

  // Handle YAML editor change
  const handleYamlChange = useCallback(
    (value: string) => {
      if (drawerMode !== 'view') {
        setFormData({ manifest: value });
      }
    },
    [drawerMode]
  );

  // Handle create button click
  const handleCreateClick = useCallback(() => {
    handleOpenDrawer('create');
  }, [handleOpenDrawer]);

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
          <Stack flexDirection="row" gap={1}>
            <TextField
              size="small"
              value={kindFilter}
              onChange={handleKindFilterChange}
              placeholder="e.g., CivoVpc, AwsRdsInstance"
              sx={{ minWidth: 250 }}
            />
            <Button variant="contained" startIcon={<Add />} onClick={handleCreateClick}>
              Create
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              startIcon={<Refresh />}
              onClick={handleRefresh}
            >
              Refresh
            </Button>
          </Stack>
        </Box>

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
              <Typography variant="subtitle2">Manifest (YAML)</Typography>
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
              value={drawerMode === 'view' ? selectedResource?.manifest || '' : formData.manifest}
              onChange={handleYamlChange}
              readOnly={drawerMode === 'view'}
              height="500px"
            />
          </Box>
        </Stack>
      </Drawer>
    </CloudResourceContainer>
  );
}
