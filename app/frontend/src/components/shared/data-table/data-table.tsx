'use client';
import React, { useMemo, useState } from 'react';
import {
  Checkbox,
  CircularProgress,
  IconButton,
  Menu,
  MenuItem,
  TableBody,
  TablePagination,
  Typography,
} from '@mui/material';
import { MoreVert, Visibility } from '@mui/icons-material';
import {
  StyledTableContainer,
  StyledTable,
  StyledTableHead,
  StyledTableRow,
  StyledTableCell,
  StyledTableSortLabel,
  StyledCheckboxTableCell,
  EmptyStateContainer,
  LoadingContainer,
  ActionCell,
  MenuItemIconContainer,
  StyledTablePagination,
  ActionsTableCell,
} from '@/components/shared/data-table/styled';

export type Order = 'asc' | 'desc';

export interface Column<T> {
  id: keyof T | string;
  label: string;
  minWidth?: number;
  align?: 'right' | 'left' | 'center';
  sortable?: boolean;
  render?: (value: any, row: T) => React.ReactNode;
}

export interface Action<T> {
  label: string;
  icon?: React.ReactNode;
  onClick: (row: T) => void;
  color?: 'primary' | 'secondary' | 'error' | 'warning' | 'info' | 'success';
}

export interface DataTableProps<T> {
  columns: Column<T>[];
  data: T[];
  loading?: boolean;
  emptyMessage?: string;
  selectable?: boolean;
  onSelectAll?: (selected: boolean) => void;
  onSelectRow?: (row: T, selected: boolean) => void;
  selectedRows?: T[];
  actions?: Action<T>[];
  pagination?: boolean;
  page?: number;
  rowsPerPage?: number;
  totalRows?: number;
  onPageChange?: (page: number) => void;
  onRowsPerPageChange?: (rowsPerPage: number) => void;
  rowsPerPageOptions?: number[];
  onSort?: (columnId: keyof T | string, order: Order) => void;
  defaultSortColumn?: keyof T | string;
  defaultSortOrder?: Order;
}

