package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	util "github.com/mreines/go-users-api/util"
)

func GetDB() *gorm.DB {
	c := util.ReadConfig()
	var DB *gorm.DB
	var err error
	if DB == nil {
		dsn := c.DSN //"root:@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
	}
	return DB
}
