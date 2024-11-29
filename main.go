package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/rnjassis/api-bouncer/argparser"
	"github.com/rnjassis/api-bouncer/models"
	"github.com/rnjassis/api-bouncer/server"
	"github.com/rnjassis/api-bouncer/sqllite"
)

func main() {
	args, error := argparser.ParseArgs()
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}

	db := sqllite.Init()
	defer db.Close()

	error = run(db, args)
	if error != nil {
		fmt.Println(error.Error())
	}
}

func run(db *sql.DB, args argparser.Arguments) error {
	if args.ListProjects {
		projects, _ := sqllite.GetProjects(db)
		if len(projects) > 0 {
			fmt.Println("Available projects:")
			for _, v := range projects {
				fmt.Println("- ", v.Name)
			}
		} else {
			return errors.New("no registered projects")
		}

		return nil
	}

	if args.RunProject {
		project, error := sqllite.GetFullProject(db, args.Name)
		if error == nil {
			server.RunServer(project)
			return nil
		}
		return error
	}

	if args.CreateProject {
		project := &models.Project{
			Name:        args.ProjectName,
			Port:        args.Port,
			Description: args.Description,
		}

		error := sqllite.CreateProject(db, project)

		if error == nil {
			fmt.Println("Project created")
			return nil
		} else {
			return error
		}
	}

	if args.CreateRequest {
		project := &models.Project{Name: args.ProjectName}

		requestMethod, error := models.GetStatus(args.Method)
		if error != nil {
			return error
		}
		request := &models.Request{RequestMethod: requestMethod, Url: args.Url}

		error = sqllite.CreateRequest(db, project, request)
		if error == nil {
			fmt.Println("Request Created")
			return nil
		} else {
			return error
		}

	}

	if args.CreateResponse {
		project := &models.Project{Name: args.ProjectName}
		request := &models.Request{Url: args.RequestMethodUrl}
		response := &models.Response{Identifier: args.Identifier, Mime: args.Mime, Body: args.Body, Active: true}

		error := sqllite.CreateResponse(db, project, request, response)

		if error == nil {
			fmt.Println("Response Created")
			return nil
		} else {
			return error
		}
	}

	return errors.New("arguments not found")
}
