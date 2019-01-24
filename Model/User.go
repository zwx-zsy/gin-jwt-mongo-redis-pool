package Models

import (
	"TimeLine/Lib"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionName_User string = "bb_User"

type User struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	WxOpenId       string        `bson:"WxOpenId"`
	UserCode       string        `bson:"UserCode"`
	PersonId       string        `bson:"PersonId"`
	CreateDateTime time.Time     `bson:"CreateDateTime"`
}

func Users() *mgo.Collection {
	return Lib.DB.C(CollectionName_User)
}
