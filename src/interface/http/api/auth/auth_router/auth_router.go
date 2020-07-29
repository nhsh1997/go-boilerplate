package auth_router

import (
	"github.com/gorilla/mux"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/auth/auth_controller"
	"net/http"
)

type Router struct {
	controller auth_controller.IAuthController
}

func (r Router) MakeUserHandler(mainRouter *mux.Router) {
	router := mainRouter.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/token", r.controller.GetToken).Methods(http.MethodPost)
	router.HandleFunc("/verifyToken", r.controller.VerifyToken)
}

func NewAuthRouter(controller auth_controller.IAuthController) *Router {
	return &Router{
		controller: controller,
	}
}


