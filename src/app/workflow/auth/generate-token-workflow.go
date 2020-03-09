package auth

import (
	users "image-review/src/domain/user"
	"image-review/src/infra/utils"
)

type GenerateTokenWorkFlow struct {
	jwtHelper *utils.JwtHelper
	userRepository *users.Repository
}

func NewGenerateTokenWorkFlow (helper *utils.JwtHelper, repository *users.Repository) *GenerateTokenWorkFlow{
	return &GenerateTokenWorkFlow{
		jwtHelper:      helper,
		userRepository: repository,
	}
}

func (g *GenerateTokenWorkFlow) Execute() {
}