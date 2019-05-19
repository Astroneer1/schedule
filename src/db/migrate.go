package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"project/Schedule/src/models"
)

func gormConnect() *gorm.DB {
	DBMS    := "mysql"
	USER    := "root"
	PASS    := "mysql"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME  := "schedule"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main(){
	db := gormConnect()
	db.CreateTable(&models.Attendamce{})
	db.CreateTable(&models.Days{})
	db.CreateTable(&models.Event{})
	db.CreateTable(&models.People{})
	defer db.Close()
}