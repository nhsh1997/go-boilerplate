package ioc

import generate_token "github.com/nhsh1997/go-boilerplate/src/usecase/auth/generate-token"


func inject_workflows()  {
	//user module
	container.Provide(generate_token.NewGenerateTokenWorkFlow)
}