'use client';

import { Drawer } from '@/components/shared/drawer';
import { StackUpdatesList } from './stack-jobs-list';
export interface StackUpdatesDrawerProps {
  open: boolean;
  cloudResourceId: string;
  onClose: () => void;
}

export function StackUpdatesDrawer({ open, cloudResourceId, onClose }: StackUpdatesDrawerProps) {
  return (
    <Drawer open={open} onClose={onClose} title="Stack Jobs" width={900}>
      <StackUpdatesList cloudResourceId={cloudResourceId} />
    </Drawer>
  );
}

