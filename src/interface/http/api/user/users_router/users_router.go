package users_router

import (
	"github.com/gorilla/mux"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/user/users_controller"
)

type Router struct {
	controller users_controller.IUserController
}

func (r Router) MakeUserHandler(mainRouter *mux.Router) {
	router := mainRouter.PathPrefix("/users").Subrouter()
	router.HandleFunc("/", r.controller.HomePage1)
}

func NewUserRouter(controller users_controller.IUserController) *Router {
	return &Router{
		controller: controller,
	}
}


