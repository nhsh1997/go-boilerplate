package apps

import servers "github.com/nhsh1997/go-boilerplate/src/interface/http"

type Application struct {
	Server servers.IServer
}

type IApplication interface {
	Start()
}

func NewApplication (server servers.IServer) IApplication {
	return &Application{
		Server: server,
	}
}

func (app Application) Start( ){
	app.Server.Start()
}

