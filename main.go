package main

import (
	"github.com/gin-gonic/gin"
	"hairdressing/http/handler"
)

func main() {

	r := gin.Default()

	r.GET("get/client/v1", func(context *gin.Context) {
		handler.GetPerson(context)
	})

	r.POST("/create/client/v1", func(context *gin.Context) {
		handler.CreatePerson(context)
	})

}
