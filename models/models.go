package models

import "errors"

type Project struct {
	Id          int
	Port        string
	Name        string
	Description string
	Requests    []Request
}

type Request struct {
	Id            int
	RequestMethod RequestMethod
	Url           string
	Responses     []Response
	Active        bool
}

type Response struct {
	Id         int
	Identifier string
	StatusCode int
	Active     bool
	Mime       string
	Body       string
}

type RequestMethod string

const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
)

func GetStatus(value string) (RequestMethod, error) {
	switch value {
	case string(GET):
		return GET, nil
	case string(POST):
		return POST, nil
	case string(PUT):
		return PUT, nil
	case string(DELETE):
		return DELETE, nil
	default:
		return "", errors.New("Request Method not found")
	}
}
