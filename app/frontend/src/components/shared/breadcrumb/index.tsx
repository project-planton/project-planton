'use client';
import { FC, ReactNode } from 'react';
import { Typography, Skeleton, SvgIconProps } from '@mui/material';
import { Icon, ICON_NAMES } from '@/components/shared/icon';
import { BreadcrumbLabel, StyledBreadcrumbs, StyledBreadcrumbStartIcon } from './styled';

export interface IBreadcrumbItem {
  name: string;
  handler?: () => void;
}

interface IBreadcrumb {
  breadcrumbs: IBreadcrumbItem[];
  startBreadcrumb?: ReactNode;
}

interface IBreadcrumbStartIcon {
  icon: ICON_NAMES;
  iconProps?: SvgIconProps;
  label: string;
  handler?: () => void;
}

export const BreadcrumbStartIcon: FC<IBreadcrumbStartIcon> = ({
  icon,
  iconProps,
  label,
  handler,
}) => {
  return (
    <StyledBreadcrumbStartIcon onClick={handler}>
      <Icon name={icon} {...iconProps} />
      <Typography variant="caption">{label}</Typography>
    </StyledBreadcrumbStartIcon>
  );
};

export const Breadcrumb: FC<IBreadcrumb> = ({ breadcrumbs, startBreadcrumb }) => {
  const isEmpty = breadcrumbs.length === 0;

  return (
    <StyledBreadcrumbs aria-label="breadcrumb">
      {isEmpty ? (
        <BreadcrumbLabel>
          <Skeleton variant="text" width={120} height={20} />
        </BreadcrumbLabel>
      ) : (
        startBreadcrumb
      )}
      {isEmpty
        ? [80, 100, 60].map((width, index) => (
            <BreadcrumbLabel key={`skeleton-${index}`}>
              <Skeleton variant="text" width={width} height={20} />
            </BreadcrumbLabel>
          ))
        : breadcrumbs.map(({ name, handler }, index) => (
            <BreadcrumbLabel key={`${name}-${index}`} onClick={handler} $hasLink={!!handler}>
              {name}
            </BreadcrumbLabel>
          ))}
    </StyledBreadcrumbs>
  );
};
