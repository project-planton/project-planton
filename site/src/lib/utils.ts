import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: Array<string | undefined | false | null>): string {
  return twMerge(clsx(inputs));
}


