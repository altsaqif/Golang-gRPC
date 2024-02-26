package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:4sH13ybd@19^_^@tcp(localhost:3306)/go_grpc"))
	if err != nil {
		log.Fatalf("Database Connection Fail! %v", err.Error())
	}

	return db
}
