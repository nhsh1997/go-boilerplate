package users

import "time"

type User struct {
	ID  int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status int16 `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(id int64, firstName, lastName, email, password string, status int16) *User{
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		Status: status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
