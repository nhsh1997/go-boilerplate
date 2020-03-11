package ioc

import (
	"go.uber.org/dig"
	configs "image-review/config"
	apps "image-review/src/app"
	dbs "image-review/src/infra/db"
	"image-review/src/infra/utils"
	servers "image-review/src/interface/http"
	Photo "image-review/src/interface/http/api/photo"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(configs.NewConfig)
	container.Provide(utils.NewJwtHelper)
	container.Provide(Photo.NewPhotoRouter)
	container.Provide(servers.NewMainRouter)
	container.Provide(servers.NewServer)
	container.Provide(dbs.NewMongoDBClient)
	container.Provide(apps.NewApplication)

	return container
}