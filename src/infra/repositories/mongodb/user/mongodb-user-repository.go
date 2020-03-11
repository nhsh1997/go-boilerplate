package user

import (
	"github.com/juju/mgosession"
	users "image-review/src/domain/user"
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
