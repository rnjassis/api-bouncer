package sqllite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const db_file string = "config_bouncer.db"

func Init() {
	db := createOrOpen()
	createTables(db)
}

func createOrOpen() *sql.DB {
	_, err := os.Stat(db_file)
	if err != nil {
		file, err := os.Create(db_file)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}

	database, _ := sql.Open("sqlite3", "./"+db_file)

	return database
}

func createTables(db *sql.DB) {
	createEndpointUrl := `CREATE TABLE endpoint (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"verb" TEXT,
	"url" TEXT,
	"return" TEXT)`

	statement, err := db.Prepare(createEndpointUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func Test_two() {
	fmt.Println("test 2")
}
