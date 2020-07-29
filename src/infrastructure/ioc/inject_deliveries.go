package ioc

import (
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/auth/auth_controller"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/auth/auth_router"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/user/users_controller"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/user/users_router"
)

func inject_deliveries()  {
	//user module
	container.Provide(users_controller.NewUserController)
	container.Provide(users_router.NewUserRouter)
	container.Provide(auth_controller.NewAuthController)
	container.Provide(auth_router.NewAuthRouter)
}