package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
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
		project, error := sqllite.GetFullProject(db, args.ProjectName, true)
		if error == nil {
			server.RunServer(project)
			return nil
		}
		return error
	}

	if args.CreateProject {
		port := args.ProjectPort
		if !strings.Contains(port, ":") {
			port = ":" + args.ProjectPort
		} else if strings.Index(port, ":") > 0 {
			return errors.New("incorrect port format")
		}
		project := &models.Project{
			Name:        args.ProjectName,
			Port:        port,
			Description: args.ProjectDescription,
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

		requestMethod, error := models.GetStatus(args.RequestMethod)
		if error != nil {
			return error
		}
		request := &models.Request{RequestMethod: requestMethod, Url: args.RequestUrl, Active: true}

		error = sqllite.CreateRequest(db, project, request)
		if error == nil {
			fmt.Println("Request Created")
		} else {
			return error
		}

		if args.RequestPreflight {
			request_pre := request
			request_pre.RequestMethod = models.RequestMethod("OPTIONS")
			response_pre := &models.Response{Identifier: uuid.NewString(), Mime: "text/html", Body: "", StatusCode: 200, Active: true}
			error = sqllite.CreateRequest(db, project, request_pre)
			if error != nil {
				error = sqllite.CreateResponse(db, project, request_pre, response_pre)
			}
			if error == nil {
				fmt.Println("Request pre-flight Created")
			} else {
				return error
			}
		}

		return nil
	}

	if args.CreateResponse {
		project := &models.Project{Name: args.ProjectName}
		request := &models.Request{Url: args.RequestUrl, RequestMethod: models.RequestMethod(args.RequestMethod)}
		response := &models.Response{Identifier: args.ResponseIdentifier, Mime: args.ResponseMime, Body: args.ResponseBody, StatusCode: args.ResponseStatusCode, Active: true, Redirect: args.ResponseIsRedirect, Headers: args.ResponseHeaders, Proxy: args.ResponseIsProxy}

		if response.Redirect {
			if request.RequestMethod == models.GET {
				response.StatusCode = 301 // moved permanently
			} else if request.RequestMethod == models.POST {
				response.StatusCode = 308 // permanent redirect
			}
		}

		error := sqllite.CreateResponse(db, project, request, response)

		if error == nil {
			fmt.Println("Response Created")
			return nil
		} else {
			return error
		}
	}

	if args.DeleteProject {
		// TODO
	}

	if args.DeleteRequest {
		// TODO
	}

	if args.DeleteResponse {
		// TODO
	}

	return errors.New("arguments not found")
}
