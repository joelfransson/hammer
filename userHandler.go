package main

import (
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

type apiUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type userHandler struct {
	repo Users
}

func NewUserHandler(repo Users) *userHandler {
	return &userHandler{repo}
}

func NewApiUser(id string, name string, age int64) *apiUser {
	return &apiUser{id, name, age}
}

func (uc *userHandler) addUser(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)

	n := c.PostForm("name")
	a := c.PostForm("age")

	age, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	uc.repo.insertUser(db, n, age)
}

func (uc *userHandler) updateUser(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)

	var a apiUser
	if err := c.BindJSON(&a); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	u := NewUser(a.ID, a.Name, a.Age)

	uc.repo.updateUser(db, u)
}

func (uc *userHandler) getUsers(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)
	res, err := uc.repo.getAllUsers(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	data := make([]*apiUser, len(res))
	for i := range data {
		data[i] = NewApiUser(res[i].ID.Hex(), res[i].Name, res[i].Age)
	}

	c.JSON(http.StatusOK, data)
}

func (uc *userHandler) getUser(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)
	res, err := uc.repo.getUserByID(db, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	u := NewApiUser(res.ID.Hex(), res.Name, res.Age)

	c.JSON(http.StatusOK, u)
}
