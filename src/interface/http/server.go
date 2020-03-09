package servers

import (
	"fmt"
	Configuration "image-review/config"
	"net/http"
	"time"
)

type Server struct {
	//logger log.Logger
	Port int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	HttpServer *http.Server
	MainRouter *MainRouter
}

func NewServer (config *Configuration.Configuration, router *MainRouter) *Server {
	return &Server{
		Port: config.Web.Port,
		ReadTimeout: config.Web.ReadTimeout,
		WriteTimeout: config.Web.WriteTimeout,
		MainRouter: router,
	}
}

func (server *Server) Start() {
	conn := fmt.Sprintf(":%d", server.Port)
	fmt.Print(conn)

	server.HttpServer = &http.Server{
		Addr:           conn,
		Handler:        server.MainRouter.Router,
		ReadTimeout:    server.ReadTimeout * time.Second,
		WriteTimeout:   server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.HttpServer.ListenAndServe()
}

func (server *Server) Stop() {
	server.HttpServer.Close()
}