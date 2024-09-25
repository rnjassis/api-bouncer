package models

type Project struct {
	Id          int
	Port        int
	Name        string
	Description string
	Requests   []Request
}

type Request struct {
	Id     int
	Verb   string
	Url    string
	Responses []Response
}

type Response struct {
    Id int
    StatusCode int
    Active bool
    Mime string
    Body string
}
