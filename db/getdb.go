package db

import (
	//"github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(sqlite.Open("baoxiu.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
