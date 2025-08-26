import { createClient } from '@base44/sdk';
// import { getAccessToken } from '@base44/sdk/utils/auth-utils';

// Create a client with authentication required
export const base44 = createClient({
  appId: "68a140b1ea5b74497f609f38", 
  requiresAuth: true // Ensure authentication is required for all operations
});
