package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/user", addUser)

	router.Run(":3000")

}
