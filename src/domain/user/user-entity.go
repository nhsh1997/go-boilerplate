package users

type User struct {
	ID  int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName string
	Email string
	Password string
}

func NewUser(id int64, firstName, lastName, email, password string) *User{
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
	}
}
