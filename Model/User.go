package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const collectionName_User string = "bb_User"

type User struct {
	Id             bson.ObjectId `bson:"_id"`
	WxOpenId       string        `json:"WxOpenId" bson:"WxOpenId""`
	UserCode       string        `json:"UserCode" bson:"UserCode"`
	PersonId       string        `json:"PersonId" bson:"PersonId"`
	CreateDateTime *time.Time    `json:"CreateDateTime" bson:"CreateDateTime"`
}

func Users() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(collectionName_User)
}
