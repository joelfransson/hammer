package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	id   interface{}
	Name string
	Age  int64
}

func getUC() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session.DB("hammer").C("users"), session
}

func insertUser(name string, age int64) {
	c, session := getUC()
	err := c.Insert(&user{nil, name, age})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
}

func getAllUsers() []user {
	c, session := getUC()

	var results []user
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	return results
}

func getUserByID(id string) user {
	c, session := getUC()

	result := user{}
	err := c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	return result
}
