package servers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/auth/auth_router"
	"github.com/nhsh1997/go-boilerplate/src/interface/http/api/user/users_router"
	"net/http"
)

type MainRouter struct {
	Router *mux.Router
}

func NewMainRouter(
	userRouter *users_router.Router,
	authRouter *auth_router.Router,
	) *MainRouter {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/", homePage)

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	userRouter.MakeUserHandler(apiRouter)
	authRouter.MakeUserHandler(apiRouter)

	return &MainRouter{
		mainRouter,
	}
}

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "The homepage.")
}
