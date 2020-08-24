package auth_controller

import (
	"encoding/json"
	"fmt"
	generate_token "github.com/nhsh1997/go-boilerplate/src/usecase/auth/generate-token"
	"net/http"
)

type IAuthController interface {
	GetToken(http.ResponseWriter, *http.Request) (interface{}, error)
	VerifyToken(http.ResponseWriter, *http.Request) (interface{}, error)
}

type AuthController struct {
	generateTokenWorkFlow generate_token.IGenerateTokenWorkFlow
}

func (c *AuthController) VerifyToken(res http.ResponseWriter, req *http.Request) (interface{}, error) {
	return nil, nil
}

func NewAuthController( generateToken generate_token.IGenerateTokenWorkFlow) IAuthController {
	return &AuthController{
		generateTokenWorkFlow: generateToken,
	}
}

func (c *AuthController) GetToken(res http.ResponseWriter, req *http.Request) (interface{}, error) {
	var credential generate_token.GenerateTokenWorkflowInput

	err := json.NewDecoder(req.Body).Decode(&credential)

	if err != nil {
		return nil, err
	}

	result, err := c.generateTokenWorkFlow.Execute(credential)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func homePage2(res http.ResponseWriter, req http.Request) {
	fmt.Print("a")
	fmt.Fprint(res, "The homepage. 2")
}
