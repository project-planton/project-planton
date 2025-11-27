import { useContext, useEffect, useState } from 'react';
import { createClient, Client } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import { AppContext } from '@/contexts';
import { GenService } from '@bufbuild/protobuf/codegenv1';

function addGlobalErrorHandling(client: any) {
  for (const methodName of Object.keys(client)) {
    if (typeof client[methodName] === 'function') {
      const originalMethod = client[methodName];
      client[methodName] = (...args: any[]) => {
        try {
          const result = originalMethod(...args);
          if (result && typeof result.then === 'function') {
            return result.catch((error: any) => {
              if (error.message) error.message = error?.message?.replace(/^\[.*?\]\s*/, '');
              throw error;
            });
          }
          return result;
        } catch (error: any) {
          if (error.message) error.message = error?.message?.replace(/^\[.*?\]\s*/, '');
          throw error;
        }
      };
    }
  }

  return client;
}

export const useConnectRpcClient = <T extends GenService<any>>(service: T): Client<T> | null => {
  const { connectHost } = useContext(AppContext);
  const [client, setClient] = useState<Client<T> | null>(null);

  useEffect(() => {
    if (!service || !connectHost) return;

    const transport = createGrpcWebTransport({
      baseUrl: connectHost,
      useBinaryFormat: true, // Binary format now works with proper field definitions in schemas
    });

    const newClient = createClient(service, transport);

    const clientWithGlobalErrorHandling = addGlobalErrorHandling(newClient);

    setClient(clientWithGlobalErrorHandling);
  }, [connectHost, service]);

  return client;
};
