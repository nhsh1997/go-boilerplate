package main

import (
	types "image-review/src"
	"image-review/src/infra/ioc"
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

