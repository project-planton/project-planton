import React from 'react';
import { ButtonProps } from '@mui/material';

export type DotPrefix<T extends string> = T extends '' ? '' : `.${T}`;

export type DotNestedKeys<T> = T extends Date | ((...args: any) => any) | Array<any>
  ? ''
  : T extends object
    ? {
        [K in Exclude<keyof T, symbol>]: T[K] extends object
          ? `${K}` | `${K}${DotPrefix<DotNestedKeys<T[K]>>}`
          : `${K}`;
      }[Exclude<keyof T, symbol>]
    : '';

export enum PAGINATION_MODE {
  CLIENT = 'client',
  SERVER = 'server',
}

export enum ActionType {
  UNSPECIFIED = 'UNSPECIFIED',
  DELETE = 'DELETE',
  REFRESH = 'REFRESH',
}

export interface ActionMenuProps<T> {
  text?: string;
  textTemplate?: (row: T) => React.ReactNode;
  btnProps?: ButtonProps;
  handler?: (row: T, data?: any) => void;
  previewHandler?: (row: T, reason?: string) => void;
  type?: ActionType;
  isMenuAction?: boolean;
  icon?: React.ReactNode;
  iconUrl?: string;
  iconTemplate?: (row: T) => React.ReactNode;
  requiresProceedConfirmation?: boolean;
  label?: string;
  filterAction?: (row: T) => boolean;
  message?: string;
  deleteMessageKey?: DotNestedKeys<Partial<T>>;
  skipDeleteIdCheck?: boolean;
}

