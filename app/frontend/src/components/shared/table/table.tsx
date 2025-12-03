'use client';
import Image from 'next/image';
import React, { MouseEvent, useEffect, useMemo, useState } from 'react';
import { Checkbox, Stack, Table, TableBody, Typography, Skeleton } from '@mui/material';
import { getDataFromPath } from '@/lib';
import { ActionMenuProps, DotNestedKeys, PAGINATION_MODE } from '@/models/table';
import { ActionType } from '@/models/table';
import { ConfirmationDialog } from '@/components/shared/confirmation-dialog';
import { TablePagination } from '@/components/shared/table/pagination';
import { Icon, ICON_NAMES } from '@/components/shared/icon';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';
import {
  StyledActionButton,
  StyledCheckboxTableCell,
  StyledHeaderLabel,
  StyledIconButton,
  StyledInputLabel,
  StyledLink,
  StyledMenu,
  StyledMenuIcon,
  StyledMenuItem,
  StyledTableCell,
  StyledTableContainer,
  StyledTableHead,
  StyledTableRow,
  StyledTextCopy,
} from '@/components/shared/table/styled';
import {
  AlertDialog,
  AlertDialogProps,
  defaultAlertDialogProps,
} from '@/components/shared/alert-dialog';

interface ConfirmActionProps {
  openModal: boolean;
  confirmHandler: (reason?: string) => () => void;
  actionId: string;
  actionLabel: string;
  actionMessage: string;
  actionType: ActionType;
}

const defaultConfirmAction: ConfirmActionProps = {
  openModal: false,
  confirmHandler: null,
  actionId: null,
  actionLabel: 'Confirm',
  actionMessage: null,
  actionType: ActionType.UNSPECIFIED,
};

export interface TableProps<T> {
  options: {
    /**
     * Defines table head title columns
     */
    headers?: string[];
    /**
     * Defines data path of each column
     */
    dataPath?: Array<DotNestedKeys<Partial<T>>>;
    /**
     * Defines the copyable columns
     */
    copyableColumns?: Array<DotNestedKeys<Partial<T>>>;
    /**
     * Defines the clickable columns
     */
    clickableColumns?: { [key in DotNestedKeys<Partial<T>>]?: (row: T) => void };
    /**
     * Defines the date columns
     */
    dateColumns?: string[];
    /**
     * add colspan to table fields
     */
    colSpan?: { [key in DotNestedKeys<Partial<T>>]?: number };
    /**
     * Provide template for any column value to customize it. ex. string -> link with icon
     */
    dataTemplate?: {
      [key in DotNestedKeys<Partial<T>>]?: (val: string, row: T) => React.ReactNode;
    };
    /**
     * Override value of any dataPath with custom values
     */
    override?: { [key in DotNestedKeys<Partial<T>>]?: (val: T) => React.ReactNode };
    /**
     * Defines table actions
     */
    actions?: ActionMenuProps<T>[];
    /**
     * Defines table action direction
     */
    actionDirection?: 'flex-start' | 'center' | 'flex-end';
    /**
     * Defines Delete action
     */
    deleteBtn?: ActionMenuProps<T>;
    /**
     * Defines Edit action
     */
    editBtn?: ActionMenuProps<T>;
    /**
     * Defines empty message to display when no data rows
     */
    emptyMessage?: string;
    /**
     * Show or hide table pagination
     */
    onPageChange?: (page: number, rowsPerPage: number) => void;
    showPagination?: boolean;
    rowsPerPage?: number;
    currentPage?: number;
    paginationMode?: PAGINATION_MODE;
    /**
     * Set total pages when it is server pagination mode
     */
    totalPages?: number;
    /**
     * Show checkbox for each row
     */
    isRowSelectable?: boolean;
    /**
     * Returns selected rows as an array to the caller.
     */
    selectRowHandler?: (rows: T[]) => void;
    /**
     * Sticky header and action column
     */
    stickyHeader?: boolean;
    stickyAction?: boolean;
    headerBgColor?: string;
    border?: boolean;
    paginationborder?: boolean;
    borderColor?: string;
    bgColor?: string;
  };
  /**
   * Defines table data
   */
  data?: T[];
  loading?: boolean;
  onActionMenuClick?: (T) => void;
}

/**
 * Usage
 * <TableComp
    options={{
      headers: ['Id', 'Organization Name', 'Email', 'City', 'Actions'],
      dataPath: ['address.suite', 'organization.name', 'organization.catchPhrase', 'address.city'], // intellisense automatically provides all json paths in the data object
      dataTemplate: {
          'address.suite': templateFn // Provide template for any column value to customize it. Eg. string -> link with icon
      },
      actions: [
        {
            props: {color: 'info'},
            iconUrl: '/images/view.svg', ---|
            icon: <Icon />							 ---|----- Anyone of these 2
            text: "btn text"
            handler: handleClick, // Handler fn returns the whole row object which is clicked
        },
        {
            props: {color: 'info'},
            text: 'Edit',
            handler: handleClick,
        },
        {
            props: {color: 'info'},
            text: 'Delete',
            handler: handleClick,
        }
      ],
      onPageChange: onPageChange
    }}
    data={dummy}
  />
 */

