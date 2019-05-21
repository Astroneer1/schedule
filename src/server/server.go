package server

import (
	"github.com/gin-gonic/gin"
	"os"
	controller "project/Schedule/src/controllers"
)

func Init() {
	router := router()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func router() (router *gin.Engine) {
	router = gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./templates/*.tmpl")

	eCtrl := controller.EventController{}
	uCtrl := controller.UserController{}

	router.GET("/", eCtrl.EventIndex)
	router.POST("/Event", eCtrl.EventCreate)
	router.GET("/AddSchedule", uCtrl.AddSchedule)

	return
}