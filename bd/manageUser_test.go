package bd

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	value := AddUser()
	if value == nil{
		t.Error("error in get rows")
	}
}
