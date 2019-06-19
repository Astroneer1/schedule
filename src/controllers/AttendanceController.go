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

func (ac AttendanceController) Index (c *gin.Context) {
	dbConnection := db.GetConnection()

	daysModel := models.Days{}
	dbConnection.Find(&daysModel, "event_id=?", c.Query("eventId"))

	db.CloseConnection(dbConnection)
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": daysModel.EventID,"day": daysModel.Day})
}

func (ac AttendanceController) Create (c *gin.Context) {
	dbConnection := db.GetConnection()

	peopleModel := models.People{
		Name:    c.PostForm("name"),
		Comment: c.PostForm("comment"),
	}

	peopleModel.EventID,_ = strconv.Atoi(c.PostForm("eventId"))
	dbConnection.Create(&peopleModel)


	attendanceModel := models.Attendamce{
		PeopleID: peopleModel.PeopleID,
	}

	attendanceModel.Attendamce,_ = strconv.Atoi(c.PostForm("options"))
	attendanceModel.DayID,_ = strconv.Atoi(c.PostForm("dayId"))

	dbConnection.Create(&attendanceModel)

	daysModel := models.Days{}
	eventModel := models.Event{}

	dbConnection.Find(&daysModel, "event_Id=?", peopleModel.EventID)
	dbConnection.Find(&eventModel, "event_Id=?", peopleModel.EventID)

	db.CloseConnection(dbConnection)

	c.HTML(http.StatusOK, "Schedule.tmpl", gin.H{"eventId": eventModel.EventID,
			"eventName": eventModel.EventName, "memo": eventModel.Memo, "days": daysModel.Day, "dayId": daysModel.DayID})
}
