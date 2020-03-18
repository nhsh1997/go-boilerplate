package servers

import (
	"fmt"
	user_apis "go-boilerplate/src/interface/http/api/user"
	"net/http"
	"github.com/gorilla/mux"
)

type MainRouter struct {
	Router *mux.Router
}

func NewMainRouter(
	userRouter *user_apis.Router,
	) *MainRouter {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/", homePage)

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	userRouter.MakeUserHandler(apiRouter)

	return &MainRouter{
		mainRouter,
	}
}

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "The homepage.")
}
