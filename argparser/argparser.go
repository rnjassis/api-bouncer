package argparser

import (
	"flag"
)

type Arguments struct {
	ListProjects    bool
	DescribeProject string

	RunProject  bool
	ProjectName string
}

func ParseArgs() (Arguments, error) {
	listProjects := flag.Bool("list-projects", false, "List all available projects.")
	describeProject := flag.String("describe-project", "", "Show all the info about the project")
	runProject := flag.Bool("run-project", false, "Start the project server")
	projectName := flag.String("project-name", "", "The name of the project")

	flag.Parse()

	args := Arguments{
		ListProjects:    *listProjects,
		DescribeProject: *describeProject,
		RunProject:      *runProject,
		ProjectName:     *projectName,
	}

	error := argValidate(args)

	return args, error
}
