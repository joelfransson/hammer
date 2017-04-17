package main

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

var SESSION *mgo.Session
var DBNAME = "hammer"

func init() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("DB issues")
	}
	SESSION = session
}

func main() {
	router := gin.Default()
	router.Use(mapMongo)
	router.Use(errorHandler)

	ur := NewUserRepository()
	uc := NewUserHandler(ur)

	router.GET("/users", uc.getUsers)
	router.GET("/users/:id", uc.getUser)
	router.POST("/user", uc.addUser)
	router.PUT("/user/:id", uc.updateUser)

	router.Run(":3000")
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors) // -1 == not override the current error code
	}
}

func mapMongo(c *gin.Context) {
	s := SESSION.Clone()

	defer s.Close()

	c.Set("mongo", s.DB(DBNAME))
	c.Next()
}
