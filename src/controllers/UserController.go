package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"project/Schedule/src/models"
	"fmt"
)

type UserController struct {}

func (uc UserController) AddSchedule (c *gin.Context) {
	db,_ := gorm.Open("mysql", "root:mysql@/schedule?charset=utf8&parseTime=True&loc=Local")

	days := models.Days{}
	fmt.Println(c.Query("eventId"))
	db.Find(&days, "event_id=?", c.Query("eventId"))

	defer db.Close()
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{"eventId": days.EventID,"day": days.Day})
}