package model
import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type TimeStamp struct {
	CreatedTime time.Time `json:"created_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type Deleted struct  {
	IsDeleted bool `json:"deleted"`
}

type Common struct {
	TimeStamp
	Deleted
	Id bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
}