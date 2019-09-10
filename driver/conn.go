package driver

import (
	"database/sql"
	"fmt"
	"hairdressing/models"
	"log"

	_ "github.com/lib/pq"
	"github.com/tkanos/gonfig"
)

var conf models.Configuration

func init() {
	Conf := models.Configuration{}
	err := gonfig.GetConf("environment.json", &Conf)
	if err != nil {
		log.Println("missing data base connection file", err)
	}
}

func Coonection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		string(conf.Host), conf.Port, string(conf.User), string(conf.Password), string(conf.Password))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error database connection", err)
	}
	//defer db.Close()
	return db
}
