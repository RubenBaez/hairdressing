package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hairdressing/bd"
	"hairdressing/models"
	"log"
)

func GetPerson(c *gin.Context) {
	var p []models.Person
	var person models.Person
	rows := bd.GetUser()
	for rows.Next() {
		err := rows.Scan(&person.Id, &person.Age, &person.First_name, &person.Last_name, &person.Email)
		if err != nil {
			log.Println("Error in get table rows")
		}
		_, err = json.Marshal(person)
		p = append(p, person)
		if err != nil {
			log.Println("Error in person to try parse json")
		}
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")
	c.JSON(200, map[string][]models.Person{
		"ConsultClients":p,
	})
	p = nil
}
