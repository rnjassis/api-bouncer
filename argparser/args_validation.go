//lint:file-ignore ST1005 errors that will be returned to the user
package argparser

import (
	"errors"
)

func argsValidation(args Arguments) error {
	// Run Project
	if args.RunProject {
		if args.ProjectName == "" {
			return errors.New("Project name not found")
		}
	}

	// Create Project
	if args.CreateProject {
		if args.ProjectName == "" {
			return errors.New("Project name missing")
		}
		if args.ProjectPort == "" {
			return errors.New("Port missing")
		}
	}

	// Create Request
	if args.CreateRequest {
		if args.ProjectName == "" {
			return errors.New("Provide the name of the project this request will be related to")
		}
		if args.RequestMethod == "" {
			return errors.New("Request method missing")
		}
		if args.RequestUrl == "" {
			return errors.New("Request url missing")
		}
	}

	// Create Response
	if args.CreateResponse {
		if args.ProjectName == "" {
			return errors.New("Provide the name of the project this response will related to")
		}
		if args.RequestMethod == "" {
			return errors.New("Provide the request method this response will be related to")
		}
		if args.RequestUrl == "" {
			return errors.New("Provide the url the response will be related")
		}
		if !args.ResponseIsRedirect {
			if args.ResponseStatusCode == 0 {
				return errors.New("Status code missing")
			}
		}
		if args.ResponseIdentifier == "" {
			return errors.New("Identifier missing")
		}
		if args.ResponseBody != "" {
			if args.ResponseMime == "" {
				return errors.New("Mime type missing")
			}
		}
	}

	// Deletion
	if args.DeleteProject || args.DeleteRequest || args.DeleteResponse {
		if args.ProjectName == "" {
			return errors.New("Provide project name")
		}
	}
	if args.DeleteRequest || args.DeleteResponse {
		if args.RequestUrl == "" {
			return errors.New("Provide request Url")
		}
	}
	if args.DeleteResponse {
		if args.ResponseIdentifier == "" {
			return errors.New("Provide unique identifier for the response")
		}
	}

	// No validation errors
	return nil
}
