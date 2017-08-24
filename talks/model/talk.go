package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Talk struct {
	ID      bson.ObjectId `json:"talkid" bson:"_id"`
	Title   string        `json:"title" bson:"title"`
	Summary string        `json:"summary" bson:"summary"`
	When    time.Time     `json:"when" bson:"when"`
}
