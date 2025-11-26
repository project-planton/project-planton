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

