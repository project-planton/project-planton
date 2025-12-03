import { Timestamp } from 'node_modules/@bufbuild/protobuf/dist/esm/wkt/gen/google/protobuf/timestamp_pb';
import moment from 'moment';

export const Utils = {
  setStorage(key: string, data: any): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem(key, JSON.stringify(data));
    }
  },
  getStorage(key: string): any {
    if (typeof window !== 'undefined') {
      try {
        const data = localStorage.getItem(key);
        return data ? safeParseJson(data) : undefined;
      } catch {
        return;
      }
    } else {
      return;
    }
  },
  removeStorageItem(key: string): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem(key);
    }
  },
  clearStorage(): void {
    if (typeof window !== 'undefined') {
      localStorage.clear();
    }
  },
};

export function safeParseJson(json: any): any {
  try {
    return JSON.parse(typeof json === 'object' ? JSON.stringify(json, null, 2) : json);
  } catch {
    return;
  }
}

export const placeholderErrHandler = () => {
  /**
   * Placeholder err handler for API catch.
   * If required to perform any other action on API error, then use your own error function.
   */
};

export const formatTimestampToDate = (timestamp: Timestamp, format = 'DD/MM/YYYY, HH:mm:ss') => {
  let date: Date;
  if (typeof timestamp === 'string') date = new Date(timestamp);
  else if (timestamp?.seconds) date = new Date(Number(timestamp.seconds) * 1000);
  return moment(date).format(format);
};

/**
 * Get data from nested object path (e.g., 'user.profile.name')
 */
export function getDataFromPath<T>(path: string, data: T): any {
  if (!path) {
    return '';
  }

  const pathArr = path?.split('.');
  if (pathArr.length === 1) {
    return (data && (data as any)[pathArr[0]]) || '';
  }
  const parsed = pathArr.shift();
  return getDataFromPath(pathArr.join('.'), data ? (data as any)[parsed] : null);
}

/**
 * Copy text to clipboard
 */
export const copyText = async (text: string): Promise<void> => {
  if (typeof navigator !== 'undefined' && navigator.clipboard) {
    await navigator.clipboard.writeText(text);
  }
};

/**
 * Capitalize words in a string (e.g., 'hello_world' -> 'Hello World')
 */
export const capitalizeWords = (str: string): string => {
  return str
    .split('_')
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join(' ');
};
