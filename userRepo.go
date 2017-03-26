package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	id   interface{}
	Name string
	Age  int64
}

func getUC() (*mgo.Collection, *mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, nil, err
	}
	return session.DB("hammer").C("users"), session, err
}

func insertUser(name string, age int64) error {
	c, session, err := getUC()
	if err != nil {
		return fmt.Errorf("Failed to open DB connection %v", err)
	}

	if err = c.Insert(&user{nil, name, age}); err != nil {
		return fmt.Errorf("Failed when inserting user %v", err)
	}
	defer session.Close()

	return nil
}

func getAllUsers() ([]user, error) {
	c, session, err := getUC()
	if err != nil {
		return nil, fmt.Errorf("Failed to open DB connection %v", err)
	}

	var results []user
	if err = c.Find(nil).All(&results); err != nil {
		return nil, fmt.Errorf("Failed when getting all users %v", err)
	}
	defer session.Close()

	return results, err
}

func getUserByID(id string) (user, error) {
	var u user
	c, session, err := getUC()
	if err != nil {
		return u, fmt.Errorf("Failed to open DB connection %v", err)
	}

	u = user{}
	if err = c.FindId(bson.ObjectIdHex(id)).One(&u); err != nil {
		return u, fmt.Errorf("Failed when getting a user with id %s, %v", id, err)
	}
	defer session.Close()

	return u, nil
}
