package http

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// HTTPServerOption is a function for setting fields of HTTPServer
type HTTPServerOption func(server *HTTPServer)

// WithAddress returns an option for setting HTTPServer.address
func WithAddress(address string) HTTPServerOption {
	return func(server *HTTPServer) {
		server.address = address
	}
}

// WithPort returns an option for setting HTTPServer.port
func WithPort(port string) HTTPServerOption {
	return func(server *HTTPServer) {
		server.port = port
	}
}

// WithHandler returns an option for setting HTTPServer.handler
func WithHandler(handler http.Handler) HTTPServerOption {
	return func(server *HTTPServer) {
		server.handler = handler
	}
}

// WithTLSConfig returns an option for setting HTTPServer.tls
func WithTLSConfig(config *tls.Config) HTTPServerOption {
	return func(server *HTTPServer) {
		server.tls = config
	}
}

// WithReadTimeout returns an option for setting HTTPServer.readTimeout
func WithReadTimeout(duration time.Duration) HTTPServerOption {
	return func(server *HTTPServer) {
		server.readTimeout = duration
	}
}

// WithReadHeaderTimeout returns an option for setting HTTPServer.readHeaderTimeout
func WithReadHeaderTimeout(duration time.Duration) HTTPServerOption {
	return func(server *HTTPServer) {
		server.readHeaderTimeout = duration
	}
}

// WithWriteTimeout returns an option for setting HTTPServer.writeTimeout
func WithWriteTimeout(duration time.Duration) HTTPServerOption {
	return func(server *HTTPServer) {
		server.writeTimeout = duration
	}
}

// WithMiddlewares returns an option for setting HTTPServer.middlewares
func WithMiddlewares(middlewares ...mux.MiddlewareFunc) HTTPServerOption {
	return func(server *HTTPServer) {
		server.middlewares = middlewares
	}
}
