package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionName_GrowthStandard = "bb_GrowthStandards"

type GrowthStandard struct {
	Id      bson.ObjectId `bson:"_id"`
	Days    int           `bson:"Days"`
	Title   string        `bson:"Title"`
	Context string        `bson:"Context"`
	Type    int           `bson:"Type"`
}

func GrowthStandards() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(CollectionName_GrowthStandard)
}
