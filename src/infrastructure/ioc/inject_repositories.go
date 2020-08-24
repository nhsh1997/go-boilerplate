package ioc

import (
	user_repository "github.com/nhsh1997/go-boilerplate/src/infrastructure/repositories/postgres/user"
)

func inject_repositories()  {
	//user module
	container.Provide(user_repository.NewPostgreUserMapper)
	container.Provide(user_repository.NewPostgresUserRepository)
}
