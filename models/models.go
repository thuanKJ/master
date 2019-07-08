package models

import (
	"github.com/globalsign/mgo/bson"
)

type (
	Song struct {
		Id     string   `bson:"_id" json:"_id"`
		Name   string   `bson:"name" json:"name"`
		Artist []string `bson:"artist" json:"artist"`
		Kind   string   `bson:"kind" json:"kind"`
		Album  string   `bson:"album" json:"album"`
	}
	Account struct {
		Id       bson.ObjectId `bson:"_id" json:"_id"`
		Username string        `bson:"username" json "username"`
		Password string        `bson:"password" json:"password"`
	}
)
