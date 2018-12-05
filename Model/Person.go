package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
)

const collectionName = "SS_Person"

type Person struct {
	Name     string `bson:"Name"`
	PassWord string `bson:"PassWord"`
}

type Men []Person

func Persons() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(collectionName)
}
