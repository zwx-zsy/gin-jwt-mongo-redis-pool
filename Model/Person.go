package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const CollectionName_Person = "bb_Person"

type Person struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	NickName       string        `json:"Nickname" bson:"NickName"`
	Sex            int           `json:"Sex" bson:"Sex"`
	Birthday       string        `json:"Birthday" bson:"Birthday"`
	Born           int           `json:"Born" bson:"Born"`
	Role           int           `json:"Role" bson:"Role"`
	CreateDateTime time.Time     `json:"CreateDateTime" bson:"CreateDateTime"`
}

type Men []Person

func Persons() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(CollectionName_Person)
}
