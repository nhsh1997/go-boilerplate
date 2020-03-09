package apps

import servers "image-review/src/interface/http"

type Application struct {
	Server *servers.Server

}

func NewApplication (server *servers.Server) *Application {
	return &Application{
		Server: server,
	}
}

func (app Application) Start( ){
	app.Server.Start()
}

