package ioc

import (
	"go-boilerplate/src/interface/http/api/user/users_controller"
	"go-boilerplate/src/interface/http/api/user/users_router"
)

func inject_deliveries()  {
	//user module
	container.Provide(users_controller.NewUserController)
	container.Provide(users_router.NewUserRouter)
}