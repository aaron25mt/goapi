package dao

import (
  "log"

  . "github.com/aaron25mt/applications-api/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type ApplicationsDAO struct {
  Server   string
  Database string
}

var db *mgo.Database

const (
  COLLECTION = "applications"
)

func (a *ApplicationsDAO) Connect() {
  session, err := mgo.Dial(a.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(a.Database)
}

func (a *ApplicationsDAO) GetAll() ([]Application, error) {
  var applications []Application
  err := db.C(COLLECTION).Find(bson.M{}).All(&applications)
  return applications, err
}

func (a *ApplicationsDAO) GetById(id string) (Application, error) {
  var application Application
  err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&application)
  return application, err
}

func (a *ApplicationsDAO) Insert(application Application) error {
  err := db.C(COLLECTION).Insert(&application)
  return err
}

func (a *ApplicationsDAO) Update(id string, application Application) error {
  err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &application)
  return err
}

func (a *ApplicationsDAO) Delete(id string) error {
  err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
  return err
}
