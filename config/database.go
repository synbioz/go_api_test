package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=theo dbname=goapi")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()
}

func TestDatabaseInit() {
	connection, err := sql.Open("postgres", "user=theo")
	_, err = connection.Exec("CREATE DATABASE goapi_test")

	connection.Close()

	db, err = sql.Open("postgres", "user=theo dbname=goapi_test")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()

}

func TestDatabaseDestroy() {
	db.Close()

	connection, err := sql.Open("postgres", "user=theo")
	_, err = connection.Exec("DROP DATABASE goapi_test")

	if err != nil {
		log.Fatal(err)
	}
}

func createCarsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
