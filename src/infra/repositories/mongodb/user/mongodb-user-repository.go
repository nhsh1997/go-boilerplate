package user_mongo_repositories

import (
	"github.com/juju/mgosession"
	users "go-boilerplate/src/domain/user"
)

type repo struct {
	pool *mgosession.Pool
}

//NewMongoRepository create new repository
func NewMongoRepository(p *mgosession.Pool) users.Repository {
	return &repo{
		pool: p,
	}
}


func (r *repo) FindByEmail(email string) (*users.User, error) {
}
