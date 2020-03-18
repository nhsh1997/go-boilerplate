package generate_token

import (
	"fmt"
	user_domain "go-boilerplate/src/domain/user"
	jwt_helper "go-boilerplate/src/infrastructure/utils/jwt-helper"
)

type GenerateTokenWorkFlow struct {
	jwtHelper jwt_helper.IJwtHelper
	userRepository user_domain.Repository
}

func NewGenerateTokenWorkFlow (helper jwt_helper.IJwtHelper, repository user_domain.Repository) *GenerateTokenWorkFlow{
	return &GenerateTokenWorkFlow{
		jwtHelper:      helper,
		userRepository: repository,
	}
}

func (g *GenerateTokenWorkFlow) Execute(credential *user_domain.User) error {
	user, err := g.userRepository.FindByEmail(credential.Email)

	if user != nil {
		return fmt.Errorf("%s already exists", credential.Email)
	}
	if err != nil {
		return err
	}
	return nil
}