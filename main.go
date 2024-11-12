package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/rnjassis/api-bouncer/argparser"
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

	return errors.New("argument not found")
}
