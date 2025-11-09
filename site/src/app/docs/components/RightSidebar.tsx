'use client';

import React, { useEffect, useState } from 'react';
import { Box, Typography } from '@mui/material';
import { Author } from '@/lib/mdx';

interface RightSidebarProps {
  author?: Author[];
  content?: string;
}

interface Heading {
  id: string;
  text: string;
  level: number;
}

const RightSidebar: React.FC<RightSidebarProps> = ({ author = [], content }) => {
  const [headings, setHeadings] = useState<Heading[]>([]);
  const [activeId, setActiveId] = useState<string>('');

  useEffect(() => {
    if (!content) return;

    // Extract headings from markdown content
    const lines = content.split('\n');
    const extractedHeadings: Heading[] = [];

    lines.forEach((line) => {
      const match = line.match(/^(#{1,6})\s+(.+)$/);
      if (match) {
        const level = match[1].length;
        const text = match[2];
        const id = text
          .toLowerCase()
          .replace(/[^a-z0-9\s-]/g, '')
          .replace(/\s+/g, '-');

        // Only show h2 and h3 in TOC
        if (level === 2 || level === 3) {
          extractedHeadings.push({ id, text, level });
        }
      }
    });

    setHeadings(extractedHeadings);
  }, [content]);

  useEffect(() => {
    // Track active heading based on scroll position
    const handleScroll = () => {
      const headingElements = headings.map((h) => ({
        id: h.id,
        element: document.getElementById(h.id),
      }));

      let currentActiveId = '';
      
      for (const { id, element } of headingElements) {
        if (element) {
          const rect = element.getBoundingClientRect();
          if (rect.top <= 100) {
            currentActiveId = id;
          }
        }
      }

      setActiveId(currentActiveId);
    };

    window.addEventListener('scroll', handleScroll);
    handleScroll(); // Initial check

    return () => window.removeEventListener('scroll', handleScroll);
  }, [headings]);

  const scrollToHeading = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      const yOffset = -80; // Offset for fixed header
      const y = element.getBoundingClientRect().top + window.pageYOffset + yOffset;
      window.scrollTo({ top: y, behavior: 'smooth' });
    }
  };

  return (
    <Box className="p-6">
      {/* Table of Contents */}
      {headings.length > 0 && (
        <Box className="mb-6">
          <Typography variant="subtitle2" className="text-gray-400 font-semibold mb-3 uppercase text-xs">
            On This Page
          </Typography>
          <nav>
            <ul className="space-y-2">
              {headings.map((heading) => (
                <li
                  key={heading.id}
                  className={`${heading.level === 3 ? 'ml-4' : ''}`}
                >
                  <button
                    onClick={() => scrollToHeading(heading.id)}
                    className={`text-sm text-left w-full hover:text-purple-400 transition-colors ${
                      activeId === heading.id
                        ? 'text-purple-400 font-medium'
                        : 'text-gray-400'
                    }`}
                  >
                    {heading.text}
                  </button>
                </li>
              ))}
            </ul>
          </nav>
        </Box>
      )}

      {/* Author Information */}
      {author && author.length > 0 && (
        <Box className="border-t border-purple-900/30 pt-6">
          <Typography variant="subtitle2" className="text-gray-400 font-semibold mb-3 uppercase text-xs">
            Author{author.length > 1 ? 's' : ''}
          </Typography>
          <div className="space-y-3">
            {author.map((a, index) => (
              <div key={index} className="flex items-start gap-3">
                {a.image && (
                  <img
                    src={a.image}
                    alt={a.name}
                    className="w-10 h-10 rounded-full"
                  />
                )}
                <div className="flex-1">
                  <Typography className="text-white text-sm font-medium">
                    {a.name}
                  </Typography>
                  {a.role && (
                    <Typography className="text-gray-400 text-xs">
                      {a.role}
                    </Typography>
                  )}
                </div>
              </div>
            ))}
          </div>
        </Box>
      )}
    </Box>
  );
};

export default RightSidebar;

