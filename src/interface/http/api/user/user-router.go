package user_apis

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Router struct {
}

func (r Router) MakeUserHandler(mainRouter *mux.Router) {
	router := mainRouter.PathPrefix("/users").Subrouter()
	router.HandleFunc("/", homePage1)
}

func NewUserRouter() *Router {
	return &Router{
	}
}

func homePage1(res http.ResponseWriter, req *http.Request) {
	fmt.Print("b")

	fmt.Fprint(res, "The homepage. 1")
}

func homePage2(res http.ResponseWriter, req *http.Request) {
	fmt.Print("a")
	fmt.Fprint(res, "The homepage. 2")
}

