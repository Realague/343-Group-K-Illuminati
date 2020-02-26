package utils

import "gopkg.in/mgo.v2"

func Sort(query *mgo.Query, fields []string) *mgo.Query {
	for _, field := range fields {
		query.Sort(field)
	}
	return query
}
