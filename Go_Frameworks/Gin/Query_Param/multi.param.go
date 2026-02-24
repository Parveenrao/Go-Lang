package main

import (
	"net/http"
	"strconv"
    "github.com/gin-gonic/gin"
)



func main() {

	r := gin.Default()

	r.GET("/add" , func(c *gin.Context) {

		aStr := c.Query("a")
		bStr := c.Query("b")


		a , err1 := strconv.Atoi(aStr)
		b , err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {

			c.JSON(http.StatusBadRequest , gin.H{

				"error" : "both a and b should be numbers",
			})
		}

		c.JSON(http.StatusOK , gin.H{

			"a" : a,
			"b" : b,
			"addition" : a + b,
		})


	})

	r.Run(":8080")
}