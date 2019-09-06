package models

type Person struct {
	Id         int    `json:id`
	Age        int    `json:age`
	First_name string `json:first_name`
	Last_name  string `json:last_name`
	Email      string `json:email`
}
