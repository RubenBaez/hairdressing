package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hairdressing/bd"
	"hairdressing/models"
	"log"
)

func CreatePerson(c *gin.Context) {
	person := models.Person{}
	data, err := c.GetRawData()
	if err != nil {
		log.Println("error in ger data from post", err)
	}
	json.Unmarshal(data, &person)
	bd.AddUser(&person)
	c.JSON(201, map[string]string{
		"code":"0",
		"message":"registro guardado",
	})
}
