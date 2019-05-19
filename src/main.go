package main

import (
	"github.com/gin-gonic/gin"
	controller "project/Schedule/src/controllers"
)


func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./templates/*.tmpl")

	eCtrl := controller.EventController{}
	uCtrl := controller.UserController{}

	router.GET("/", eCtrl.EventIndex)
	router.POST("/Event", eCtrl.EventCreate)
	router.POST("/AddSchedule", uCtrl.AddSchedule)
	router.Run(":8080")
}

