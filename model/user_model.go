package model

type User struct {
	ID   uint
	Name string
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}
