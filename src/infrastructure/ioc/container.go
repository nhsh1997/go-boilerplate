package ioc

import (
	configs "go-boilerplate/config"
	apps "go-boilerplate/src"
	"go-boilerplate/src/infrastructure/db/mongo-client"
	"go-boilerplate/src/infrastructure/utils/jwt_helper"
	servers "go-boilerplate/src/interface/http"
	"go.uber.org/dig"
)

var(
	container = dig.New()
)


func BuildContainer() *dig.Container {

	inject_deliveries()

	container.Provide(configs.NewConfig)
	container.Provide(jwt_helper.NewJwtHelper)
	container.Provide(servers.NewMainRouter)
	container.Provide(servers.NewServer)
	container.Provide(mongo_client.NewMongoDBClient)
	container.Provide(apps.NewApplication)

	return container
}