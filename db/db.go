package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kondo1018008/golangPractice/entity"
)

var (
	db *gorm.DB
	err error
)

func Init(){
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB{
	return db
}

func Close(){
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration(){
	db.AutoMigrate(&entity.User{})
}