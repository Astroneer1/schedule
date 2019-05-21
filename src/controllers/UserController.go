package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"project/Schedule/src/db"
	"project/Schedule/src/models"
)

type UserController struct {}

func (uc UserController) AddSchedule (c *gin.Context) {
	dbConnection := db.GetConnection()

	days := models.Days{}
	fmt.Println(c.Query("eventId"))
	dbConnection.Find(&days, "event_id=?", c.Query("eventId"))

	db.CloseConnection(dbConnection)
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": days.EventID,"day": days.Day})
}