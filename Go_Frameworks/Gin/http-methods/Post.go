package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`      // binding  --> now if field are missig - Gin give auto erro
	Age  int    `json:"age" binding:"required"`
}



func main() {

	r := gin.Default()

	r.POST("/users" , func(c *gin.Context) {

		var user User

		// Bind Json Body 

		if err := c.ShouldBindJSON(&user); err != nil {

			c.JSON(http.StatusBadRequest , gin.H{
				"error" : err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK , gin.H{
			"name" : user.Name,
			"age" : user.Age,
		})


	})

	r.Run(":8080")
}
