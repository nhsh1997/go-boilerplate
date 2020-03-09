package servers

import (
	"fmt"
	photos "image-review/src/interface/http/api/photo"
	"net/http"
	"github.com/gorilla/mux"
)

type MainRouter struct {
	Router *mux.Router
}

func NewMainRouter(photoRouter *photos.Router) *MainRouter {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/", homePage)

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	photoRouter.MakePhotoHandler(apiRouter)

	return &MainRouter{
		mainRouter,
	}
}

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprint(res, "The homepage.")
}
