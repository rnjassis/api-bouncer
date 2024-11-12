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
		if args.Port == "" {
			return errors.New("Port missing")
		}
	}

	// Create Request
	if args.CreateRequest {
		if args.ProjectName == "" {
			return errors.New("Provide the name of the project this request will be related to")
		}
		if args.Method == "" {
			return errors.New("Request method missing")
		}
		if args.Url == "" {
			return errors.New("Request url missing")
		}
	}

	// Create Response
	if args.CreateResponse {
		if args.ProjectName == "" {
			return errors.New("Provide the name of the project this response will related to")
		}
		if args.RequestMethodUrl == "" {
			return errors.New("Provide the request url this response will be related to")
		}
		if args.StatusCode == "" {
			return errors.New("Status code missing")
		}
		if args.Identifier == "" {
			return errors.New("Ideitifier missing")
		}
		if args.Body != "" {
			if args.Mime == "" {
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
		if args.RequestMethodUrl == "" {
			return errors.New("Provide request method Url")
		}
	}
	if args.DeleteResponse {
		if args.Identifier == "" {
			return errors.New("Provide unique identifier for the response")
		}
	}

	// No validation errors
	return nil
}
