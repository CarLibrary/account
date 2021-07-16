package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitMYSQL() {

	DB, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/carlibrary_user?charset=utf8mb4"))
	if err != nil {
		log.Fatal(err.Error())
	}
	db = DB
	if err:=DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{});err !=nil{
		log.Fatal(err.Error())
	}

}