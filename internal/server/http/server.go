package http

import (
	"context"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/vench/rotator/internal/config"
	"go.uber.org/zap"
)

// Server contains and produce maintance web service.
type Server struct {
	logger *zap.Logger
	conf   *config.App
}

// NewServer create instance of Server.
func NewServer(logger *zap.Logger, conf *config.App) (*Server, error) {
	return &Server{
		logger: logger,
		conf:   conf,
	}, nil
}

func (s *Server) Serve(ctx context.Context) error {
	srv := &fasthttp.Server{
		Handler:            s.router(ctx),
		Name:               s.conf.Name + " http server",
		ReadTimeout:        time.Second,
		WriteTimeout:       time.Second,
		CloseOnShutdown:    true,
		TCPKeepalive:       true,
		TCPKeepalivePeriod: time.Minute,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe(fmt.Sprintf(":%d", s.conf.HTTP.Port))
	}()

	s.logger.Info("HTTP server is running",
		zap.Int("port", s.conf.HTTP.Port),
	)

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(); err != nil {
			return err
		}
		return nil
	case err := <-errCh:
		return err
	}

	return nil
}
