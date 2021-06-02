package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	var DB *sql.DB
	var err error
	if DB == nil {
		DB, err = sql.Open("mysql", "root:@/users")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		DB.SetConnMaxLifetime(time.Minute * 3)
		DB.SetMaxOpenConns(10)
		DB.SetMaxIdleConns(10)
	}

	return DB
}
