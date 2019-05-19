package controller

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"project/Schedule/src/models"
)

type EventController struct{}

func (ec EventController) EventIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}


func (ec EventController) EventCreate (c *gin.Context) {
	db,_ := gorm.Open("mysql", "root:mysql@/schedule?charset=utf8&parseTime=True&loc=Local")

	eventName := c.PostForm("eventName")
	memo := c.PostForm("memo")
	days := c.PostForm("days")

	event := models.Event{}
	event.EventName = eventName
	event.Memo = memo
	db.Create(&event)

	fmt.Println(event)

	dayModel := models.Days{}
	dayModel.EventID = event.EventID
	dayModel.Day = days
	db.Create(&dayModel)

	fmt.Println(event.EventID)
	fmt.Println(event.Memo)
	fmt.Println(event.EventName)
	fmt.Println(days)

	defer db.Close()

	c.HTML(http.StatusOK, "Schedule.tmpl", gin.H{"eventName": eventName, "memo": memo, "days": days})
}