package main

import (
	types "go-boilerplate/src"
	"go-boilerplate/src/infra/ioc"
)


func main() {
	container := ioc.BuildContainer()

	err := container.Invoke(func(app types.IApplication) {
		app.Start()
	})

	if err != nil {
		panic(err)
	}
}