export function DataTable<T extends { id?: string | number }>({
  columns,
  data,
  loading = false,
  emptyMessage = 'No data available',
  selectable = false,
  onSelectAll,
  onSelectRow,
  selectedRows = [],
  actions = [],
  pagination = true,
  page = 0,
  rowsPerPage = 10,
  totalRows,
  onPageChange,
  onRowsPerPageChange,
  rowsPerPageOptions = [5, 10, 25, 50],
  onSort,
  defaultSortColumn,
  defaultSortOrder = 'asc',
}: DataTableProps<T>) {
  const [order, setOrder] = useState<Order>(defaultSortOrder || 'asc');
  const [orderBy, setOrderBy] = useState<keyof T | string>(defaultSortColumn || '');
  const [anchorEl, setAnchorEl] = useState<{ [key: string]: HTMLElement | null }>({});
  const [selectedRow, setSelectedRow] = useState<T | null>(null);

  const handleRequestSort = (property: keyof T | string) => {
    if (!onSort) return;

    const isAsc = orderBy === property && order === 'asc';
    const newOrder = isAsc ? 'desc' : 'asc';
    setOrder(newOrder);
    setOrderBy(property);
    onSort(property, newOrder);
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (onSelectAll) {
      onSelectAll(event.target.checked);
    }
  };

  const handleSelectRow = (row: T, checked: boolean) => {
    if (onSelectRow) {
      onSelectRow(row, checked);
    }
  };

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>, row: T) => {
    const rowId = String(row.id || Math.random());
    setAnchorEl({ ...anchorEl, [rowId]: event.currentTarget });
    setSelectedRow(row);
  };

  const handleMenuClose = (rowId: string) => {
    setAnchorEl({ ...anchorEl, [rowId]: null });
  };

  const handleActionClick = (action: Action<T>, row: T) => {
    action.onClick(row);
    const rowId = String(row.id || Math.random());
    handleMenuClose(rowId);
  };

  const isSelected = (row: T) => {
    return selectedRows.some((selected) => selected.id === row.id);
  };

  const isAllSelected = useMemo(() => {
    if (!selectable || data.length === 0) return false;
    return data.every((row) => isSelected(row));
  }, [data, selectedRows, selectable]);

  const isIndeterminate = useMemo(() => {
    if (!selectable || data.length === 0) return false;
    const selectedCount = data.filter((row) => isSelected(row)).length;
    return selectedCount > 0 && selectedCount < data.length;
  }, [data, selectedRows, selectable]);

  const handleChangePage = (_event: unknown, newPage: number) => {
    if (onPageChange) {
      onPageChange(newPage);
    }
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (onRowsPerPageChange) {
      onRowsPerPageChange(parseInt(event.target.value, 10));
    }
  };

  const getValue = (row: T, columnId: keyof T | string): any => {
    if (typeof columnId === 'string' && columnId.includes('.')) {
      return columnId.split('.').reduce((obj: any, key) => obj?.[key], row);
    }
    return row[columnId as keyof T];
  };

  if (loading) {
    return (
      <StyledTableContainer>
        <LoadingContainer>
          <CircularProgress />
        </LoadingContainer>
      </StyledTableContainer>
    );
  }

  if (data.length === 0) {
    return (
      <StyledTableContainer>
        <EmptyStateContainer>
          <Typography variant="body2" color="text.secondary">
            {emptyMessage}
          </Typography>
        </EmptyStateContainer>
      </StyledTableContainer>
    );
  }

  return (
    <StyledTableContainer>
      <StyledTable>
        <StyledTableHead>
          <StyledTableRow>
            {selectable && (
              <StyledCheckboxTableCell>
                <Checkbox
                  indeterminate={isIndeterminate}
                  checked={isAllSelected}
                  onChange={handleSelectAllClick}
                  color="primary"
                />
              </StyledCheckboxTableCell>
            )}
            {columns.map((column) => (
              <StyledTableCell
                key={String(column.id)}
                align={column.align || 'left'}
                style={{ minWidth: column.minWidth }}
              >
                {column.sortable !== false && onSort ? (
                  <StyledTableSortLabel
                    active={orderBy === column.id}
                    direction={orderBy === column.id ? order : 'asc'}
                    onClick={() => handleRequestSort(column.id)}
                  >
                    {column.label}
                  </StyledTableSortLabel>
                ) : (
                  column.label
                )}
              </StyledTableCell>
            ))}
            {actions.length > 0 && (
              <ActionsTableCell align="right">
                Actions
              </ActionsTableCell>
            )}
          </StyledTableRow>
        </StyledTableHead>
        <TableBody>
          {data.map((row, index) => {
            const rowId = String(row.id || index);
            const selected = isSelected(row);
            const menuOpen = Boolean(anchorEl[rowId]);

            return (
              <StyledTableRow key={rowId} selected={selected} hover>
                {selectable && (
                  <StyledCheckboxTableCell>
                    <Checkbox
                      checked={selected}
                      onChange={(e) => handleSelectRow(row, e.target.checked)}
                      color="primary"
                    />
                  </StyledCheckboxTableCell>
                )}
                {columns.map((column) => {
                  const value = getValue(row, column.id);
                  return (
                    <StyledTableCell key={String(column.id)} align={column.align || 'left'}>
                      {column.render ? column.render(value, row) : value ?? '-'}
                    </StyledTableCell>
                  );
                })}
                {actions.length > 0 && (
                  <StyledTableCell align="right">
                    <ActionCell>
                      {actions.length === 1 ? (
                        <IconButton
                          size="small"
                          onClick={() => actions[0].onClick(row)}
                          color={actions[0].color || 'primary'}
                        >
                          {actions[0].icon || <Visibility />}
                        </IconButton>
                      ) : (
                        <>
                          <IconButton size="small" onClick={(e) => handleMenuOpen(e, row)}>
                            <MoreVert />
                          </IconButton>
                          <Menu
                            anchorEl={anchorEl[rowId]}
                            open={menuOpen}
                            onClose={() => handleMenuClose(rowId)}
                            anchorOrigin={{
                              vertical: 'bottom',
                              horizontal: 'right',
                            }}
                            transformOrigin={{
                              vertical: 'top',
                              horizontal: 'right',
                            }}
                          >
                            {actions.map((action, actionIndex) => (
                              <MenuItem
                                key={actionIndex}
                                onClick={() => handleActionClick(action, row)}
                              >
                                {action.icon && <MenuItemIconContainer>{action.icon}</MenuItemIconContainer>}
                                {action.label}
                              </MenuItem>
                            ))}
                          </Menu>
                        </>
                      )}
                    </ActionCell>
                  </StyledTableCell>
                )}
              </StyledTableRow>
            );
          })}
        </TableBody>
      </StyledTable>
      {pagination && (
        <StyledTablePagination>
          <TablePagination
            component="div"
            count={totalRows || data.length}
            page={page}
            onPageChange={handleChangePage}
            rowsPerPage={rowsPerPage}
            onRowsPerPageChange={handleChangeRowsPerPage}
            rowsPerPageOptions={rowsPerPageOptions}
          />
        </StyledTablePagination>
      )}
    </StyledTableContainer>
  );
}

