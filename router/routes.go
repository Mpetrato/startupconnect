package router

import (
	handler "github.com/Mpetrato/startupconnect/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/")
	{
		v1.GET("register", handler.RegisterUserHandler)
	}
}
