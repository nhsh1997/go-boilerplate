package auth_controller

import (
	"encoding/json"
	"fmt"
	rest_errors "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/error"
	generate_token "github.com/nhsh1997/go-boilerplate/src/usecase/auth/generate-token"
	"net/http"
)

type IAuthController interface {
	GetToken(http.ResponseWriter, *http.Request) error
	VerifyToken(http.ResponseWriter, *http.Request) error
}

type AuthController struct {
	generateTokenWorkFlow generate_token.IGenerateTokenWorkFlow
}

func (c *AuthController) VerifyToken(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func NewAuthController( generateToken generate_token.IGenerateTokenWorkFlow) IAuthController {
	return &AuthController{
		generateTokenWorkFlow: generateToken,
	}
}

func handleError(err error, res http.ResponseWriter){
	if err != nil {
		switch err.(type) {
		case *rest_errors.RestError:
			err := err.(*rest_errors.RestError)
			http.Error(res, err.Message, err.Status)
		default:
			err := rest_errors.NewInternalServerError(`The server failed to handle this request`, err)
			http.Error(res, err.Message, err.Status)
		}
	}
}

func (c *AuthController) GetToken(res http.ResponseWriter, req *http.Request) error {
	var credential generate_token.GenerateTokenWorkflowInput

	err := json.NewDecoder(req.Body).Decode(&credential)

	if err != nil {
		return err
	}

	result, err := c.generateTokenWorkFlow.Execute(credential)

	if err != nil {
		return err
	}

	json.NewEncoder(res).Encode(result)

	return nil
}

func homePage2(res http.ResponseWriter, req http.Request) {
	fmt.Print("a")
	fmt.Fprint(res, "The homepage. 2")
}
