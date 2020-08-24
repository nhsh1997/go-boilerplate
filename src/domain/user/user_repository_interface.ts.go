package user_domain

type Repository interface {
	FindByEmail(email string) (*User, error)
	FindByEmailAsSystem(email string) (*User, error)
	FindByAccountId(accountId int64) (*User, error)
}