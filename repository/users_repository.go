package repository

import (
	usersdb "github.com/mreines/go-users-api/db"
	usermodel "github.com/mreines/go-users-api/model"
)

// UsersGet - Returns a list of users.
func UsersGet() []usermodel.User {
	rows, err := usersdb.GetDB().Query("select * from users")

	var users []usermodel.User
	if err != nil {
		return users
	}

	for rows.Next() {
		var name string
		var id int64
		err = rows.Scan(&id, &name)
		if err != nil {
			break
		}
		users = append(users, usermodel.NewUser(id, name))
	}
	return users
}

func UsersPost(username string) (usermodel.User, error) {
	result, err := usersdb.GetDB().Exec("insert into users (name) values (?)", username)
	if err == nil {
		id, err := result.LastInsertId()
		if err == nil {
			return usermodel.NewUser(id, username), nil
		}
	}
	return usermodel.NewUser(0, ""), err
}
