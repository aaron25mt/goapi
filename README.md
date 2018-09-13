# GoAPI

A simple CRUD REST API designed using [Golang](https://golang.org) and [MongoDB](https://www.mongodb.com).

### Requirements
- [MUX](https://github.com/gorilla/mux)
- [BSON](https://godoc.org/labix.org/v2/mgo/bson)
- [TOML](https://github.com/BurntSushi/toml)

### How to Run
```
go build && ./goapi
```

### Routes

| Route | Description | Body |
| ------ | ------ | ------ |
| GET / | Get all applications | |
| POST / | Create new application | Application |
| GET /{id} | Get a single application | |
| PUT /{id} | Update a single application | Application |
| DELETE /{id} | Delete a single application | |

### Objects

##### Application

```
{
    "company": {
        "name": "",
        "location": "",
    },
    "status": ""
}
```
