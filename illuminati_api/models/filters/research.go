package filters

import (
	"gopkg.in/mgo.v2/bson"
)

type Attribute struct {
	Key   string
	Value interface{}
}

type Parameter struct {
	Attributes []Attribute
	Name       string
}

type QueryString struct {
	Parameters []Parameter
}

type PaginationData struct {
	Offset int
	Limit  int
}

type ResearchData struct {
	Pagination          PaginationData
	Sorting             []string
	QueryStringToVerify QueryString
	QueryString         bson.M
}
