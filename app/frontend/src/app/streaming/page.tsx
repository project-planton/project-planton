'use client';
import { useState, useRef, useCallback } from 'react';
import {
  Typography,
  Button,
  Box,
  Paper,
  TextField,
  Divider,
} from '@mui/material';
import { StreamingContainer, LogContainer, LogEntry } from '@/app/streaming/styled';
import { useStreamingQuery } from '@/app/streaming/_services';

export default function StreamingPage() {
  const { query } = useStreamingQuery();
  const [isStreaming, setIsStreaming] = useState(false);
  const [logs, setLogs] = useState<Array<{ timestamp: Date; message: string; type: 'info' | 'error' | 'success' }>>([]);
  const [messageCount, setMessageCount] = useState<number>(10);
  const [intervalMs, setIntervalMs] = useState<number>(1000);
  const abortControllerRef = useRef<AbortController | null>(null);

  const addLog = useCallback((message: string, type: 'info' | 'error' | 'success' = 'info') => {
    setLogs((prev) => [
      ...prev,
      {
        timestamp: new Date(),
        message,
        type,
      },
    ]);
    // Auto-scroll to bottom
    setTimeout(() => {
      const logContainer = document.getElementById('log-container');
      if (logContainer) {
        logContainer.scrollTop = logContainer.scrollHeight;
      }
    }, 100);
  }, []);

  const handleStartStream = useCallback(async () => {
    if (!query) {
      addLog('Streaming client not ready', 'error');
      return;
    }

    setIsStreaming(true);
    setLogs([]);
    addLog('Starting stream...', 'info');

    // Create AbortController for cancellation
    const abortController = new AbortController();
    abortControllerRef.current = abortController;

    try {
      // Start streaming with options
      const stream = query.streamData(
        {
          messageCount: messageCount > 0 ? messageCount : undefined,
          intervalMs: intervalMs > 0 ? intervalMs : undefined,
        },
        abortController.signal
      );

      addLog(`Stream started with messageCount: ${messageCount || 'infinite'}, intervalMs: ${intervalMs}ms`, 'success');

      // Iterate over the stream
      for await (const response of stream) {
        if (abortController.signal.aborted) {
          addLog('Stream aborted by user', 'info');
          break;
        }

        // Log the stream response
        const logMessage = `[Sequence: ${response.sequence}] ${response.data} | Status: ${response.status}`;
        addLog(logMessage, 'info');

        // Also log to console
        console.log('Stream Data:', {
          sequence: response.sequence,
          data: response.data,
          timestamp: response.timestamp,
          status: response.status,
        });

        // Check if stream is completed
        if (response.status === 'completed') {
          addLog('Stream completed successfully', 'success');
          setIsStreaming(false);
          break;
        }
      }
    } catch (error: any) {
      if (error.name === 'AbortError') {
        addLog('Stream cancelled by user', 'info');
      } else {
        const errorMessage = error.message || 'Unknown stream error';
        addLog(`Stream error: ${errorMessage}`, 'error');
        console.error('Stream error:', error);
      }
      setIsStreaming(false);
    } finally {
      abortControllerRef.current = null;
    }
  }, [query, messageCount, intervalMs, addLog]);

  const handleStopStream = useCallback(() => {
    if (abortControllerRef.current) {
      abortControllerRef.current.abort();
      abortControllerRef.current = null;
      setIsStreaming(false);
      addLog('Stream stopped by user', 'info');
    }
  }, [addLog]);

  const handleClearLogs = useCallback(() => {
    setLogs([]);
  }, []);

  return (
    <StreamingContainer>
      <Typography variant="h4" gutterBottom>
        Streaming API Test
      </Typography>

      <Paper sx={{ p: 3, mb: 3 }}>
        <Typography variant="h6" gutterBottom>
          Stream Configuration
        </Typography>

        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mb: 2 }}>
          <TextField
            label="Message Count"
            type="number"
            value={messageCount}
            onChange={(e) => setMessageCount(parseInt(e.target.value) || 0)}
            helperText="Number of messages to stream (0 = infinite)"
            disabled={isStreaming}
            inputProps={{ min: 0 }}
          />

          <TextField
            label="Interval (ms)"
            type="number"
            value={intervalMs}
            onChange={(e) => setIntervalMs(parseInt(e.target.value) || 1000)}
            helperText="Interval between messages in milliseconds"
            disabled={isStreaming}
            inputProps={{ min: 100 }}
          />
        </Box>

        <Box sx={{ display: 'flex', gap: 2 }}>
          <Button
            variant="contained"
            color="primary"
            onClick={handleStartStream}
            disabled={isStreaming || !query}
          >
            Start Stream
          </Button>
          <Button
            variant="outlined"
            color="error"
            onClick={handleStopStream}
            disabled={!isStreaming}
          >
            Stop Stream
          </Button>
          <Button variant="outlined" onClick={handleClearLogs} disabled={isStreaming}>
            Clear Logs
          </Button>
        </Box>
      </Paper>

      <Paper sx={{ p: 3 }}>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
          <Typography variant="h6">Stream Logs</Typography>
          <Typography variant="body2" color="text.secondary">
            {logs.length} log entries
          </Typography>
        </Box>

        <Divider sx={{ mb: 2 }} />

        <LogContainer id="log-container">
          {logs.length === 0 ? (
            <Typography variant="body2" color="text.secondary" sx={{ textAlign: 'center', py: 4 }}>
              No logs yet. Start a stream to see logs here.
            </Typography>
          ) : (
            logs.map((log, index) => (
              <LogEntry key={index} type={log.type}>
                <Box sx={{ display: 'flex', gap: 1, alignItems: 'flex-start' }}>
                  <Typography
                    variant="caption"
                    color="text.secondary"
                    sx={{ minWidth: 180, fontFamily: 'monospace' }}
                  >
                    {log.timestamp.toLocaleTimeString()}
                  </Typography>
                  <Typography variant="body2" component="pre" sx={{ margin: 0, whiteSpace: 'pre-wrap', wordBreak: 'break-word' }}>
                    {log.message}
                  </Typography>
                </Box>
              </LogEntry>
            ))
          )}
        </LogContainer>
      </Paper>
    </StreamingContainer>
  );
}

