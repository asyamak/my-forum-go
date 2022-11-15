package httpserver

import (
	"context"
	"fmt"
	"forum/config"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Srv             *http.Server
	notify          chan error
	shutdownTimeOut time.Duration
	// config     *config.Config
}

func NewServer(conf *config.Config, router *http.ServeMux) *Server {
	server := &Server{
		Srv: &http.Server{
			Addr:           ":" + conf.Port,
			Handler:        router,
			ReadTimeout:    conf.ReadTimeout,
			WriteTimeout:   conf.WriteTimeout,
			MaxHeaderBytes: conf.MaxHeaderBytes,
		},
		notify:          make(chan error, 1),
		shutdownTimeOut: conf.ShutdownTimeOut,
	}

	server.start()
	return server
}

func (s *Server) start() {
	log.Printf("server has been initiated on http://localhost%v\n", s.Srv.Addr)
	go func() {
		s.notify <- s.Srv.ListenAndServe()
		close(s.notify)
		fmt.Println("notify chan")
	}()
}

// figure out how it catch signal
func (s *Server) Notify() <-chan error {
	return s.notify
}

// gracefully shutdowns server
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeOut)
	defer cancel()
	defer fmt.Println("gracefully shutdowning server")
	return s.Srv.Shutdown(ctx)
}
