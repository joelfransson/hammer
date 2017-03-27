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

type Users interface {
	insertUser(db *mgo.Database, name string, age int64) error
	getAllUsers(db *mgo.Database) ([]user, error)
	getUserByID(db *mgo.Database, id string) (user, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func getUC(db *mgo.Database) *mgo.Collection {
	return db.C("users")
}

func (r *UserRepository) insertUser(db *mgo.Database, name string, age int64) error {
	c := getUC(db)

	if err := c.Insert(&user{nil, name, age}); err != nil {
		return fmt.Errorf("Failed when inserting user %v", err)
	}

	return nil
}

func (r *UserRepository) getAllUsers(db *mgo.Database) ([]user, error) {
	c := getUC(db)

	var results []user
	if err := c.Find(nil).All(&results); err != nil {
		return nil, fmt.Errorf("Failed when getting all users %v", err)
	}

	return results, nil
}

func (r *UserRepository) getUserByID(db *mgo.Database, id string) (user, error) {
	c := getUC(db)

	var u = user{}
	if err := c.FindId(bson.ObjectIdHex(id)).One(&u); err != nil {
		return u, fmt.Errorf("Failed when getting a user with id %s, %v", id, err)
	}

	return u, nil
}
