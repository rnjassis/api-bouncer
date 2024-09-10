package sqllite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
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

	database, _ := sql.Open("sqlite3", db_file)

	return database
}

func createTables(db *sql.DB) {
	createProject := `CREATE TABLE IF NOT EXISTS project (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"port" integer NOT NULL
	)`
	execStatement(db, createProject)

	createEndpointUrl := `CREATE TABLE IF NOT EXISTS endpoint (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"projectId" integer NOT NULL,
	"verb" TEXT,
	"url" TEXT,
	"return" TEXT,
	FOREIGN KEY("projectId") REFERENCES project("id")
	)`
	execStatement(db, createEndpointUrl)

}

func execStatement(db *sql.DB, _statement string) (sql.Result, error) {
	statement, err := db.Prepare(_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, error := statement.Exec()
	return result, error
}
