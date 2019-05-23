package controller

import (
	"Schedule/src/db"
	"Schedule/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type AttendanceController struct {}

func (ac AttendanceController) AttendanceIndex (c *gin.Context) {
	dbConnection := db.GetConnection()

	daysModel := models.Days{}
	dbConnection.Find(&daysModel, "event_id=?", c.Query("eventId"))

	db.CloseConnection(dbConnection)
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": daysModel.EventID,"day": daysModel.Day})
}

func (ac AttendanceController) AttendanceCreate (c *gin.Context) {
	dbConnection := db.GetConnection()
	peopleModel := models.People{}
	attendanceModel := models.Attendamce{}
	daysModel := models.Days{}
	eventModel := models.Event{}

	eventId := c.PostForm("eventId")
	name := c.PostForm("name")
	comment := c.PostForm("comment")

	peopleModel.EventID,_ = strconv.Atoi(eventId)
	peopleModel.Name = name
	peopleModel.Comment = comment

	dbConnection.Create(&peopleModel)


	peopleId := peopleModel.PeopleID
	attendance := c.PostForm("options")
	dayId := c.PostForm("dayId")

	attendanceModel.PeopleID = peopleId
	attendanceModel.Attendamce,_ = strconv.Atoi(attendance)
	attendanceModel.DayID,_ = strconv.Atoi(dayId)

	dbConnection.Create(&attendanceModel)

	dbConnection.Find(&daysModel, "event_Id=?", eventId)
	dbConnection.Find(&eventModel, "event_Id=?", eventId)

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl", gin.H{"eventId": eventId,
			"eventName": eventModel.EventName, "memo": eventModel.Memo, "days": daysModel.Day, "dayId": dayId})
}
