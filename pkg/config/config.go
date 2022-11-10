package config

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:Bega030394/@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
