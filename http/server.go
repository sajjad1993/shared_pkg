package http

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	nethttp "net/http"
	"time"
)

const (
	_defaultAddress = "0.0.0.0"
	_defaultPort    = "8080"
)

type HTTPServer struct {
	server            *nethttp.Server
	handler           nethttp.Handler
	middlewares       []mux.MiddlewareFunc
	address           string
	port              string
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	tls               *tls.Config
}

// Serve starts serving
func (s *HTTPServer) Serve() error {
	err := s.server.ListenAndServe()
	return err
}

// ServeTLS starts serving HTTPS
func (s *HTTPServer) ServeTLS(certFile, keyFile string) error {
	return s.server.ListenAndServeTLS(certFile, keyFile)
}

// Shutdown shuts the server down
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// NewHTTPServer returns a server specified by options
func NewHTTPServer(options ...HTTPServerOption) Server {
	s := &HTTPServer{
		address: _defaultAddress,
		port:    _defaultPort,
	}

	for _, o := range options {
		o(s)
	}
	s.server = &nethttp.Server{
		Addr:              s.address,
		Handler:           s.handler,
		TLSConfig:         s.tls,
		ReadTimeout:       s.readTimeout,
		WriteTimeout:      s.writeTimeout,
		ReadHeaderTimeout: s.readHeaderTimeout,
	}
	return s
}
