package main

import (
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserControllers(getsession())
	router.GET("/user/:id",uc.GetUser)
	router.POST("/user/:id",uc.CreateUser)
	router.DELETE("/user/:id",uc.DeleteUser)
	log.Println(http.ListenAndServe(":8080",router))

}

func getsession() *mgo.Session{
	s, err:= mgo.Dial("mongodb://localhost")
	if err != nil{
		panic(err)
	}
	return s
}
