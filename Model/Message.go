package Models

import (
	"TimeLine/Lib"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionName_Message = "bb_Message"

type Message struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	Title          string        `json:"Title" bson:"Title"`
	Content        string        `json:"Content" bson:"Content"`
	CreateDateTime time.Time     `json:"CreateDateTime,omitempty" bson:"CreateDateTime"`
}

func Messages() *mgo.Collection {
	return Lib.DB.C(CollectionName_Message)
}
