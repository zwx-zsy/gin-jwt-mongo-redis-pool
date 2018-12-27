package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collectionName_Person = "bb_Person"

type Person struct {
	Id       bson.ObjectId `bson:"_id"`
	NickName string        `json:"Nickname" bson:"NickName"`
	Sex      int           `json:"Sex" bson:"Sex"`
	Birthday string        `json:"Birthday" bson:"Birthday"`
}

type Men []Person

func Persons() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(collectionName_Person)
}
