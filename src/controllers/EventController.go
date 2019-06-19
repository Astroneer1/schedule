package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type EventController struct{}

func (ec EventController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

/*
	Event作成
　　　Eventsテーブル・Daysテーブルにインサート
 */
func (ec EventController) Create (c *gin.Context) {
	dbConnection := db.GetConnection()

	eventModel := models.Event{
		EventName: c.PostForm("eventName"),
		Memo:c.PostForm("memo"),
	}

	dbConnection.Create(&eventModel)


	dayModel := models.Days{
		EventID: eventModel.EventID,
		Day: c.PostForm("days"),
	}

	dbConnection.Create(&dayModel)

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl",
		gin.H{"eventId": eventModel.EventID, "eventName": eventModel.EventName,
			"memo": eventModel.Memo, "days": dayModel.Day, "dayId": dayModel.DayID})
}

func (ec EventController) EventUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{})
}