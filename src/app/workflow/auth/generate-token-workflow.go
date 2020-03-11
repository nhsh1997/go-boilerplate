package auth

import (
	"fmt"
	types "image-review/src"
	users "image-review/src/domain/user"
)

type GenerateTokenWorkFlow struct {
	jwtHelper types.IJwtHelper
	userRepository users.Repository
}

func NewGenerateTokenWorkFlow (helper types.IJwtHelper, repository users.Repository) *GenerateTokenWorkFlow{
	return &GenerateTokenWorkFlow{
		jwtHelper:      helper,
		userRepository: repository,
	}
}

func (g *GenerateTokenWorkFlow) Execute(credential *users.User) error {
	user, err := g.userRepository.FindByEmail(credential.Email)

	if user != nil {
		return fmt.Errorf("%s already exists", credential.Email)
	}
	if err != nil {
		return err
	}
	return nil
}