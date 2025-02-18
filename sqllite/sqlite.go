package sqllite

import (
	"database/sql"
	"errors"
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

func GetFullProject(db *sql.DB, projectName string, isActive bool) (*models.Project, error) {
	var err error
	project, err := GetProjectByName(db, projectName)
	if err != nil {
		return nil, err
	}
	project.Requests, err = GetRequests(db, project.Id, isActive)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(project.Requests); i++ {
		project.Requests[i].Responses, err = GetResponses(db, project.Requests[i].Id, isActive)
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
	rows, err := selectStatement(db, getProjectByNameSql().sql, projectName)
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	if rows.Next() {
		project := &models.Project{}
		err = rows.Scan(&project.Id, &project.Port, &project.Name, &project.Description)
		if err != nil {
			return nil, nil //TODO add error
		}
		return project, nil
	}
	return nil, nil
}

func GetRequests(db *sql.DB, projectId int, isActive bool) ([]models.Request, error) {
	rows, err := selectStatement(db, getRequestsSql(isActive).sql, strconv.Itoa(projectId))
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	slice := []models.Request{}
	for rows.Next() {
		request := models.Request{}
		err = rows.Scan(&request.Id, &request.RequestMethod, &request.Url, &request.Active)
		if err != nil {
			return nil, nil //TODO add error
		}

		slice = append(slice, request)
	}

	return slice, nil
}

func GetRequestByProjectUrl(db *sql.DB, projectName string, requestUrl string, requestMethod string) (*models.Request, error) {
	rows, err := selectStatement(db, getRequestByProjectUrlSql().sql, projectName, requestUrl, requestMethod)
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	if rows.Next() {
		request := &models.Request{}
		err = rows.Scan(&request.Id, &request.RequestMethod, &request.Url, &request.Active)
		if err != nil {
			return nil, nil //TODO add error
		}
		return request, nil
	}

	return nil, nil
}

func GetResponses(db *sql.DB, requestId int, isActive bool) ([]models.Response, error) {
	rows, err := selectStatement(db, getResponseByRequestIdSql(isActive).sql, strconv.Itoa(requestId))
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()

	slice := []models.Response{}
	for rows.Next() {
		response := models.Response{}
		err = rows.Scan(&response.Id, &response.StatusCode, &response.Active, &response.Body, &response.Mime, &response.Identifier, &response.Redirect)
		if err != nil {
			return nil, nil //TODO add error
		}
		slice = append(slice, response)
	}
	return slice, nil
}

func GetResponseByProjectRequestResponse(db *sql.DB, project string, request string, response string) (*models.Response, error) {
	rows, err := selectStatement(db, getResponseByProjectRequestResponseSql().sql,
		project, request, response)
	if err != nil {
		return nil, nil //TODO add error
	}
	defer rows.Close()
	if rows.Next() {
		response := &models.Response{}
		err = rows.Scan(&response.Id, &response.StatusCode, &response.Active, &response.Body, &response.Mime, &response.Identifier, &response.Redirect)
		if err != nil {
			return nil, nil //TODO add error
		}
		return response, nil
	}

	return nil, nil
}

func CreateProject(db *sql.DB, project *models.Project) error {
	proj, _ := GetProjectByName(db, project.Name)
	if proj != nil {
		return errors.New("Project \"" + project.Name + "\" does exist")
	}
	_, error := execStatement(db, createProjectSql().sql, project.Name, project.Port, project.Description)
	if error != nil {
		return errors.New("Error creating project: " + error.Error())
	}
	return nil
}

func CreateRequest(db *sql.DB, project *models.Project, request *models.Request) error {
	proj, _ := GetProjectByName(db, project.Name)
	if proj == nil {
		return errors.New("Project \"" + project.Name + "\" does not exist")
	}
	req, _ := GetRequestByProjectUrl(db, project.Name, request.Url, string(request.RequestMethod))
	if req != nil {
		return errors.New("Request" + request.Url + "already exist")
	}

	result, error := execStatement(db, createRequestSql().sql, request.RequestMethod, request.Url, request.Active, project.Name)
	if result != nil {
	}
	if error != nil {
		return errors.New("Error creating request: " + error.Error())
	}
	return nil
}

func CreateResponse(db *sql.DB, project *models.Project, request *models.Request, response *models.Response) error {
	proj, _ := GetProjectByName(db, project.Name)
	if proj == nil {
		return errors.New("Project \"" + project.Name + "\" does not exist")
	}
	req, _ := GetRequestByProjectUrl(db, project.Name, request.Url, string(request.RequestMethod))
	if req == nil {
		return errors.New("Request \"" + request.Url + "\" does not exist")
	}
	res, _ := GetResponseByProjectRequestResponse(db, project.Name, request.Url, response.Identifier)
	if res != nil {
		return errors.New("Response \"" + res.Identifier + "\" already exist")
	}
	_, error := execStatement(db, createResponseSql().sql, response.StatusCode, response.Active, response.Body, response.Mime, response.Identifier, project.Name, request.Url)
	if error != nil {
		return errors.New("Error creating response: " + error.Error())
	}
	return nil
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

func execStatement(db *sql.DB, _statement string, params ...any) (sql.Result, error) {
	statement, err := db.Prepare(_statement)
	if err != nil {
		return nil, err
	}
	result, error := statement.Exec(params...)
	return result, error
}

func selectStatement(db *sql.DB, statement string, params ...any) (*sql.Rows, error) {
	rows, err := db.Query(statement, params...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
