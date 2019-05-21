package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Init() (db *gorm.DB) {
	DBMS     := os.Getenv("DBMS")
	USER     := os.Getenv("USER_NAME")
	PASS     := os.Getenv("PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DB_IP") + ":" + os.Getenv("DB_PORT") + ")"
	_         = "?charset=utf8&parseTime=True&loc=Local"
	DBNAME   := os.Getenv("DB_NAME")

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	fmt.Println(CONNECT)
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return
}

func GetConnection() (db *gorm.DB) {
	db = Init()
	return
}

func CloseConnection(db *gorm.DB) {
	defer db.Close()
}
