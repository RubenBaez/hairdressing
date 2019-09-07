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
	rows := bd.GetUser()
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.Id, &person.Age, &person.First_name, &person.Last_name, &person.Email)
		if err != nil {
			log.Println("Error in get table rows")
		}
		fmt.Println(person.Id, person.Age, person.First_name, person.Last_name, person.Email)
		value, err := json.Marshal(person)
		if err != nil {
			log.Println("Error in person to try parse json")
		}
		c.Writer.Header().Set("Content-Type","application/json")
		c.Writer.WriteHeader(200)
		c.Writer.Write(value)
	}
}
