package server

import (
	controller "Schedule/src/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	router := router()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func router() (router *gin.Engine) {
	router = gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("src/templates/*.tmpl")

	eCtrl := controller.EventController{}
	aCtrl := controller.AttendanceController{}

	router.GET("/", eCtrl.EventIndex)
	router.POST("/Event", eCtrl.EventCreate)
	router.GET("/Attendance", aCtrl.AttendanceIndex)

	return
}