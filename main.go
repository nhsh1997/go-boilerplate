package main

import (
	types "github.com/nhsh1997/go-boilerplate/src"
	"github.com/nhsh1997/go-boilerplate/src/infrastructure/ioc"
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

