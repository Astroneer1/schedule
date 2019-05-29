package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type EventController struct{
	eventModel models.Event
	dayModel models.Days
}

func (ec EventController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

/*
	Event作成
　　　Eventsテーブル・Daysテーブルにインサート
 */
func (ec EventController) Create (c *gin.Context) {
	dbConnection := db.GetConnection()

	eventName := c.PostForm("eventName")
	memo := c.PostForm("memo")
	days := c.PostForm("days")

	ec.eventModel.EventName = eventName
	ec.eventModel.Memo = memo
	dbConnection.Create(&ec.eventModel)
	eventId := ec.eventModel.EventID

	ec.dayModel.EventID = eventId
	ec.dayModel.Day = days
	dbConnection.Create(&ec.dayModel)
	dayId := ec.dayModel.DayID

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl",
		gin.H{"eventId": eventId, "eventName": eventName, "memo": memo, "days": days, "dayId": dayId})
}

func (ec EventController) EventUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{})
}