package db

import "gopkg.in/mgo.v2/bson"

type ForbiddenMailAddress struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name" binding:"required"`
}
