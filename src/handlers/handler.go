package handler

import (
	"net/http"

	service "github.com/Mpetrato/startupconnect/src/services"
	"github.com/Mpetrato/startupconnect/src/types"
	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := service.RegisterUserService(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
