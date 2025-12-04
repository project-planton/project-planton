'use client';

import { Drawer } from '@/components/shared/drawer';
import { StackJobsList } from './stack-jobs-list';
export interface StackJobsDrawerProps {
  open: boolean;
  cloudResourceId: string;
  onClose: () => void;
}

export function StackJobsDrawer({ open, cloudResourceId, onClose }: StackJobsDrawerProps) {
  return (
    <Drawer open={open} onClose={onClose} title="Stack Jobs" width={900}>
      <StackJobsList cloudResourceId={cloudResourceId} />
    </Drawer>
  );
}

