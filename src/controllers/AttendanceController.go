package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type AttendanceController struct {
	daysModel models.Days
	eventModel models.Event
	attendanceModel models.Attendamce
	peopleModel models.People
}

func (ac AttendanceController) Index (c *gin.Context) {
	dbConnection := db.GetConnection()

	dbConnection.Find(&ac.daysModel, "event_id=?", c.Query("eventId"))

	db.CloseConnection(dbConnection)
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": ac.daysModel.EventID,"day": ac.daysModel.Day})
}

func (ac AttendanceController) Create (c *gin.Context) {
	dbConnection := db.GetConnection()

	eventId := c.PostForm("eventId")
	name := c.PostForm("name")
	comment := c.PostForm("comment")

	ac.peopleModel.EventID,_ = strconv.Atoi(eventId)
	ac.peopleModel.Name = name
	ac.peopleModel.Comment = comment

	dbConnection.Create(&ac.peopleModel)


	peopleId := ac.peopleModel.PeopleID
	attendance := c.PostForm("options")
	dayId := c.PostForm("dayId")

	ac.attendanceModel.PeopleID = peopleId
	ac.attendanceModel.Attendamce,_ = strconv.Atoi(attendance)
	ac.attendanceModel.DayID,_ = strconv.Atoi(dayId)

	dbConnection.Create(&ac.attendanceModel)

	dbConnection.Find(&ac.daysModel, "event_Id=?", eventId)
	dbConnection.Find(&ac.eventModel, "event_Id=?", eventId)

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl", gin.H{"eventId": eventId,
			"eventName": ac.eventModel.EventName, "memo": ac.eventModel.Memo, "days": ac.daysModel.Day, "dayId": dayId})
}
