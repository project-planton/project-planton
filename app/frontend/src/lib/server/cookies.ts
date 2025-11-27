import { cookies } from 'next/headers';
import { COOKIE_KEYS } from '@/lib/cookie-constants';

/**
 * Get all cookies with their parsed values
 * Attempts to parse JSON values, returns raw string if parsing fails
 */
export function getAllCookiesParsed(): Record<string, any> {
  const cookieStore = cookies();
  const allCookies: Record<string, any> = {};

  cookieStore.getAll().forEach((cookie) => {
    try {
      // Try to parse as JSON
      allCookies[cookie.name] = JSON.parse(cookie.value);
    } catch {
      // If not JSON, store as string
      allCookies[cookie.name] = cookie.value;
    }
  });

  return allCookies;
}

/**
 * Get navbar open state from cookies (Server-side)
 */
export function getCookieNavbarOpen(cookies?: Record<string, any>): boolean | null {
  if (!cookies) return null;
  const value = cookies[COOKIE_KEYS.NAVBAR_OPEN];
  return value !== undefined ? value : null;
}

/**
 * Get theme mode from cookies (Server-side)
 */
export function getCookieThemeMode(cookies?: Record<string, any>): 'dark' | 'light' | null {
  if (!cookies) return null;
  const value = cookies[COOKIE_KEYS.THEME_MODE];
  return value !== undefined ? value : null;
}

/**
 * Set a cookie (Server-side)
 * @param name - Cookie name
 * @param value - Cookie value (will be JSON stringified if not a string)
 * @param days - Number of days until cookie expires (default: 365)
 */
export function setCookie(name: string, value: any, days = 365): void {
  const cookieStore = cookies();
  const stringValue = typeof value === 'string' ? value : JSON.stringify(value);

  const maxAge = days * 24 * 60 * 60; // Convert days to seconds

  cookieStore.set(name, stringValue, {
    path: '/',
    maxAge,
    sameSite: 'strict',
    httpOnly: false,
  });
}

/**
 * Delete a specific cookie (Server-side)
 */
export function deleteCookie(name: string): void {
  const cookieStore = cookies();
  cookieStore.delete(name);
}

/**
 * Delete all app cookies (Server-side)
 * Note: We use cookieStore.delete() directly here instead of calling deleteCookie()
 * to avoid multiple cookies() calls which would be inefficient
 */
export function clearAllAppCookies(): void {
  const cookieStore = cookies();
  const allCookies = cookieStore.getAll();

  // Delete each cookie individually
  allCookies.forEach((cookie) => {
    cookieStore.delete(cookie.name);
  });
}
