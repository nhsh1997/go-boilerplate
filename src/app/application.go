package apps

import (
	types "image-review/src"
)

type Application struct {
	Server types.IServer

}

func NewApplication (server types.IServer) types.IApplication {
	return &Application{
		Server: server,
	}
}

func (app Application) Start( ){
	app.Server.Start()
}

