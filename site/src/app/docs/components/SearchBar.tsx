'use client';

import React, { useState, useEffect, useDeferredValue, useRef } from 'react';
import { TextField, InputAdornment, Popper, Paper, List, ListItem, ListItemButton, ListItemText, Typography, Box } from '@mui/material';
import { Search as SearchIcon, InfoOutlined as InfoIcon } from '@mui/icons-material';
import { useRouter } from 'next/navigation';
import { addBasePath } from 'next/dist/client/add-base-path';

// Type definitions for Pagefind
type PagefindOptions = {
  baseUrl?: string;
};

declare global {
  interface Window {
    pagefind?: {
      options: (opts: PagefindOptions) => Promise<void>;
      debouncedSearch: <T>(query: string) => Promise<{
        results: Array<{ data: () => Promise<T> }>;
      } | null>;
    };
  }
}

type PagefindResult = {
  excerpt: string;
  meta: {
    title: string;
  };
  url: string;
  sub_results: {
    excerpt: string;
    title: string;
    url: string;
  }[];
};

const INPUTS = new Set(['INPUT', 'SELECT', 'BUTTON', 'TEXTAREA']);

const DEV_SEARCH_NOTICE = (
  <Box sx={{ p: 2, textAlign: 'left' }}>
    <Typography variant="body2" sx={{ mb: 1 }}>
      Search isn&apos;t available in development because Pagefind indexes built HTML files instead of markdown source files.
    </Typography>
    <Typography variant="body2">
      To test search during development, run <code>yarn build</code> and then <code>yarn start</code>.
    </Typography>
  </Box>
);

async function importPagefind() {
  // Dynamic import of pagefind from the built static files
  // Use addBasePath to ensure correct path in production (GitHub Pages)
  const pagefindPath = addBasePath('/_pagefind/pagefind.js');
  window.pagefind = await import(
    /* webpackIgnore: true */ pagefindPath
  ) as typeof window.pagefind;
  await window.pagefind!.options({
    baseUrl: '/',
  });
}

