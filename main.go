package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.Use(errorHandler)

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/user", addUser)

	router.Run(":3000")
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors) // -1 == not override the current error code
	}
}
