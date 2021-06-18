package repository

import (
	usersdb "github.com/mreines/go-users-api/db"
	usermodel "github.com/mreines/go-users-api/model"
)

// UsersGet - Returns a list of users.
func UserGet(id uint) usermodel.User {
	var user usermodel.User
	user.ID = id
	usersdb.GetDB().First(&user)

	return user
}

// UsersGet - Returns a list of users.
func UsersGet() []usermodel.User {
	var users []usermodel.User
	usersdb.GetDB().Find(&users)

	return users
}

func UsersPost(username string) usermodel.User {
	user := usermodel.NewUser(username)
	usersdb.GetDB().Create(&user)

	return user
}

func UserPatch(id uint, user *usermodel.User) *usermodel.User {
	original := usermodel.NewUser("")
	usersdb.GetDB().First(&original)
	if original.Name != user.Name {
		user.ID = id
		usersdb.GetDB().Save(&user)
	}
	return user
}
