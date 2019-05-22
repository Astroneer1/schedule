package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type AttendanceController struct {}

func (ac AttendanceController) AttendanceIndex (c *gin.Context) {
	dbConnection := db.GetConnection()

	days := models.Days{}
	fmt.Println(c.Query("eventId"))
	dbConnection.Find(&days, "event_id=?", c.Query("eventId"))

	db.CloseConnection(dbConnection)
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": days.EventID,"day": days.Day})
}
