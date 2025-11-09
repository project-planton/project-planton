"use client";

import Image from "next/image";
import { useEffect, useState } from "react";

interface GitHubStarBadgeProps {
  repo: string; // Format: "owner/repo"
  className?: string;
}

export function GitHubStarBadge({ repo, className = "" }: GitHubStarBadgeProps) {
  const [stars, setStars] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchStars() {
      try {
        const response = await fetch(`https://api.github.com/repos/${repo}`);
        if (response.ok) {
          const data = await response.json();
          setStars(data.stargazers_count);
        }
      } catch (error) {
        console.error("Failed to fetch GitHub stars:", error);
      } finally {
        setLoading(false);
      }
    }

    fetchStars();
  }, [repo]);

  const formatStars = (count: number): string => {
    if (count >= 1000) {
      return `${(count / 1000).toFixed(1)}k`;
    }
    return count.toString();
  };

  return (
    <a
      href={`https://github.com/${repo}`}
      target="_blank"
      rel="noopener noreferrer"
      className={`inline-flex items-center gap-2 px-3 py-1.5 rounded-md border border-slate-700 bg-slate-800/50 text-slate-300 hover:bg-slate-700 hover:text-white hover:border-slate-600 transition-all duration-200 ${className}`}
    >
      <Image src="/images/providers/github-dark.svg" alt="GitHub" width={16} height={16} />
      <span className="text-sm font-medium">Star</span>
      {!loading && stars !== null && (
        <>
          <span className="w-px h-4 bg-slate-600" />
          <span className="text-sm font-semibold">{formatStars(stars)}</span>
        </>
      )}
    </a>
  );
}

