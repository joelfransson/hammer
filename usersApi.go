package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type apiUser struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func addUser(c *gin.Context) {
	n := c.PostForm("name")
	a := c.PostForm("age")

	age, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	insertUser(n, age)

	c.JSON(http.StatusOK, "")
}

func getUsers(c *gin.Context) {
	res, err := getAllUsers()
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

func getUser(c *gin.Context) {
	res, err := getUserByID(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	u := apiUser{res.Name, res.Age}

	c.JSON(http.StatusOK, u)
}
