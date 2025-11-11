package argparser

import (
	"flag"
)

type Arguments struct {
	ListProjects    bool
	DescribeProject string

	RunProject  bool
	ProjectName string

	CreateProject      bool
	ProjectPort        string
	ProjectDescription string

	CreateRequest    bool
	RequestMethod    string
	RequestUrl       string
	RequestPreflight bool

	CreateResponse     bool
	ResponseStatusCode int
	ResponseBody       string
	ResponseMime       string
	ResponseIdentifier string
	ResponseHeaders    string
	ResponseIsRedirect bool
	ResponseIsProxy    bool

	DeleteProject  bool
	DeleteRequest  bool
	DeleteResponse bool
}

func ParseArgs() (Arguments, error) {
	listProjects := flag.Bool("list-projects", false, "List all available projects.")
	describeProject := flag.String("describe-project", "", "Show all the info about the project")

	runProject := flag.Bool("run-project", false, "Start the project server")
	projectName := flag.String("project-name", "", "Name of the project")

	createProject := flag.Bool("create-project", false, "Create a new project")
	projectPort := flag.String("project-port", "", "The port that will be used")
	projectDescription := flag.String("project-description", "", "The description of the new project")

	createRequest := flag.Bool("create-request", false, "Create a new request")
	requestMethod := flag.String("request-method", "", "Request method")
	requestUrl := flag.String("request-url", "", "URL of the new request")
	requestpreflight := flag.Bool("request-preflight", false, "Set if the url needs pre-flight response")

	createResponse := flag.Bool("create-response", false, "Create a new response for an existing request")
	responseStatusCode := flag.Int("response-status-code", 0, "Status code for the response")
	responseBody := flag.String("response-body", "", "Body that will be returned")
	responseMime := flag.String("response-mime", "", "Mime type")
	responseIdentifier := flag.String("response-identifier", "", "Exclusive identifier")
	responseHeaders := flag.String("response-headers", "", "Set the response headers")
	responseIsRedirect := flag.Bool("response-is-redirect", false, "Set if the response will redirect to another url")
	responseIsProxy := flag.Bool("response-is-proxy", false, "Set if the url is  proxy to another")

	deleteProject := flag.Bool("delete-project", false, "Deletes the entire project")
	deleteRequest := flag.Bool("delete-request", false, "Delete the request and all reponses associated with it")
	deleteResponse := flag.Bool("delete-response", false, "Delete a specific response")

	flag.Parse()

	args := Arguments{
		ListProjects:    *listProjects,
		DescribeProject: *describeProject,

		RunProject:  *runProject,
		ProjectName: *projectName,

		CreateProject:      *createProject,
		ProjectPort:        *projectPort,
		ProjectDescription: *projectDescription,

		CreateRequest:    *createRequest,
		RequestMethod:    *requestMethod,
		RequestUrl:       *requestUrl,
		RequestPreflight: *requestpreflight,

		CreateResponse:     *createResponse,
		ResponseStatusCode: *responseStatusCode,
		ResponseBody:       *responseBody,
		ResponseMime:       *responseMime,
		ResponseIdentifier: *responseIdentifier,
		ResponseHeaders:    *responseHeaders,
		ResponseIsRedirect: *responseIsRedirect,
		ResponseIsProxy:    *responseIsProxy,

		DeleteProject:  *deleteProject,
		DeleteRequest:  *deleteRequest,
		DeleteResponse: *deleteResponse,
	}

	error := argsValidation(args)

	return args, error
}
