package bd

import (
	"database/sql"
	"fmt"
	"hairdressing/driver"
	"hairdressing/models"
	"log"
)

func GetUser() *sql.Rows {

	rows, err  := driver.Coonection().Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error in make query", err)
	}
	return rows
}

func AddUser(person *models.Person) sql.Result {

	stmt, err := driver.Coonection().Prepare("insert into users values($1,$2,$3,$4,$5)")
	if err != nil {
		log.Println("Error in prepare statement to insert values", err)
	}

	rows, err := stmt.Exec(person.Id, person.Age, person.First_name, person.Last_name, person.Email)
	if err != nil {
		log.Println("Error in insert values", err)
	}
	fmt.Println(rows)
	return rows
}
