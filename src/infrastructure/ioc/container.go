package ioc

import (
	configs "go-boilerplate/config"
	apps "go-boilerplate/src"
	dbs "go-boilerplate/src/infrastructure/db"
	"go-boilerplate/src/infrastructure/utils/jwt-helper"
	servers "go-boilerplate/src/interface/http"
	user_apis "go-boilerplate/src/interface/http/api/user"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(configs.NewConfig)
	container.Provide(jwt_helper.NewJwtHelper)
	container.Provide(servers.NewMainRouter)
	container.Provide(servers.NewServer)
	container.Provide(dbs.NewMongoDBClient)
	container.Provide(apps.NewApplication)

	//Apis
	container.Provide(user_apis.NewUserRouter)

	return container
}