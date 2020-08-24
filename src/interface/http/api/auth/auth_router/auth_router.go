package auth_router

import (
	"github.com/gorilla/mux"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/auth/auth_controller"
	"net/http"

	"github.com/nhsh1997/go-boilerplate/src/interface/http/middlewares"
)

type Router struct {
	controller auth_controller.IAuthController
}

func (r Router) MakeAuthHandler(mainRouter *mux.Router) {
	router := mainRouter.PathPrefix("/auth").Subrouter()
	router.Handle("/token", middlewares.Interceptor(r.controller.GetToken)).Methods(http.MethodPost)
	router.Handle("/verifyToken", middlewares.Interceptor(r.controller.VerifyToken))
}

func NewAuthRouter(controller auth_controller.IAuthController) *Router {
	return &Router{
		controller: controller,
	}
}


