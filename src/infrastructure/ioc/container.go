package ioc

import (
	configs "github.com/nhsh1997/go-boilerplate/config"
	apps "github.com/nhsh1997/go-boilerplate/src"
	db "github.com/nhsh1997/go-boilerplate/src/infrastructure/db/models"
	"github.com/nhsh1997/go-boilerplate/src/infrastructure/db/mongo-client"
	postgresql_client "github.com/nhsh1997/go-boilerplate/src/infrastructure/db/postgresql"
	"github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/jwt_helper"
	servers "github.com/nhsh1997/go-boilerplate/src/interface/http"
	"go.uber.org/dig"
)

var(
	container = dig.New()
)


func BuildContainer() *dig.Container {
	container.Provide(configs.NewConfig)
	container.Provide(jwt_helper.NewJwtHelper)
	container.Provide(servers.NewMainRouter)
	container.Provide(servers.NewServer)
	container.Provide(mongo_client.NewMongoDBClient)
	container.Provide(postgresql_client.NewPostgresClient)
	container.Provide(db.NewUserModel)
	container.Provide(apps.NewApplication)
	inject_deliveries()
	inject_repositories()
	inject_workflows()

	return container
}