package model

import (
	"GoPass/lib/es"
	"github.com/olivere/elastic/v7"
)

type Employee struct {
	///*Base
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func (l *Employee) GetIndex() string {
	return "megacorp"
}

func (l *Employee) GetConn() string {
	return "default"
}

func (l *Employee) Model() *elastic.IndexService {
	return es.Es(l.GetConn()).Index().Index(l.GetIndex()).Type("_doc")
}
