package main

import (
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

type apiUser struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type userController struct {
	repo UserRepository
}

func NewUserController() *userController {
	return &userController{}
}

func (uc *userController) addUser(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)

	n := c.PostForm("name")
	a := c.PostForm("age")

	age, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	uc.repo.insertUser(db, n, age)

	c.JSON(http.StatusOK, "")
}

func (uc *userController) getUsers(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)
	res, err := uc.repo.getAllUsers(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	data := make([]apiUser, len(res))
	for i := range data {
		data[i] = apiUser{res[i].Name, res[i].Age}
	}

	c.JSON(http.StatusOK, data)
}

func (uc *userController) getUser(c *gin.Context) {
	db := c.MustGet("mongo").(*mgo.Database)
	res, err := uc.repo.getUserByID(db, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	u := apiUser{res.Name, res.Age}

	c.JSON(http.StatusOK, u)
}
