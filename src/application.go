package apps

type Application struct {
	Server IServer
}

func NewApplication (server IServer) IApplication {
	return &Application{
		Server: server,
	}
}

func (app Application) Start( ){
	app.Server.Start()
}

