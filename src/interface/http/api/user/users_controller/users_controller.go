package users_controller

import (
	"fmt"
	"net/http"
)

type IUserController interface {
	HomePage1(http.ResponseWriter, *http.Request)
}

type UserController struct {

}

func NewUserController() IUserController  {
	return &UserController{}
}

func (c *UserController) HomePage1(res http.ResponseWriter, req *http.Request) {
	fmt.Print("b")

	fmt.Fprint(res, "The homepage. 1")
}

func homePage2(res http.ResponseWriter, req http.Request) {
	fmt.Print("a")
	fmt.Fprint(res, "The homepage. 2")
}
