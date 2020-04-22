package servers

import (
	"fmt"
	configs "github.com/nhsh1997/go-boilerplate/config"
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

type IServer interface  {
	Start()
	Stop()
}

func NewServer (config *configs.Configuration, router *MainRouter) IServer {
	return &Server{
		Port: config.Web.Port,
		ReadTimeout: config.Web.ReadTimeout,
		WriteTimeout: config.Web.WriteTimeout,
		MainRouter: router,
	}
}

func (server *Server) Start() {
	conn := fmt.Sprintf(":%d", server.Port)
	fmt.Print("Server run on port" + conn)

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