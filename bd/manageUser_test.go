package bd

import (
	"hairdressing/models"
	"testing"
)

func TestGetUser(t *testing.T) {
	value := GetUser()
	if value == nil{
		t.Error("error in get rows")
	}
}

func TestAddUser(t *testing.T) {

	p := models.Person {
		Id:123,
		Age:24,
		First_name:"ruben",
		Last_name:"baez",
		Email:"rubenba@gmail.com",
	}

	value := AddUser(&p)
	if value == nil {
		t.Error("Error in insert values")
	}
}
