package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {}

func (uc UserController) AddSchedule (c *gin.Context) {
	c.HTML(http.StatusOK, "AddSchedule.tmpl", gin.H{})
}