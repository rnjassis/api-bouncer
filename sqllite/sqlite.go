package sqllite

import (
	"database/sql"
	"log"
	"os"
    "strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rnjassis/api-bouncer/models"
)

const db_file_name string = "config_bouncer.db"

func Init() *sql.DB {
	db := createOrOpen()
	createTables(db)

	return db
}

func GetProjects(db *sql.DB) ([]models.Project, error) {
	rows, err := selectStatement(db, getProjectsSql().sql)
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	slice := []models.Project{}
	for rows.Next() {
		project := models.Project{}
		err = rows.Scan(&project.Id, &project.Port, &project.Name, &project.Description)

		if err != nil {
			return nil, nil //TODO add error
		}

		slice = append(slice, project)
	}

	return slice, nil
}

func GetRequests(db *sql.DB, projectName string) ([]models.Request, error) {
	rows, err := selectStatement(db, getRequestsSql().sql, projectName)
    if err != nil {
        return nil, nil //TODO add error
    }
    defer rows.Close()

    slice := []models.Request{}
    for rows.Next() {
        request := models.Request{}
        err = rows.Scan(&request.Id, &request.Verb, &request.Url)
        if err != nil {
            return nil, nil //TODO add error
        }

        resp, err := getResponse(db, request.Id)
        if err != nil {
            return nil, nil //TODO add error
        }
        request.Return = resp

        slice = append(slice, request)
    }

	return slice, nil
}

func getResponse(db *sql.DB, requestId int) (*models.Response, error) {
    response := &models.Response{}
	rows, err := selectStatement(db, getResponseSQL().sql, strconv.Itoa(requestId))
    if err != nil {
        return response, nil //TODO add error
    }
    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&response.Id, &response.StatusCode, &response.Body, &response.Mime)
        if err!= nil {
            return nil, nil //TODO add error
        }
        return response, nil
    }
    return nil, nil //TODO add error
}

func createOrOpen() *sql.DB {
	_, err := os.Stat(db_file_name)
	if err != nil {
		file, err := os.Create(db_file_name)
		if err != nil {
			log.Fatal("Could not reach to the database due to the following error: " + err.Error())
		}
		file.Close()
	}

	database, _ := sql.Open("sqlite3", db_file_name)

	return database
}

func createTables(db *sql.DB) {
	execStatement(db, createProjectTableSql().sql)
	execStatement(db, createResponseTableSql().sql)
	execStatement(db, createRequestTableSql().sql)
}

func execStatement(db *sql.DB, _statement string) (sql.Result, error) {
	statement, err := db.Prepare(_statement)
	if err != nil {
		return nil, err
	}
	result, error := statement.Exec()
	return result, error
}

func selectStatement(db *sql.DB, statement string, params ...string) (*sql.Rows, error) {
	rows, err := db.Query(statement, params)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
