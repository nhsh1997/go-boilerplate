package main

import (
	Application "image-review/src/app"
	"image-review/src/infra/ioc"
)


func main() {
	container := ioc.BuildContainer()

	err := container.Invoke(func(app *Application.Application) {
		app.Start()
	})

	if err != nil {
		panic(err)
	}
}

