package storage

type Link struct {
	Id  string
	Url string
}

var LinkMap = map[string]*Link{}
