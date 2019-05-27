package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//get the user id from the URL
	id := p.ByName("id")
	//verify that id is objectid
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //404
	}
	oid := bson.ObjectId(id) // return object id from hex object id
	u := models.User{}
	err := uc.session.DB("training").C("test").Find(oid).One(&u)
		if err != nil {
		w.WriteHeader(http.StatusNotFound) //Not found
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectId(id)
	err := uc.session.DB("training").C("test").Insert(oid)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectId(id)
	err := uc.session.DB("training").C("test").RemoveId(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
