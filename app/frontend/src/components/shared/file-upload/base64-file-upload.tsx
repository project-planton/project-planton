'use client';

import React, { FC, useCallback, useRef, useState, useEffect } from 'react';
import { UseFormSetValue } from 'react-hook-form';
import { Button, Stack, Typography, FormHelperText, IconButton } from '@mui/material';
import { CloudUpload, Clear, Download } from '@mui/icons-material';
import { readFileAsBase64 } from '@/lib/utils';

interface UseBase64FileUploadProps {
  setValue: UseFormSetValue<any>;
  path: string;
  maxSizeBytes?: number;
  acceptedExtensions?: string[];
  onError?: (error: string) => void;
}

export const useBase64FileUpload = ({
  setValue,
  path,
  maxSizeBytes,
  acceptedExtensions,
  onError,
}: UseBase64FileUploadProps) => {
  const inputFileRef = useRef<HTMLInputElement>(null);
  const [selectedFile, setSelectedFile] = useState<string>('');
  const [error, setError] = useState<string>('');

  const handleFileChange = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      if (event.target.files && event.target.files.length > 0) {
        const file: File = event.target.files[0];

        // Check file size if maxSizeBytes is provided
        if (maxSizeBytes && file.size > maxSizeBytes) {
          const errorMsg = `File size exceeds the maximum allowed size of ${(maxSizeBytes / 1024).toFixed(0)} KB`;
          setError(errorMsg);
          if (onError) onError(errorMsg);
          return;
        }

        // Check file extension if acceptedExtensions is provided
        if (acceptedExtensions && acceptedExtensions.length > 0) {
          const extension = `.${file.name.split('.').pop()?.toLowerCase()}`;
          if (!acceptedExtensions.includes(extension)) {
            const errorMsg = `File type not allowed. Accepted types: ${acceptedExtensions.join(', ')}`;
            setError(errorMsg);
            if (onError) onError(errorMsg);
            return;
          }
        }

        // Clear any previous errors and set filename immediately
        setError('');
        setSelectedFile(file.name);

        // Read and encode the file
        readFileAsBase64(file)
          .then((base64Content) => {
            setValue(path, base64Content);
          })
          .catch((err) => {
            const errorMsg = 'Failed to read file';
            setError(errorMsg);
            setSelectedFile(''); // Clear filename on error
            if (onError) onError(errorMsg);
            console.error('File reading error:', err);
          });
      }
    },
    [setValue, path, maxSizeBytes, acceptedExtensions, onError]
  );

  const clearFile = useCallback(() => {
    setSelectedFile('');
    setError('');
    setValue(path, '');
    if (inputFileRef.current) {
      inputFileRef.current.value = '';
    }
  }, [setValue, path]);

  const triggerFileClick = useCallback(() => {
    if (inputFileRef.current) {
      inputFileRef.current.click();
    }
  }, []);

  return {
    selectedFile,
    clearFile,
    error,
    triggerFileClick,
    inputFileRef,
    handleFileChange,
    acceptedExtensions,
  };
};

interface FileUploadWithClearProps {
  label?: string;
  buttonText?: string;
  maxSizeBytes: number;
  setValue: UseFormSetValue<any>;
  path: string;
  watch?: (path: string) => any;
  helpText?: string;
  disabled?: boolean;
  downloadFileName?: string;
}

export const FileUploadWithClear: FC<FileUploadWithClearProps> = ({
  label,
  buttonText = 'Upload File',
  maxSizeBytes,
  setValue,
  path,
  watch,
  helpText,
  disabled,
  downloadFileName,
}) => {
  const {
    selectedFile,
    clearFile,
    error,
    triggerFileClick,
    inputFileRef,
    handleFileChange,
    acceptedExtensions,
  } = useBase64FileUpload({
    setValue,
    path,
    maxSizeBytes,
  });

  // Watch the form value to detect existing base64 content (for view/edit mode)
  const currentValue = watch ? watch(path) : undefined;
  const [hasExistingValue, setHasExistingValue] = useState(false);

  useEffect(() => {
    if (currentValue && typeof currentValue === 'string' && currentValue.length > 0) {
      setHasExistingValue(true);
    } else {
      setHasExistingValue(false);
    }
  }, [currentValue]);

  const handleDownload = useCallback(() => {
    if (!currentValue) return;

    try {
      // Backend returns decoded JSON string, so we can use it directly
      let jsonContent = '';
      try {
        // Try to parse and format as JSON
        const jsonObj = JSON.parse(currentValue);
        jsonContent = JSON.stringify(jsonObj, null, 2);
      } catch {
        // If not JSON, use value as-is (shouldn't happen for valid credentials)
        jsonContent = currentValue;
      }

      // Create blob and trigger download
      const blob = new Blob([jsonContent], { type: 'application/json' });
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `${downloadFileName || 'service-account-key'}.json`;
      a.click();
      URL.revokeObjectURL(url);
    } catch (error) {
      console.error('Failed to download file:', error);
    }
  }, [currentValue, downloadFileName]);

  return (
    <Stack gap={1}>
      {label && (
        <Typography variant="subtitle2" color="text.secondary">
          {label}
        </Typography>
      )}
      <Stack direction="row" gap={1} alignItems="center">
        {selectedFile ? (
          <Stack direction="row" gap={1} alignItems="center">
            <Typography variant="body2">{selectedFile}</Typography>
            <IconButton size="small" onClick={clearFile} disabled={disabled}>
              <Clear fontSize="small" />
            </IconButton>
          </Stack>
        ) : hasExistingValue && disabled ? (
          // View mode: show download button for existing credential
          <Stack direction="row" gap={1} alignItems="center">
            <Typography variant="body2" color="text.secondary">
              Service account key file uploaded
            </Typography>
            <Button
              variant="outlined"
              size="small"
              startIcon={<Download />}
              onClick={handleDownload}
              sx={{ width: 'fit-content' }}
            >
              Download Key File
            </Button>
          </Stack>
        ) : (
          // Create/Edit mode: show upload button
          <>
            <Button
              variant="contained"
              color="secondary"
              size="medium"
              startIcon={<CloudUpload />}
              onClick={triggerFileClick}
              disabled={disabled}
              sx={{ width: 'fit-content' }}
            >
              {buttonText}
            </Button>
            <input
              ref={inputFileRef}
              type="file"
              onChange={handleFileChange}
              accept={acceptedExtensions?.join(',') || '*'}
              style={{ display: 'none' }}
              disabled={disabled}
            />
            <Typography variant="body2" color="text.secondary">
              No File Selected
            </Typography>
          </>
        )}
      </Stack>
      {error && <FormHelperText error>{error}</FormHelperText>}
      {helpText && !error && <FormHelperText>{helpText}</FormHelperText>}
    </Stack>
  );
};
