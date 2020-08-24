package user_repository

import (
	"github.com/nhsh1997/go-boilerplate/ent"
	user_domain "github.com/nhsh1997/go-boilerplate/src/domain/user"
)

type IPostgresUserMapper interface {
	ToEntity(data *ent.User) *user_domain.User
}

type PostgresUserMapper struct {
}

func NewPostgreUserMapper() IPostgresUserMapper {
	return PostgresUserMapper{}
}

func (p PostgresUserMapper) ToEntity(data *ent.User) *user_domain.User {
	return &user_domain.User{
		ID:           data.ID,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		IsSuperAdmin: data.IsSuperAdmin,
		Status:       data.Status,
		UpdatedAt:    data.UpdatedAt,
		CreatedAt:    data.CreatedAt,
	}
}