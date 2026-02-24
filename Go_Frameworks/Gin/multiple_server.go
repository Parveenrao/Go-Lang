package main

import "github.com/gin-gonic/gin"

func main1() {

	r := gin.Default() 

	// Home
	r.GET("/home", func(c *gin.Context) {
		c.String(200, "Home Page")
	})


	// about
	r.GET("/about", func(c *gin.Context) {
		c.String(200, "About Page")
	})

	// contact
	r.GET("/contact", func(c *gin.Context) {
		c.String(200, "Contact Page")
	})

	r.Run(":8080")
}

