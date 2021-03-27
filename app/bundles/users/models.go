package users

import "time"

type User struct {
	Id         int       `json:"id"`
	Dni        string    `json:"dni"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	DateJoined time.Time `json:"dateJoined"`
}

func NewUser(dni string, name string, lastName string,
	email string, password string) *User {
	return &User{
		Id:         0,
		Dni:        dni,
		Name:       name,
		LastName:   lastName,
		Email:      email,
		Password:   password,
		DateJoined: time.Now(),
	}
}
