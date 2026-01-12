'use client';
import React, { ReactNode } from 'react';
import { Box, Skeleton, styled, alpha, Typography } from '@mui/material';
import { TrendingUp, TrendingDown, ArrowForward } from '@mui/icons-material';
import Link from 'next/link';

interface StatCardProps {
  title: string;
  value: string | number | null;
  icon?: ReactNode;
  trend?: {
    value: string;
    positive: boolean;
  };
  subtitle?: string;
  loading?: boolean;
  accent?: boolean;
  href?: string;
  onClick?: () => void;
}

const CardContainer = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$accent' && prop !== '$clickable',
})<{ $accent?: boolean; $clickable?: boolean }>(({ theme, $accent, $clickable }) => ({
  padding: theme.spacing(3),
  borderRadius: 16,
  height: '100%',
  minHeight: 160,
  boxSizing: 'border-box',
  display: 'flex',
  flexDirection: 'column',
  position: 'relative',
  backgroundColor: $accent
    ? theme.palette.mode === 'dark'
      ? alpha(theme.palette.primary.main, 0.08)
      : alpha(theme.palette.primary.main, 0.04)
    : theme.palette.background.paper,
  border: `1px solid ${
    $accent
      ? theme.palette.mode === 'dark'
        ? alpha(theme.palette.primary.main, 0.25)
        : alpha(theme.palette.primary.main, 0.15)
      : theme.palette.divider
  }`,
  transition: 'all 200ms cubic-bezier(0.4, 0, 0.2, 1)',
  cursor: $clickable ? 'pointer' : 'default',
  textDecoration: 'none',
  overflow: 'hidden',

  // Subtle gradient overlay for depth
  '&::before': $accent ? {
    content: '""',
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    background: theme.palette.mode === 'dark'
      ? `linear-gradient(135deg, ${alpha(theme.palette.primary.main, 0.12)} 0%, transparent 50%)`
      : `linear-gradient(135deg, ${alpha(theme.palette.primary.main, 0.08)} 0%, transparent 50%)`,
    pointerEvents: 'none',
  } : {},

  '&:hover': {
    borderColor: $accent
      ? theme.palette.primary.main
      : theme.palette.mode === 'dark'
        ? alpha(theme.palette.common.white, 0.2)
        : alpha(theme.palette.common.black, 0.15),
    boxShadow: $accent
      ? `0 8px 32px ${alpha(theme.palette.primary.main, 0.25)}`
      : theme.palette.mode === 'dark'
        ? '0 8px 32px rgba(0, 0, 0, 0.35)'
        : '0 8px 32px rgba(0, 0, 0, 0.08)',
    transform: $clickable ? 'translateY(-3px)' : 'none',

    '& .stat-card-arrow': {
      opacity: 1,
      transform: 'translateX(0)',
    },
  },
}));

const CardHeader = styled(Box)({
  display: 'flex',
  alignItems: 'flex-start',
  justifyContent: 'space-between',
  marginBottom: 'auto',
});

const IconContainer = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$accent',
})<{ $accent?: boolean }>(({ theme, $accent }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 48,
  height: 48,
  borderRadius: 14,
  backgroundColor: $accent
    ? theme.palette.mode === 'dark'
      ? alpha(theme.palette.primary.main, 0.2)
      : alpha(theme.palette.primary.main, 0.12)
    : theme.palette.mode === 'dark'
      ? alpha(theme.palette.text.secondary, 0.1)
      : alpha(theme.palette.text.secondary, 0.08),
  color: $accent
    ? theme.palette.primary.main
    : theme.palette.text.secondary,
  transition: 'all 200ms ease',

  '& svg': {
    fontSize: 24,
  },
}));

const ArrowContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 28,
  height: 28,
  borderRadius: 8,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.06)
    : alpha(theme.palette.common.black, 0.04),
  color: theme.palette.text.secondary,
  opacity: 0,
  transform: 'translateX(-4px)',
  transition: 'all 200ms cubic-bezier(0.4, 0, 0.2, 1)',

  '& svg': {
    fontSize: 16,
  },
}));

const ContentSection = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(2),
}));

const Title = styled(Typography)(({ theme }) => ({
  fontSize: '0.8125rem',
  fontWeight: 500,
  color: theme.palette.text.secondary,
  marginBottom: theme.spacing(0.75),
  lineHeight: 1.4,
  letterSpacing: '0.01em',
  textTransform: 'uppercase',
}));

const Value = styled(Typography)(({ theme }) => ({
  fontSize: '2.25rem',
  fontWeight: 700,
  color: theme.palette.text.primary,
  lineHeight: 1.1,
  letterSpacing: '-0.02em',
}));

const TrendBadge = styled(Box, {
  shouldForwardProp: (prop) => prop !== '$positive',
})<{ $positive?: boolean }>(({ theme, $positive }) => ({
  display: 'inline-flex',
  alignItems: 'center',
  gap: theme.spacing(0.5),
  marginTop: theme.spacing(1.5),
  padding: theme.spacing(0.5, 1),
  borderRadius: 6,
  fontSize: '0.75rem',
  fontWeight: 600,
  backgroundColor: $positive
    ? alpha(theme.palette.success.main, 0.12)
    : alpha(theme.palette.error.main, 0.12),
  color: $positive
    ? theme.palette.success.main
    : theme.palette.error.main,

  '& svg': {
    fontSize: 14,
  },
}));

const Subtitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.75rem',
  color: theme.palette.text.secondary,
  marginTop: theme.spacing(1),
}));

export const StatCard: React.FC<StatCardProps> = ({
  title,
  value,
  icon,
  trend,
  subtitle,
  loading = false,
  accent = false,
  href,
  onClick,
}) => {
  const isClickable = !!href || !!onClick;

  const cardContent = (
    <CardContainer
      $accent={accent}
      $clickable={isClickable}
      onClick={onClick}
    >
      <CardHeader>
        {icon && <IconContainer $accent={accent}>{icon}</IconContainer>}
        {isClickable && (
          <ArrowContainer className="stat-card-arrow">
            <ArrowForward />
          </ArrowContainer>
        )}
      </CardHeader>

      <ContentSection>
        <Title>{title}</Title>

        {loading ? (
          <Skeleton 
            variant="text" 
            width={100} 
            height={48} 
            sx={{ borderRadius: 1 }}
          />
        ) : (
          <Value>{value ?? '—'}</Value>
        )}

        {trend && !loading && (
          <TrendBadge $positive={trend.positive}>
            {trend.positive ? <TrendingUp /> : <TrendingDown />}
            {trend.value}
          </TrendBadge>
        )}

        {subtitle && !loading && <Subtitle>{subtitle}</Subtitle>}
      </ContentSection>
    </CardContainer>
  );

  if (href) {
    return (
      <Link href={href} style={{ textDecoration: 'none', display: 'block', height: '100%' }}>
        {cardContent}
      </Link>
    );
  }

  return cardContent;
};
