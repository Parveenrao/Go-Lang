
//  Query parameters live after ? in the URL.       /add?a=10&b=20






package main

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
)


func main1() {

	r := gin.Default()

	r.GET("/search" , func(c *gin.Context) {

		name := c.Query("name")

		c.JSON(http.StatusOK , gin.H{

			"name" : name,
		})


	})

	r.Run(":8080")
}