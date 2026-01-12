'use client';
import React from 'react';
import { Box, styled, keyframes, alpha } from '@mui/material';

const shimmer = keyframes`
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
`;

interface ShimmerSkeletonProps {
  width?: number | string;
  height?: number | string;
  borderRadius?: number | string;
  className?: string;
}

const SkeletonBase = styled(Box, {
  shouldForwardProp: (prop) => 
    prop !== '$width' && prop !== '$height' && prop !== '$borderRadius',
})<{
  $width?: number | string;
  $height?: number | string;
  $borderRadius?: number | string;
}>(({ theme, $width = '100%', $height = 20, $borderRadius = 6 }) => ({
  width: typeof $width === 'number' ? `${$width}px` : $width,
  height: typeof $height === 'number' ? `${$height}px` : $height,
  borderRadius: typeof $borderRadius === 'number' ? `${$borderRadius}px` : $borderRadius,
  backgroundColor: theme.palette.mode === 'dark'
    ? alpha(theme.palette.common.white, 0.08)
    : alpha(theme.palette.common.black, 0.06),
  backgroundImage: `linear-gradient(
    90deg,
    transparent 0%,
    ${theme.palette.mode === 'dark'
      ? alpha(theme.palette.common.white, 0.08)
      : alpha(theme.palette.common.white, 0.4)} 50%,
    transparent 100%
  )`,
  backgroundSize: '200% 100%',
  animation: `${shimmer} 1.5s infinite ease-in-out`,
}));

export const ShimmerSkeleton: React.FC<ShimmerSkeletonProps> = ({
  width,
  height,
  borderRadius,
  className,
}) => {
  return (
    <SkeletonBase
      $width={width}
      $height={height}
      $borderRadius={borderRadius}
      className={className}
    />
  );
};

// Pre-built skeleton variants
export const TextSkeleton = styled(ShimmerSkeleton)({});

export const CircleSkeleton = styled(ShimmerSkeleton)({
  borderRadius: '50%',
});

export const CardSkeleton = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2.5),
  borderRadius: 12,
  backgroundColor: theme.palette.background.paper,
  border: `1px solid ${theme.palette.divider}`,
}));

export const TableRowSkeleton: React.FC<{ columns?: number }> = ({ columns = 5 }) => {
  return (
    <Box sx={{ display: 'flex', gap: 2, py: 1.5, px: 2 }}>
      {Array.from({ length: columns }).map((_, i) => (
        <ShimmerSkeleton
          key={i}
          width={i === 0 ? 40 : `${100 / columns}%`}
          height={16}
        />
      ))}
    </Box>
  );
};

export const StatCardSkeleton: React.FC = () => {
  return (
    <CardSkeleton>
      <ShimmerSkeleton width={40} height={40} borderRadius={10} />
      <Box sx={{ mt: 2 }}>
        <ShimmerSkeleton width={80} height={14} />
      </Box>
      <Box sx={{ mt: 1 }}>
        <ShimmerSkeleton width={60} height={28} />
      </Box>
    </CardSkeleton>
  );
};
