package driver

import (
	"github.com/tkanos/gonfig"
	"hairdressing/models"
	"log"
	"testing"
)

func init() {
	Conf := models.Configuration{}
	err := gonfig.GetConf("environment.json", &Conf)
	if err != nil {
		log.Println("missing data base connection file", err)
	}
}

func TestCoonection(t *testing.T) {
	value := Coonection()
	if value == nil {
		t.Error("Error in database connection")
	}
}