package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:mysql@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
