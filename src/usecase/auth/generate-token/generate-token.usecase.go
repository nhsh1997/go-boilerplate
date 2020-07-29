package generate_token

import (
	jwt_helper "github.com/nhsh1997/go-boilerplate/src/infrastructure/utils/jwt_helper"
)

type IGenerateTokenWorkFlow interface {
	Execute(string) (string, error)
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

func (g *GenerateTokenWorkFlow) Execute(email string) (string, error) {
	/*user, err := g.userRepository.FindByEmail(credential.Email)

	if user != nil {
		return fmt.Errorf("%s already exists", credential.Email)
	}
	if err != nil {
		return err
	}*/
	return email, nil
}