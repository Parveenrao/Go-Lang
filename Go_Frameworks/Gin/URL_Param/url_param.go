// What is a URL Parameter?

// A URL parameter (path parameter) is a dynamic value inside the URL path.   //users/10

package main

import "github.com/gin-gonic/gin"

func main2() {

	r := gin.Default()

	r.GET("/users/:id" , func(c *gin.Context) {

		id := c.Param("id")

		c.JSON(200 , gin.H{
			"user_id :" : id,
		})
	})

	r.Run(":8080")
}