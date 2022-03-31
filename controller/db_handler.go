package controller

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/db_bobobox?parseTime=true&loc=Asia%2fJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
