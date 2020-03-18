package user_domain

type Repository interface {
	FindByEmail(email string) (*User, error)
}
