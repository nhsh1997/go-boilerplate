package apps

import (
	types "image-review/src"
)

type Application struct {
	Server types.IServer

}

func NewApplication (server types.IServer) *Application {
	return &Application{
		Server: server,
	}
}

func (app Application) Start( ){
	app.Server.Start()
}

