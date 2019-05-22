package db

import (
	"Schedule/src/models"
)

func main(){
	createTables()
}

func createTables() {
	db := GetConnection()

	db.CreateTable(&models.Attendamce{})
	db.CreateTable(&models.Days{})
	db.CreateTable(&models.Event{})
	db.CreateTable(&models.People{})

	CloseConnection(db)
}

func deleteTables() {
	db := GetConnection()

	db.DropTable(&models.Attendamce{})
	db.DropTable(&models.Days{})
	db.DropTable(&models.Event{})
	db.DropTable(&models.People{})

	CloseConnection(db)
}