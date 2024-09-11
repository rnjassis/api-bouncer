package models

type Project struct {
	Id          int
	Port        int
	Name        string
	Description string
	Endpoints   []Endpoint
}

type Endpoint struct {
	Id     int
	Verb   string
	Url    string
	Return string
}
