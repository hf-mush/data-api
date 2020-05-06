package model

import "gopkg.in/mgo.v2/bson"

// MgoTrainingLog training struct for mgo
type MgoTrainingLog struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Date  string        `bson:"date"`
	Count int           `bson:"count"`
	Kind  string        `bson:"kind"`
}
