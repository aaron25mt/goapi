package models

import "gopkg.in/mgo.v2/bson"

type Application struct {
  ID      bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
  Company *Company      `bson:"company" json:"company"`
  Status  string        `bson:"status" json:"status"`
}

type Company struct {
  ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
  Name      string        `bson:"name" json:"name"`
  Location  string        `bson:"location" json:"location"`
}
