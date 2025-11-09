'use client';

import React, { useState } from 'react';
import { TextField, InputAdornment } from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';

export const SearchBar: React.FC = () => {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(event.target.value);
    // TODO: Implement search functionality
  };

  return (
    <TextField
      size="small"
      placeholder="Search documentation..."
      value={searchQuery}
      onChange={handleSearch}
      className="w-64"
      sx={{
        '& .MuiOutlinedInput-root': {
          color: 'white',
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
      }}
    />
  );
};