export function TableComp<T>({
  options: {
    headers = [],
    actions = [],
    actionDirection = 'center',
    deleteBtn,
    editBtn,
    emptyMessage,
    dataPath,
    copyableColumns = [],
    clickableColumns = {},
    dateColumns = [
      'status.sysAudit.createdAt.seconds',
      'status.sysAudit.updatedAt.seconds',
      'sysAudit.createdAt.seconds',
      'sysAudit.updatedAt.seconds',
      'specAudit.createdAt.seconds',
      'specAudit.updatedAt.seconds',
      'status.audit.specAudit.createdAt.seconds',
    ],
    dataTemplate,
    override,
    colSpan,
    onPageChange,
    showPagination = true,
    paginationMode = PAGINATION_MODE.SERVER,
    totalPages = 0,
    currentPage = 0,
    rowsPerPage = 25,
    isRowSelectable = false,
    selectRowHandler,
    stickyHeader = false,
    stickyAction = true,
    headerBgColor,
    border = false,
    paginationborder = false,
    borderColor,
    bgColor,
  },
  data = [],
  loading,
  onActionMenuClick,
}: TableProps<T>) {
  const [proceedConfirmationModalOpen, setProceedConfirmationModalOpen] = useState(false);
  const [confirmationModalHandler, setConfirmationModalHandler] = useState(null);
  const [selectedRows, setSelectedRows] = useState<readonly number[]>([]);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const [actionMenu, setActionMenu] = useState<{ item: T }>();
  const [confirmAction, setConfirmAction] = useState(defaultConfirmAction);
  const [alertDialogProps, setAlertDialogProps] =
    useState<AlertDialogProps>(defaultAlertDialogProps);

  useEffect(() => {
    if (isRowSelectable) {
      setSelectedRows([]);
    }
  }, [data]);

  useEffect(() => {
    if (data && data.length && isRowSelectable) {
      const checkedRows = data.filter((item, index) => selectedRows.includes(index));
      selectRowHandler(checkedRows);
    }
  }, [selectedRows]);

  const handleOnMenuItemClick = (action: ActionMenuProps<T>, item) => {
    if (action?.type) {
      setConfirmAction({
        openModal: true,
        confirmHandler: action?.handler
          ? (reason: string) => action?.handler.bind(null, item, reason)
          : null,
        actionId: !action?.skipDeleteIdCheck ? item?.id : null,
        actionLabel: action.text,
        actionType: action.type,
        actionMessage:
          action?.deleteMessageKey && action?.message
            ? action?.message?.replace(
                '${key}',
                action.deleteMessageKey
                  ? item[action.deleteMessageKey]
                  : item.id || item.metadata?.id
              )
            : action?.message
              ? action.message
              : null,
      });
    } else if (action?.requiresProceedConfirmation) {
      setConfirmationModalHandler(() => () => action?.handler.bind(null, item));
      setProceedConfirmationModalOpen(true);
      setAlertDialogProps((prev) => ({
        ...prev,
        submitLabel: action.text,
        subTitle: action.label,
      }));
    } else {
      action?.handler(item);
    }
    handleActionMenuClose();
  };

  const handleActionConfirmation = (reason: string) => {
    if (confirmAction.confirmHandler) {
      const handler = confirmAction.confirmHandler;
      handler(reason)();
      setConfirmAction(defaultConfirmAction);
    }
  };

  const handleProceedConfirmation = () => {
    if (confirmationModalHandler) {
      confirmationModalHandler()();
      setProceedConfirmationModalOpen(false);
    }
  };

  const onActionConfirmationClose = () => {
    setConfirmAction(defaultConfirmAction);
  };

  const onProceedConfirmationClose = () => {
    setProceedConfirmationModalOpen(false);
  };

  const handleAllRowSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelected = data.map((_n, index) => index);
      setSelectedRows(newSelected);
      return;
    }
    setSelectedRows([]);
  };

  const handleRowSelect = (
    event: React.ChangeEvent<HTMLInputElement>,
    _checked: boolean,
    index: number
  ) => {
    const selectedIndex = selectedRows.indexOf(index);
    let newSelected: readonly number[] = [];
    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selectedRows, index);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selectedRows.slice(1));
    } else if (selectedIndex === selectedRows.length - 1) {
      newSelected = newSelected.concat(selectedRows.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selectedRows.slice(0, selectedIndex),
        selectedRows.slice(selectedIndex + 1)
      );
    }

    setSelectedRows(newSelected);
  };

  const isSelected = (index: number) => selectedRows.indexOf(index) !== -1;

  const handleActionMenuClick = (event: React.MouseEvent<HTMLElement>, item: T) => {
    if (onActionMenuClick) onActionMenuClick(item);
    setActionMenu({ item });
    setAnchorEl(event.currentTarget);
  };

  const handleActionMenuClose = () => {
    setAnchorEl(null);
  };

  const confirmReasonPlaceholder = useMemo(() => {
    return confirmAction.actionType === ActionType.REFRESH
      ? 'Need to refresh the state to reapply computed fields'
      : 'No Longer Needed';
  }, [confirmAction]);

  return (
    <React.Fragment>
      <StyledTableContainer $stickyHeader={stickyHeader} $border={border} $bgColor={bgColor}>
        <Table stickyHeader={stickyHeader}>
          {!!headers.length && (
            <StyledTableHead
              $headerBgColor={headerBgColor}
              $border={border}
              $borderColor={borderColor}
            >
              <StyledTableRow>
                {isRowSelectable && data.length > 0 && (
                  <StyledCheckboxTableCell>
                    <Checkbox
                      checked={data.length > 0 && selectedRows.length === data.length}
                      onChange={handleAllRowSelect}
                    />
                  </StyledCheckboxTableCell>
                )}
                {headers?.map((header, index) => (
                  <StyledTableCell
                    key={`${header}-${index}`}
                    $isFixed={stickyAction && index === headers.length - 1}
                    $border={border}
                  >
                    <Stack
                      flexDirection="row"
                      justifyContent={
                        headers.length > dataPath.length && index === headers.length - 1
                          ? actionDirection
                          : 'flex-start'
                      }
                    >
                      <StyledHeaderLabel>{header}</StyledHeaderLabel>
                    </Stack>
                  </StyledTableCell>
                ))}
              </StyledTableRow>
            </StyledTableHead>
          )}
          <TableBody>
            {!!loading && (
              <>
                {Array.from({ length: 5 }).map((_, rowIndex) => (
                  <StyledTableRow
                    key={`skeleton-${rowIndex}`}
                    $border={border}
                    $borderColor={borderColor}
                  >
                    {isRowSelectable && (
                      <StyledCheckboxTableCell>
                        <Skeleton variant="rectangular" width={20} height={20} />
                      </StyledCheckboxTableCell>
                    )}
                    {headers?.map((path, colIndex) => (
                      <StyledTableCell
                        key={`skeleton-${rowIndex}-${colIndex}`}
                        colSpan={colSpan ? colSpan[path] : 0}
                        $border={border}
                      >
                        <Skeleton variant="text" width="100%" height={20} />
                      </StyledTableCell>
                    ))}
                  </StyledTableRow>
                ))}
              </>
            )}
            {!loading &&
              data?.length > 0 &&
              data?.map((item, di) => (
                <StyledTableRow key={`${item}-${di}`} $border={border} $borderColor={borderColor}>
                  {isRowSelectable && (
                    <StyledCheckboxTableCell>
                      <Checkbox
                        checked={isSelected(di)}
                        onChange={(event, checked) => handleRowSelect(event, checked, di)}
                      />
                    </StyledCheckboxTableCell>
                  )}
                  {dataPath?.map((path, ei) => {
                    const cellData = getDataFromPath(path, item);
                    const template =
                      override && override[path]
                        ? override[path](item)
                        : dataTemplate && dataTemplate[path]
                          ? dataTemplate[path](cellData, item)
                          : dateColumns.includes(path)
                            ? new Date(+cellData * 1000).toLocaleString()
                            : cellData;
                    return (
                      <StyledTableCell
                        key={`${path}-${ei}`}
                        colSpan={colSpan ? colSpan[path] : 0}
                        $border={border}
                      >
                        {clickableColumns[path] ? (
                          <StyledLink onClick={() => clickableColumns[path](item)} color="primary">
                            {template}
                          </StyledLink>
                        ) : (
                          <StyledInputLabel>{template}</StyledInputLabel>
                        )}
                        {copyableColumns.includes(path) && (
                          <StyledTextCopy text={getDataFromPath(path, item)} />
                        )}
                      </StyledTableCell>
                    );
                  })}
                  {actions
                    ?.map((action) => {
                      // For backward compatibility, keeping isMenuAction true by default
                      if (action.isMenuAction === undefined) {
                        action.isMenuAction = true;
                      }
                      return action;
                    })
                    .filter(
                      (action) =>
                        (action?.filterAction && action?.filterAction?.(item)) ||
                        !action?.filterAction
                    )?.length > 0 && (
                    <StyledTableCell
                      $isFixed={stickyAction}
                      $border={border}
                      $verticalAlign="middle"
                    >
                      <Stack flexDirection="row" justifyContent={actionDirection}>
                        {actions
                          .filter((action) => !action.isMenuAction)
                          .map((action, ai) => (
                            <StyledActionButton
                              key={`${action.text}-${ai}`}
                              {...action.btnProps}
                              onClick={() => {
                                handleOnMenuItemClick(action, item);
                              }}
                            >
                              {action?.icon && action?.icon}
                              {action?.iconTemplate && action?.iconTemplate(item)}
                              {action?.iconUrl && (
                                <Image
                                  src={action.iconUrl}
                                  alt={action?.text}
                                  height={24}
                                  width={24}
                                />
                              )}
                              {action?.text && (
                                <Typography color="primary" pl={0.5}>
                                  {action?.text}
                                </Typography>
                              )}
                            </StyledActionButton>
                          ))}
                        {actions.filter((action) => action.isMenuAction).length > 0 && (
                          <StyledIconButton
                            onClick={(event: MouseEvent<HTMLElement>) => {
                              handleActionMenuClick(event, item);
                            }}
                            color="secondary"
                            disableRipple
                          >
                            <StyledMenuIcon />
                          </StyledIconButton>
                        )}
                      </Stack>
                    </StyledTableCell>
                  )}
                  {(deleteBtn || editBtn) && (
                    <StyledTableCell $isFixed={stickyAction} $border={border}>
                      <FlexCenterRow gap={0.5} justifyContent={actionDirection}>
                        {deleteBtn && (
                          <Icon
                            name={ICON_NAMES.DELETE}
                            onClick={() => {
                              handleOnMenuItemClick(deleteBtn, item);
                            }}
                          />
                        )}
                        {editBtn && (
                          <Icon
                            name={ICON_NAMES.EDIT}
                            onClick={() => {
                              handleOnMenuItemClick(editBtn, item);
                            }}
                          />
                        )}
                      </FlexCenterRow>
                    </StyledTableCell>
                  )}
                </StyledTableRow>
              ))}
            {!loading && !data?.length && (
              <StyledTableRow $border={border} $borderColor={borderColor}>
                <StyledTableCell colSpan={headers?.length} align="center" $border={border}>
                  <StyledInputLabel>{emptyMessage || 'No records found!'}</StyledInputLabel>
                </StyledTableCell>
              </StyledTableRow>
            )}
          </TableBody>
        </Table>
        {/* Action menu*/}
        {actions.length > 0 && (
          <StyledMenu
            anchorEl={anchorEl}
            open={open}
            onClose={handleActionMenuClose}
            anchorOrigin={{ vertical: 'bottom', horizontal: 'left' }}
            transformOrigin={{ vertical: 'top', horizontal: 'left' }}
          >
            {open &&
              actions
                .filter(
                  (action) =>
                    (action?.filterAction && action?.filterAction?.(actionMenu?.item)) ||
                    !action?.filterAction
                )
                .map((action, ai) => {
                  const { item } = actionMenu;
                  return (
                    <StyledMenuItem
                      key={`${action}-${ai}`}
                      onClick={() => {
                        handleOnMenuItemClick(action, item);
                      }}
                    >
                      {action?.text}
                      {action?.textTemplate && action?.textTemplate(item)}
                    </StyledMenuItem>
                  );
                })}
          </StyledMenu>
        )}
      </StyledTableContainer>

      {showPagination && totalPages !== 0 && (
        <TablePagination
          mode={paginationMode}
          page={currentPage}
          rowsPerPage={rowsPerPage}
          totalPages={totalPages}
          onPageChange={onPageChange as () => void}
          component="div"
          border={paginationborder}
          borderColor={borderColor}
          bgColor={bgColor}
        />
      )}

      {/* Action confirmation popup */}
      <ConfirmationDialog
        open={confirmAction.openModal}
        onClose={onActionConfirmationClose}
        onSubmit={handleActionConfirmation}
        message={confirmAction.actionMessage || 'Are you absolutely Sure?'}
        id={confirmAction.actionId}
        reasonPlaceholder={confirmReasonPlaceholder}
        submitLabel={confirmAction.actionLabel}
      />
      {/* Request confirmation popup */}
      <AlertDialog
        open={proceedConfirmationModalOpen}
        onClose={onProceedConfirmationClose}
        title={alertDialogProps.title ?? 'Are you sure?'}
        subTitle={alertDialogProps.subTitle}
        onSubmit={handleProceedConfirmation}
        submitLabel={alertDialogProps.submitLabel}
      />
    </React.Fragment>
  );
}
