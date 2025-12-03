'use client';
import React, { FC, ElementType, useMemo } from 'react';
import {
  TablePaginationProps as MuiTablePaginationProps,
  TablePagination as MuiTablePagination,
  Box,
} from '@mui/material';
import { FirstPage, KeyboardArrowLeft, KeyboardArrowRight, LastPage } from '@mui/icons-material';
import { TablePaginationActionsProps } from '@mui/material/TablePagination/TablePaginationActions';
import usePagination from '@mui/material/usePagination';
import { PAGINATION_MODE } from '@/models/table';
import { capitalizeWords } from '@/lib';
import { StyledNextPrevBtn, StyledPageNumBtn, StyledPaginationContainer } from './styled';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';

type PaginationActionsProps = Omit<TablePaginationActionsProps, 'onChangePage' | 'count'> & {
  mode: PAGINATION_MODE;
  totalPages?: number;
  count: number;
  setPageData: (page: number, rowsPerPage: number) => void;
};

export const PaginationActions: FC<PaginationActionsProps> = (props) => {
  const { page, rowsPerPage, totalPages, setPageData } = props;
  const hookPage = useMemo(() => page + 1, [page]);
  const { items } = usePagination({
    count: totalPages,
    defaultPage: 1,
    page: hookPage,
    onChange: (_, changedPage) => {
      const tablePage = changedPage - 1 >= 0 ? changedPage - 1 : 0;
      if (changedPage !== hookPage) {
        setPageData(tablePage, rowsPerPage);
      }
    },
  });

  return (
    <FlexCenterRow gap={0.5}>
      {items.map(({ page: itemPage, type, selected, disabled, ...item }, index) => {
        let children = null;
        if (type === 'start-ellipsis' || type === 'end-ellipsis') {
          children = '…';
        } else if (type === 'page') {
          children = (
            <StyledPageNumBtn
              variant={selected ? 'contained' : 'text'}
              color={selected ? 'primary' : 'inherit'}
              disabled={disabled}
              {...item}
            >
              {itemPage}
            </StyledPageNumBtn>
          );
        } else {
          children = (
            <StyledNextPrevBtn
              color={disabled ? 'secondary' : 'inherit'}
              disabled={disabled}
              startIcon={
                type === 'previous' ? (
                  <KeyboardArrowLeft />
                ) : type === 'next' ? (
                  <KeyboardArrowRight />
                ) : type === 'first' ? (
                  <FirstPage />
                ) : (
                  <LastPage />
                )
              }
              {...item}
            >
              {capitalizeWords(type)}
            </StyledNextPrevBtn>
          );
        }
        return <Box key={index}>{children}</Box>;
      })}
    </FlexCenterRow>
  );
};

type TablePaginationProps = Omit<MuiTablePaginationProps, 'onPageChange' | 'count'> & {
  mode: PAGINATION_MODE;
  totalPages?: number;
  totalRecords?: number;
  onPageChange: (page: number, rowsPerPage: number) => void;
  component?: ElementType<any>;
  border?: boolean;
  borderColor?: string;
  bgColor?: string;
};

export const TablePagination = ({
  page,
  totalRecords = -1,
  rowsPerPage = 10,
  mode,
  totalPages = 0,
  onPageChange,
  component,
  border,
  borderColor,
  bgColor,
  ...otherProps
}: TablePaginationProps) => {
  return (
    <StyledPaginationContainer $border={border} $borderColor={borderColor} $bgColor={bgColor}>
      <MuiTablePagination
        component={component ?? 'td'}
        page={page}
        count={totalRecords}
        rowsPerPage={rowsPerPage}
        onPageChange={null}
        onRowsPerPageChange={null}
        rowsPerPageOptions={[]}
        labelRowsPerPage=""
        labelDisplayedRows={() => ''}
        ActionsComponent={(subProps) => (
          <PaginationActions
            {...subProps}
            mode={mode}
            count={totalRecords}
            page={page}
            rowsPerPage={rowsPerPage}
            totalPages={totalPages}
            setPageData={onPageChange}
          />
        )}
        {...otherProps}
      />
    </StyledPaginationContainer>
  );
};

