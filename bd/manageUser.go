package bd

import (
	"database/sql"
	"fmt"
	"hairdressing/driver"
	"log"
)

func AddUser() *sql.Rows {

	rows, err  := driver.Coonection().Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error in make query", err)
	}
	fmt.Println(rows)

	for rows.Next() {
		var id int
		var age int
		var first_name string
		var last_name string
		var email string
		err = rows.Scan(&id, &age, &first_name, &last_name, &email)
		if err != nil {
			log.Println("Error in get table rows")
		}
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%10v | %4v | %8v | %6v | %6v\n", id, age, first_name, last_name, email)
	}

	return rows
}
