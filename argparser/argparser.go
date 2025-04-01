package argparser

import (
	"flag"
)

type Arguments struct {
	ListProjects    bool
	DescribeProject string

	RunProject bool
	Name       string

	CreateProject bool
	Port          string
	Description   string

	CreateRequest bool
	ProjectName   string
	Method        string
	Url           string

	CreateResponse   bool
	RequestMethodUrl string
	StatusCode       int
	Body             string
	Mime             string
	Identifier       string
	Redirect         bool
	Headers          string
	Proxy            bool

	DeleteProject  bool
	DeleteRequest  bool
	DeleteResponse bool
}

func ParseArgs() (Arguments, error) {
	listProjects := flag.Bool("list-projects", false, "List all available projects.")
	describeProject := flag.String("describe-project", "", "Show all the info about the project")

	runProject := flag.Bool("run-project", false, "Start the project server")
	name := flag.String("name", "", "The name of the project")

	createProject := flag.Bool("create-project", false, "Create a new project")
	port := flag.String("port", "", "The port that will be used")
	description := flag.String("description", "", "The description of the new project")

	createRequest := flag.Bool("create-request", false, "Create a new request")
	projectName := flag.String("project-name", "", "Name of the existing project")
	method := flag.String("method", "", "Request method")
	url := flag.String("url", "", "URL of the new request")

	createResponse := flag.Bool("create-response", false, "Create a new response for an existing request")
	requestMethod := flag.String("request-method", "", "Request method that the new response will be related to")
	statusCode := flag.Int("status-code", 0, "Status code for the response")
	body := flag.String("body", "", "Body that will be returned")
	mime := flag.String("mime", "", "Mime type")
	identifier := flag.String("identifier", "", "Exclusive identifier")
	redirect := flag.Bool("is-redirect", false, "Set if the response will redirect to another url")
	headers := flag.String("headers", "", "Set the response headers")
	proxy := flag.Bool("proxy", false, "Set if the url is  proxy to another")

	deleteProject := flag.Bool("delete-project", false, "Deletes the entire project")
	deleteRequest := flag.Bool("delete-request", false, "Delete the request and all reponses associated with it")
	deleteResponse := flag.Bool("delete-response", false, "Delete a specific response")

	flag.Parse()

	args := Arguments{
		ListProjects:    *listProjects,
		DescribeProject: *describeProject,

		RunProject: *runProject,
		Name:       *name,

		CreateProject: *createProject,
		Port:          *port,
		Description:   *description,

		CreateRequest: *createRequest,
		ProjectName:   *projectName,
		Method:        *method,
		Url:           *url,

		CreateResponse:   *createResponse,
		RequestMethodUrl: *requestMethod,
		StatusCode:       *statusCode,
		Body:             *body,
		Mime:             *mime,
		Identifier:       *identifier,
		Redirect:         *redirect,
		Headers:          *headers,
		Proxy:            *proxy,

		DeleteProject:  *deleteProject,
		DeleteRequest:  *deleteRequest,
		DeleteResponse: *deleteResponse,
	}

	error := argsValidation(args)

	return args, error
}
