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

func GetFullProject(db *sql.DB, projectName string) (*models.Project, error) {
	var err error
	project, err := GetProjectByName(db, projectName)
	if err != nil {
		return nil, err
	}
	project.Requests, err = GetRequests(db, project.Id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(project.Requests); i++ {
		project.Requests[i].Responses, err = GetResponses(db, project.Requests[i].Id)
		if err != nil {
			return nil, err
		}
	}

	return project, nil
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

func GetProjectByName(db *sql.DB, projectName string) (*models.Project, error) {
	project := &models.Project{}
	rows, err := selectStatement(db, getProjectByNameSql().sql, projectName)
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Port, &project.Name, &project.Description)
		if err != nil {
			return nil, nil //TODO add error
		}
	}
	return project, nil
}

func GetRequests(db *sql.DB, projectId int) ([]models.Request, error) {
	rows, err := selectStatement(db, getRequestsSql().sql, strconv.Itoa(projectId))
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	slice := []models.Request{}
	for rows.Next() {
		request := models.Request{}
		err = rows.Scan(&request.Id, &request.RequestMethod, &request.Url)
		if err != nil {
			return nil, nil //TODO add error
		}

		slice = append(slice, request)
	}

	return slice, nil
}

func GetResponses(db *sql.DB, requestId int) ([]models.Response, error) {
	rows, err := selectStatement(db, getResponseSQL().sql, strconv.Itoa(requestId))
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	slice := []models.Response{}
	for rows.Next() {
		response := models.Response{}
		err = rows.Scan(&response.Id, &response.StatusCode, &response.Active, &response.Body, &response.Mime)
		if err != nil {
			return nil, nil //TODO add error
		}
		slice = append(slice, response)
	}
	return slice, nil
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
	execStatement(db, createRequestTableSql().sql)
	execStatement(db, createResponseTableSql().sql)
}

func execStatement(db *sql.DB, _statement string) (sql.Result, error) {
	statement, err := db.Prepare(_statement)
	if err != nil {
		return nil, err
	}
	result, error := statement.Exec()
	return result, error
}

func selectStatement(db *sql.DB, statement string, params ...any) (*sql.Rows, error) {
	rows, err := db.Query(statement, params...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
