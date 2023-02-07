package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 10 * time.Second
	_defaultWriteTimeout    = 10 * time.Second
	_defaultAddr            = ":8000"
	_defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr:              _defaultAddr,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       _defaultReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      _defaultWriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s := &Server{
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start() {
	go func() {
		s.server.ListenAndServe()
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