export const SearchBar: React.FC = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | React.ReactElement>('');
  const [results, setResults] = useState<PagefindResult[]>([]);
  const [focused, setFocused] = useState(false);
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [mounted, setMounted] = useState(false);
  
  const deferredSearch = useDeferredValue(searchQuery);
  const inputRef = useRef<HTMLInputElement>(null);
  const router = useRouter();

  // Track mounted state for SSR
  useEffect(() => {
    setMounted(true);
  }, []);

  // Handle search
  useEffect(() => {
    const handleSearch = async (value: string) => {
      if (!value) {
        setResults([]);
        setError('');
        return;
      }
      
      setIsLoading(true);
      
      if (!window.pagefind) {
        try {
          await importPagefind();
        } catch (error) {
          const message =
            error instanceof Error
              ? process.env.NODE_ENV !== 'production' &&
                error.message.includes('Failed to fetch')
                ? DEV_SEARCH_NOTICE
                : `${error.constructor.name}: ${error.message}`
              : String(error);
          setError(message);
          setIsLoading(false);
          return;
        }
      }
      
      const response = await window.pagefind!.debouncedSearch<PagefindResult>(
        value
      );
      
      if (!response) return;

      const data = await Promise.all(response.results.map((o) => o.data()));
      setIsLoading(false);
      setError('');
      setResults(
        data.map((newData) => ({
          ...newData,
          sub_results: newData.sub_results.map((r) => {
            const url = r.url.replace(/\.html$/, '').replace(/\.html#/, '#');
            return { ...r, url };
          }),
        }))
      );
    };

    handleSearch(deferredSearch);
  }, [deferredSearch]);

  // Keyboard shortcuts
  useEffect(() => {
    function handleKeyDown(event: KeyboardEvent) {
      const el = document.activeElement;
      if (
        !el ||
        INPUTS.has(el.tagName) ||
        (el as HTMLElement).isContentEditable
      ) {
        return;
      }
      if (
        event.key === '/' ||
        (event.key === 'k' &&
          !event.shiftKey &&
          (navigator.userAgent.includes('Mac') ? event.metaKey : event.ctrlKey))
      ) {
        event.preventDefault();
        inputRef.current?.focus({ preventScroll: true });
      }
    }

    window.addEventListener('keydown', handleKeyDown);
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, []);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = event.target;
    setSearchQuery(value);
    setAnchorEl(event.currentTarget);
  };

  const handleFocus = (event: React.FocusEvent<HTMLInputElement>) => {
    setFocused(true);
    setAnchorEl(event.currentTarget);
  };

  const handleBlur = () => {
    // Delay to allow click on results
    setTimeout(() => {
      setFocused(false);
      setAnchorEl(null);
    }, 200);
  };

  const handleResultClick = (result: PagefindResult['sub_results'][0]) => {
    inputRef.current?.blur();
    const [url, hash] = result.url.split('#');
    const isSamePathname = location.pathname === url;
    
    if (isSamePathname && hash) {
      location.href = `#${hash}`;
    } else {
      router.push(result.url);
    }
    setSearchQuery('');
    setAnchorEl(null);
  };

  const open = Boolean(anchorEl) && focused && (searchQuery.length > 0);
  const showResults = !error && !isLoading && results.length > 0;
  const showEmpty = !error && !isLoading && deferredSearch && results.length === 0;

  return (
    <>
      <TextField
        inputRef={inputRef}
        size="small"
        placeholder="Search documentation..."
        value={searchQuery}
        onChange={handleChange}
        onFocus={handleFocus}
        onBlur={handleBlur}
        className="w-64"
        sx={{
          '& .MuiOutlinedInput-root': {
            color: 'white',
            position: 'relative',
            '& fieldset': {
              borderColor: 'rgba(168, 85, 247, 0.3)',
            },
            '&:hover fieldset': {
              borderColor: 'rgba(168, 85, 247, 0.5)',
            },
            '&.Mui-focused fieldset': {
              borderColor: 'rgba(168, 85, 247, 0.8)',
            },
          },
          '& .MuiInputBase-input::placeholder': {
            color: 'rgba(255, 255, 255, 0.5)',
            opacity: 1,
          },
        }}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <SearchIcon className="text-gray-400" />
            </InputAdornment>
          ),
          endAdornment: (
            <InputAdornment position="end">
              <Typography
                variant="caption"
                sx={{
                  display: { xs: 'none', sm: 'block' },
                  color: 'rgba(255, 255, 255, 0.4)',
                  fontSize: '11px',
                  fontFamily: 'monospace',
                  border: '1px solid rgba(168, 85, 247, 0.3)',
                  borderRadius: '4px',
                  px: 0.5,
                  py: 0.25,
                  opacity: mounted && !focused ? 1 : 0,
                  transition: 'opacity 0.2s',
                }}
              >
                {mounted && navigator.userAgent.includes('Mac') ? '⌘K' : 'CTRL K'}
              </Typography>
            </InputAdornment>
          ),
        }}
      />
      
      <Popper
        open={open}
        anchorEl={anchorEl}
        placement="bottom-start"
        sx={{ zIndex: 1400, width: { xs: '100%', sm: 576 }, maxWidth: '90vw' }}
      >
        <Paper
          sx={{
            mt: 1,
            maxHeight: { xs: '70vh', md: 400 },
            overflow: 'auto',
            bgcolor: 'rgba(30, 41, 59, 0.95)',
            backdropFilter: 'blur(10px)',
            border: '1px solid rgba(168, 85, 247, 0.2)',
            boxShadow: '0 10px 40px rgba(0, 0, 0, 0.3)',
            '&::-webkit-scrollbar': {
              width: '8px',
            },
            '&::-webkit-scrollbar-track': {
              background: 'rgba(15, 23, 42, 0.5)',
            },
            '&::-webkit-scrollbar-thumb': {
              background: 'rgba(168, 85, 247, 0.3)',
              borderRadius: '4px',
              '&:hover': {
                background: 'rgba(168, 85, 247, 0.5)',
              },
            },
          }}
        >
          {error ? (
            <Box sx={{ p: 2, display: 'flex', gap: 1, alignItems: 'flex-start' }}>
              <InfoIcon sx={{ color: 'error.main', fontSize: 20, mt: 0.5 }} />
              <Box>
                <Typography variant="body2" sx={{ fontWeight: 600, mb: 0.5, color: 'error.main' }}>
                  Failed to load search index
                </Typography>
                {error}
              </Box>
            </Box>
          ) : isLoading ? (
            <Box sx={{ p: 2, display: 'flex', gap: 1, alignItems: 'center', justifyContent: 'center' }}>
              <Box
                sx={{
                  width: 16,
                  height: 16,
                  border: '2px solid rgba(168, 85, 247, 0.3)',
                  borderTopColor: '#a855f7',
                  borderRadius: '50%',
                  animation: 'spin 0.6s linear infinite',
                  '@keyframes spin': {
                    '0%': { transform: 'rotate(0deg)' },
                    '100%': { transform: 'rotate(360deg)' },
                  },
                }}
              />
              <Typography variant="body2" sx={{ color: 'rgba(255, 255, 255, 0.6)' }}>
                Loading…
              </Typography>
            </Box>
          ) : showEmpty ? (
            <Box sx={{ p: 2 }}>
              <Typography variant="body2" sx={{ color: 'rgba(255, 255, 255, 0.6)' }}>
                No results found.
              </Typography>
            </Box>
          ) : showResults ? (
            <List sx={{ py: 0.5 }}>
              {results.map((result) => (
                <Box key={result.url}>
                  <Box
                    sx={{
                      px: 2,
                      py: 1,
                      borderBottom: '1px solid rgba(168, 85, 247, 0.1)',
                      mb: 1,
                    }}
                  >
                    <Typography
                      variant="caption"
                      sx={{
                        color: 'rgba(255, 255, 255, 0.5)',
                        textTransform: 'uppercase',
                        fontWeight: 600,
                        fontSize: '11px',
                      }}
                    >
                      {result.meta.title}
                    </Typography>
                  </Box>
                  {result.sub_results.map((subResult) => (
                    <ListItem
                      key={subResult.url}
                      disablePadding
                      sx={{ mx: 1, mb: 0.5 }}
                    >
                      <ListItemButton
                        onClick={() => handleResultClick(subResult)}
                        sx={{
                          px: 2.5,
                          py: 1.5,
                          borderRadius: 1,
                          '&:hover': {
                            bgcolor: 'rgba(168, 85, 247, 0.15)',
                            '& .MuiListItemText-primary': {
                              color: '#c084fc',
                            },
                          },
                          transition: 'all 0.2s',
                        }}
                      >
                        <ListItemText
                          primary={
                            <Typography
                              variant="body2"
                              sx={{
                                fontWeight: 600,
                                color: 'rgba(255, 255, 255, 0.9)',
                                mb: 0.5,
                              }}
                            >
                              {subResult.title}
                            </Typography>
                          }
                          secondary={
                            <Typography
                              variant="body2"
                              sx={{
                                fontSize: '13px',
                                color: 'rgba(255, 255, 255, 0.6)',
                                '& mark': {
                                  bgcolor: 'rgba(168, 85, 247, 0.3)',
                                  color: '#c084fc',
                                  fontWeight: 600,
                                  padding: '0 2px',
                                  borderRadius: '2px',
                                },
                              }}
                              dangerouslySetInnerHTML={{ __html: subResult.excerpt }}
                            />
                          }
                        />
                      </ListItemButton>
                    </ListItem>
                  ))}
                </Box>
              ))}
            </List>
          ) : null}
        </Paper>
      </Popper>
    </>
  );
};
