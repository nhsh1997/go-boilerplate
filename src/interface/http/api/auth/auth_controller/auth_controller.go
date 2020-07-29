package auth_controller

import (
	"encoding/json"
	"fmt"
	generate_token "github.com/nhsh1997/go-boilerplate/src/usecase/auth/generate-token"
	"net/http"
)

type IAuthController interface {
	GetToken(http.ResponseWriter, *http.Request)
	VerifyToken(http.ResponseWriter, *http.Request)
}

type AuthController struct {
	generateTokenWorkFlow generate_token.IGenerateTokenWorkFlow
}

func (c *AuthController) VerifyToken(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "VerifyToken")
}

func NewAuthController( generateToken generate_token.IGenerateTokenWorkFlow) IAuthController  {
	return &AuthController{
		generateTokenWorkFlow: generateToken,
	}
}

func (c *AuthController) GetToken(res http.ResponseWriter, req *http.Request) {
	var p struct{
		Email string `json:"email"`
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(p)
	result, err := c.generateTokenWorkFlow.Execute(p.Email)
	if err != nil {
		fmt.Fprint(res, "Error")
	}

	fmt.Fprint(res, result)
}

func homePage2(res http.ResponseWriter, req http.Request) {
	fmt.Print("a")
	fmt.Fprint(res, "The homepage. 2")
}
