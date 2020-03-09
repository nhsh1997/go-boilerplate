package photos

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Router struct {
}

func (r Router) MakePhotoHandler(mainRouter *mux.Router) {
	photoRouter := mainRouter.PathPrefix("/photos").Subrouter()
	photoRouter.HandleFunc("/", homePage1)
	photoRouter.HandleFunc("/upload", homePage2)
}

func NewPhotoRouter() *Router {
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

