package users

type Repository interface {
	FindByEmail(email string) (*User, error)
}
