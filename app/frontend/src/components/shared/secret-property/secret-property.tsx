'use client';
import React, { useCallback, useContext, useEffect, useState } from 'react';
import { SvgIcon, Typography } from '@mui/material';
import { Download, VisibilityOffOutlined, VisibilityOutlined } from '@mui/icons-material';
import { AppContext } from '@/contexts';
import { isValidJSON, safeParseJson } from '@/lib';
import { StyledButton } from '@/components/shared/secret-property/styled';
import { FlexCenterRow } from '@/components/shared/resource-header/styled';
import { SecretModal } from '@/components/shared/secret-property/secret-modal';

interface IProps {
  property?: string;
  value?: string;
  message?: string;
  inlineVisibilityLength?: number;
  /**
   * secret value can come from an API call as well. Hence, the promise return type.
   */
  getSecretValue: (value: string) => Promise<string>;
  showInModal?: boolean;
  styledContent?: boolean;
  enableDownload?: boolean;
  downloadFileName?: string;
}

export function SecretProperty({
  property,
  value,
  getSecretValue,
  message,
  inlineVisibilityLength = 50,
  showInModal = false,
  styledContent = false,
  enableDownload = false,
  downloadFileName,
}: IProps) {
  const { openSnackbar } = useContext(AppContext);
  const [showSecret, setShowSecret] = useState(false);
  const [openModal, setOpenModal] = useState(false);
  const [secretValue, setSecretValue] = useState<string | null>(null);

  useEffect(() => {
    if (openModal) return;
    setSecretValue(null);
  }, [openModal]);

  const toggleSecret = async () => {
    // getSecretValue may or may not use the value parameter
    // In some cases, it ignores the parameter and fetches from context
    let secret = await getSecretValue(value || '');
    if (secret) {
      if (isValidJSON(secret)) secret = JSON.stringify(safeParseJson(secret), null, 2);
      setSecretValue(secret);
    }

    if ((secret?.length && secret.length > inlineVisibilityLength) || showInModal) {
      setOpenModal(true);
    }
  };

  useEffect(() => {
    if (showSecret) toggleSecret();
  }, [showSecret]);

  const onModalClose = () => {
    setOpenModal(false);
    setSecretValue(null);
    setShowSecret(false);
  };

  const handleDownload = useCallback(async () => {
    try {
      // Fetch secret value
      const secretValue = await getSecretValue(value || '');

      let jsonContent = '';

      // Try to parse directly as JSON first (in case it's already decoded)
      try {
        const jsonObj = JSON.parse(secretValue);
        jsonContent = JSON.stringify(jsonObj, null, 2);
      } catch {
        // If direct parsing fails, try to decode as base64 first
        try {
          const decoded = atob(secretValue);
          const jsonObj = JSON.parse(decoded);
          jsonContent = JSON.stringify(jsonObj, null, 2);
        } catch {
          // If both fail, throw error to outer catch
          throw new Error('Unable to parse secret value');
        }
      }

      // Create blob and download
      const blob = new Blob([jsonContent], { type: 'application/json' });
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `${downloadFileName || 'service-account-key'}.json`;
      a.click();
      URL.revokeObjectURL(url);

      openSnackbar('File downloaded successfully', 'success');
    } catch {
      openSnackbar('Failed to download file', 'error');
    }
  }, [getSecretValue, value, downloadFileName, openSnackbar]);

  return (
    <>
      <FlexCenterRow>
        <Typography variant="subtitle2">
          {showSecret && !openModal ? secretValue : <Typography marginTop={1}>{value}</Typography>}
        </Typography>
        <StyledButton variant="text">
          {openModal || showSecret ? (
            <SvgIcon
              component={VisibilityOffOutlined}
              color="secondary"
              sx={{ height: 12, width: 12 }}
              fontSize="small"
              onClick={() => setShowSecret(false)}
            />
          ) : (
            <SvgIcon
              component={VisibilityOutlined}
              color="secondary"
              sx={{ height: 12, width: 12 }}
              fontSize="small"
              onClick={() => setShowSecret(true)}
            />
          )}
        </StyledButton>
        {enableDownload && (
          <StyledButton variant="text">
            <SvgIcon
              component={Download}
              color="secondary"
              sx={{ height: 12, width: 12 }}
              fontSize="small"
              onClick={handleDownload}
            />
          </StyledButton>
        )}
      </FlexCenterRow>
      <SecretModal
        open={openModal}
        onClose={onModalClose}
        title={
          !!property && (
            <Typography fontSize={13} fontWeight={600}>
              {property}
            </Typography>
          )
        }
        secretValue={secretValue || ''}
        message={message}
        styledContent={styledContent}
      />
    </>
  );
}
