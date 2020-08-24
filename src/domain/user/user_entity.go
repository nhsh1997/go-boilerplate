package user_domain

import "time"

type User struct {
	ID  int  `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsSuperAdmin bool `json:"is_super_admin"`
	Status int16 `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(id int, firstName, lastName, email, password string, isSuperAdmin bool, status int16) *User{
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		Status: status,
		IsSuperAdmin: isSuperAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
