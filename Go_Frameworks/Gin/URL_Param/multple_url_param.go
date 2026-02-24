// Multiple url Path --- >users/id/post/id

package main

import "github.com/gin-gonic/gin"

func main1() {

	r := gin.Default()

	r.GET("/users/:id/post/:postId" , func(c *gin.Context) {

		userId := c.Param("id")
		postId := c.Param("postId")


		c.JSON(200 , gin.H{

			"user_ID " : userId,
			"post_ID" : postId,
		})
	})

	r.Run(":8080")
}