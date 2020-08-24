package user_repository

import (
	"context"
	"github.com/nhsh1997/go-boilerplate/ent"
	"github.com/nhsh1997/go-boilerplate/ent/user"
	user_domain "github.com/nhsh1997/go-boilerplate/src/domain/user"
)

type PostgresUserRepository struct {
	userModel *ent.UserClient
	ctx context.Context
	mapper IPostgresUserMapper
}

func NewPostgresUserRepository(client *ent.UserClient, mapper IPostgresUserMapper) user_domain.Repository  {
	return &PostgresUserRepository{
		userModel: client,
		ctx: context.Background(),
		mapper: mapper,
	}
}

func (p PostgresUserRepository) FindByEmail(email string) (*user_domain.User, error) {
	u, err := p.userModel.Query().Where(user.Email(email)).Only(p.ctx)
	if err != nil {
		return nil, err
	}
	return p.mapper.ToEntity(u), nil
}

func (p PostgresUserRepository) FindByEmailAsSystem(email string) (*user_domain.User, error) {
	panic("implement me")
}

func (p PostgresUserRepository) FindByAccountId(accountId int64) (*user_domain.User, error) {
	panic("implement me")
}
