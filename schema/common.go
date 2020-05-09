package schema

import "github.com/globalsign/mgo/bson"

type FilterParam struct {
	Condition bson.M `json:"condition"`
	Skip      int    `json:"skip"`
	Limit     int    `json:"limit"`
	Sort      string `json:"sort"`
}