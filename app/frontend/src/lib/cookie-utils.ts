import { COOKIE_KEYS } from '@/lib/cookie-constants';

/**
 * Set a cookie (Client-side)
 * @param name - Cookie name
 * @param value - Cookie value (will be JSON stringified if not a string)
 * @param days - Number of days until cookie expires (default: 365)
 */
export function setCookie(name: string, value: any, days = 365) {
  if (typeof document === 'undefined') return;

  const expires = new Date();
  expires.setTime(expires.getTime() + days * 24 * 60 * 60 * 1000);

  const stringValue = typeof value === 'string' ? value : JSON.stringify(value);
  const cookie = `${name}=${encodeURIComponent(stringValue)};expires=${expires.toUTCString()};path=/;sameSite=strict`;

  document.cookie = cookie;
}

/**
 * Delete a cookie (Client-side)
 * @param name - Cookie name to delete
 */
export function deleteCookie(name: string) {
  if (typeof document === 'undefined') return;
  document.cookie = `${name}=;max-age=0;path=/;sameSite=strict`;
}

/**
 * Set navbar open state in cookie (Client-side)
 * @param isOpen - Whether navbar is open
 */
export function setCookieNavbarOpen(isOpen: boolean): void {
  setCookie(COOKIE_KEYS.NAVBAR_OPEN, isOpen);
}

/**
 * Set theme mode in cookie (Client-side)
 * @param themeMode - Theme mode ('dark' | 'light')
 */
export function setCookieThemeMode(themeMode: 'dark' | 'light'): void {
  setCookie(COOKIE_KEYS.THEME_MODE, themeMode);
}

/**
 * Clear all cookies (Client-side)
 * Useful for logout to ensure all cookies are removed
 */
export function clearAllCookies(): void {
  if (typeof document === 'undefined') return;

  const cookies = document.cookie.split(';');

  for (const cookie of cookies) {
    const eqPos = cookie.indexOf('=');
    const name = eqPos > -1 ? cookie.substring(0, eqPos).trim() : cookie.trim();

    if (name) {
      deleteCookie(name);
    }
  }
}
