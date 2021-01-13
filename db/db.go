package db

import (
	"database/sql"
	"project_2go/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("connectionString error")
	}

	err = db.Ping()
	if err != nil {
		panic("DNS Invalid")
	}

}

func CreateCon() *sql.DB {
	return db
}
