package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type EventController struct{}

func (ec EventController) EventIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

/*
	Event作成
　　　Eventsテーブル・Daysテーブルにインサート
 */
func (ec EventController) EventCreate (c *gin.Context) {
	dbConnection := db.GetConnection()

	eventName := c.PostForm("eventName")
	memo := c.PostForm("memo")
	days := c.PostForm("days")

	eventModel := models.Event{}
	eventModel.EventName = eventName
	eventModel.Memo = memo
	dbConnection.Create(&eventModel)
	eventId := eventModel.EventID

	dayModel := models.Days{}
	dayModel.EventID = eventId
	dayModel.Day = days
	dbConnection.Create(&dayModel)
	dayId := dayModel

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl",
		gin.H{"eventId": eventId, "eventName": eventName, "memo": memo, "days": days, "dayId": dayId})
}

func (ec EventController) EventUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{})
}