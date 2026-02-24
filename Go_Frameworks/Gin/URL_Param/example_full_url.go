package main

import ("github.com/gin-gonic/gin"
        "net/http"
	    "strconv")

func main() {

	r := gin.Default()


	r.GET("multiply/:a/:b" , func(c *gin.Context) {

		aStr := c.Param("a")
		bStr := c.Param("b")

		a , err1 := strconv.Atoi(aStr)

		b , err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {

			c.JSON(http.StatusBadRequest , gin.H{
				"error" : "both a and b should be number",
			})

			return
		}

		c.JSON(http.StatusOK , gin.H{
			"a" : a,
			"b" : b,
			"result" : a * b,

		})


	})

	r.Run(":8080")
}		



