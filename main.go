package main

import (
	"github.com/gin-gonic/gin"
	"hairdressing/http/handler"
)

func main() {

	router := gin.Default()

	router.GET("get/client/v1", func(context *gin.Context) {
		handler.GetPerson(context)
	})

	router.POST("/create/client/v1", func(context *gin.Context) {
		handler.CreatePerson(context)
	})

	router.Run(":8000")
}
