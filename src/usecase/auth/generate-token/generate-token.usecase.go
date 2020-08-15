package generate_token

import (
	rest_errors "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/error"
	jwt_helper "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/jwt_helper"
	"time"
)

type IGenerateTokenWorkFlow interface {
	Execute(GenerateTokenWorkflowInput) (*GenerateTokenWorkflowOutput, error)
}

type GenerateTokenWorkFlow struct {
	jwtHelper jwt_helper.IJwtHelper/*
	userRepository user_domain.Repository*/
}

func NewGenerateTokenWorkFlow (helper jwt_helper.IJwtHelper/*, repository user_domain.Repository*/) IGenerateTokenWorkFlow {
	return &GenerateTokenWorkFlow{
		jwtHelper:      helper,
		/*userRepository: repository,*/
	}
}

func (g *GenerateTokenWorkFlow) Execute(credential GenerateTokenWorkflowInput) (*GenerateTokenWorkflowOutput, error) {
	/*user, err := g.userRepository.FindByEmail(credential.Email)

	if user != nil {
		return fmt.Errorf("%s already exists", credential.Email)
	}
	if err != nil {
		return err
	}*/

	if credential.Email != "nhsh1997@gmail.com" {
		return nil, rest_errors.NewBadRequestError("Email is wrong!")
	}
	return &GenerateTokenWorkflowOutput{
		Token:     "Hello",
		ExpiresIn: time.Time{},
	}, nil
}