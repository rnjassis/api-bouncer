package argparser

import "errors"

func argValidate(args Arguments) error {
	if args.RunProject {
		if args.ProjectName == "" {
			return errors.New("project name not found")
		}
	}

	// No validation errors
	return nil
}
