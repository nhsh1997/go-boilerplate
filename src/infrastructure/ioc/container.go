package ioc

import (
	configs "github.com/nhsh1997/go-boilerplate/config"
	apps "github.com/nhsh1997/go-boilerplate/src"
	"github.com/nhsh1997/go-boilerplate/src/infrastructure/db/mongo-client"
	"github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/jwt_helper"
	servers "github.com/nhsh1997/go-boilerplate/src/interface/http"
	"go.uber.org/dig"
)

var(
	container = dig.New()
)


func BuildContainer() *dig.Container {

	inject_deliveries()
	inject_workflows()

	container.Provide(configs.NewConfig)
	container.Provide(jwt_helper.NewJwtHelper)
	container.Provide(servers.NewMainRouter)
	container.Provide(servers.NewServer)
	container.Provide(mongo_client.NewMongoDBClient)
	container.Provide(apps.NewApplication)

	return container
}