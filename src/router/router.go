package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	InitRoutes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
