package router

import (
	"github.com/Abhi-singh-karuna/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/person/:id/info/", handler.GetData)
	router.POST("/person/create", handler.CreateData)
}
