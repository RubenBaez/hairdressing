package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hairdressing/bd"
	"hairdressing/models"
	"log"
)

func GetPerson(c *gin.Context) {
	var value []byte
	var values []byte
	var p []models.Person
	var person models.Person
	rows := bd.GetUser()
	for rows.Next() {
		err := rows.Scan(&person.Id, &person.Age, &person.First_name, &person.Last_name, &person.Email)
		if err != nil {
			log.Println("Error in get table rows")
		}
		value, err = json.Marshal(person)
		values = append(value)
		p = append(p, person)
		fmt.Println(values)
		if err != nil {
			log.Println("Error in person to try parse json")
		}
	}
	c.JSON(200, map[string][]models.Person{
		"ConsultClients":p,
	})
	p = nil
}
