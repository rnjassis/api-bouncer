package models

type Project struct {
	Id        int
	Port      int
	Endpoints []Endpoint
}

type Endpoint struct {
	Id     int
	Verb   string
	Url    string
	Return string
}
