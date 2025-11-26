package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-planton/project-planton/backend/internal/database"
	"github.com/project-planton/project-planton/backend/internal/service"
	"github.com/sirupsen/logrus"

	backendv1connect "github.com/project-planton/project-planton/backend/apis/gen/go/proto/backendv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server wraps the HTTP server and dependencies.
type Server struct {
	httpServer *http.Server
	mongo      *database.MongoDB
}

// Config holds server configuration.
type Config struct {
	Port    string
	MongoDB *database.MongoDB
}

// corsMiddleware wraps an HTTP handler with CORS headers.
// Supports Connect RPC and gRPC-Web protocols.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		origin := r.Header.Get("Origin")
		if origin != "" {
			// Allow requests from localhost:3000 (web app) and other common dev origins
			allowedOrigins := []string{
				"http://localhost:3000",
				"http://127.0.0.1:3000",
				"http://localhost:3001",
				"http://127.0.0.1:3001",
			}

			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					allowed = true
					break
				}
			}

			if allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				// Connect RPC and gRPC-Web specific headers
				w.Header().Set("Access-Control-Allow-Headers",
					"Content-Type, Authorization, X-Requested-With, "+
						"grpc-timeout, keep-alive, "+
						"x-accept-content-transfer-encoding, x-accept-response-streaming, "+
						"x-grpc-web, x-user-agent")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Max-Age", "3600")
				w.Header().Set("Access-Control-Expose-Headers",
					"grpc-status, grpc-message, grpc-status-details-bin")
			}
		}

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// NewServer creates a new server instance.
func NewServer(cfg *Config) *Server {
	deploymentComponentRepo := database.NewDeploymentComponentRepository(cfg.MongoDB)
	deploymentComponentService := service.NewDeploymentComponentService(deploymentComponentRepo)

	mux := http.NewServeMux()

	// Register the DeploymentComponentService
	path, handler := backendv1connect.NewDeploymentComponentServiceHandler(deploymentComponentService)
	// Wrap handler with CORS middleware
	mux.Handle(path, corsMiddleware(handler))

	// Create HTTP server with h2c (HTTP/2 Cleartext) for gRPC
	httpServer := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	return &Server{
		httpServer: httpServer,
		mongo:      cfg.MongoDB,
	}
}

// Start starts the server and blocks until it stops.
func (s *Server) Start() error {
	logrus.WithField("addr", s.httpServer.Addr).Info("Starting gRPC server")
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Shutting down server")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	if err := s.mongo.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB: %w", err)
	}
	return nil
}
