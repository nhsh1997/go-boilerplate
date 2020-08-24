package db

import (
	"github.com/nhsh1997/go-boilerplate/ent"
	postgresql_client "github.com/nhsh1997/go-boilerplate/src/infrastructure/db/postgresql"
)

func NewUserModel( postgres *postgresql_client.PostgreClient) *ent.UserClient {
	return postgres.Client.User
}
