package main

import (
  "encoding/json"
  "log"
  "net/http"

  "gopkg.in/mgo.v2/bson"

  "github.com/gorilla/mux"
  . "github.com/aaron25mt/goapi/config"
  . "github.com/aaron25mt/goapi/dao"
  . "github.com/aaron25mt/goapi/models"
)

var config = Config{}
var dao = ApplicationsDAO{}

func init() {
  config.Read()

  dao.Server = config.Server
  dao.Database = config.Database
  dao.Connect()
}

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/applications", GetApplications).Methods("GET")
  router.HandleFunc("/applications", CreateApplication).Methods("POST")
  router.HandleFunc("/applications/{id}", GetApplication).Methods("GET")
  router.HandleFunc("/applications/{id}", UpdateApplication).Methods("PUT")
  router.HandleFunc("/applications/{id}", DeleteApplication).Methods("DELETE")

  if err := http.ListenAndServe(":8000", router); err != nil {
    log.Fatal(err)
  }
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
  respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func GetApplications(w http.ResponseWriter, r *http.Request) {
  applications, err := dao.GetAll()

  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusOK, applications)
}

func GetApplication(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if !bson.IsObjectIdHex(params["id"]) {
    respondWithError(w, http.StatusBadRequest, "Invalid application ID format")
    return
  }

  application, err := dao.GetById(params["id"])

  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid application ID")
    return
  }

  respondWithJson(w, http.StatusOK, application)
}

func CreateApplication(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  var application Application

  if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }

  application.ID = bson.NewObjectId()
  application.Company.ID = bson.NewObjectId()
  if err := dao.Insert(application); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusCreated, application)
}

func UpdateApplication(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  params := mux.Vars(r)
  var application Application

  if !bson.IsObjectIdHex(params["id"]) {
    respondWithError(w, http.StatusBadRequest, "Invalid application ID format")
    return
  }

  if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }

  if err := dao.Update(params["id"], application); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteApplication(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  if !bson.IsObjectIdHex(params["id"]) {
    respondWithError(w, http.StatusBadRequest, "Invalid application ID format")
    return
  }

  if err := dao.Delete(params["id"]); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
