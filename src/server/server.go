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
	router.Static("../assets", "src/assets")
	router.LoadHTMLGlob("src/templates/*.tmpl")

	eCtrl := controller.EventController{}
	aCtrl := controller.AttendanceController{}

	router.GET("/", eCtrl.Index)
	router.POST("/Event", eCtrl.Create)
	router.GET("/Attendance", aCtrl.Index)
	router.POST("/Attendance", aCtrl.Create)

	return
}