package util

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type HttpServer struct {
	ln     net.Listener
	server *http.Server
}

func (s *HttpServer) Serve() error {
	return s.server.Serve(s.ln)
}

func (s *HttpServer) Shutdown(ctx context.Context) {
	_ = s.server.Shutdown(ctx)
	return
}

func NewHttpServer(port int, handler http.Handler) *HttpServer {
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Error("listen port=%d err: %s", port, err.Error())
		return nil
	}

	return &HttpServer{
		ln: ln,
		server: &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", port),
			Handler: handler,
		},
	}
}
