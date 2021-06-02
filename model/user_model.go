package model

type User struct {
	Id   int64
	Name string
}

func NewUser(id int64, name string) User {
	return User{
		Id:   id,
		Name: name,
	}
}
