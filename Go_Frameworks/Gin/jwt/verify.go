/*
   => JWT generated on login
       but server still doesn't know who is calling protected API


   => Now

   Read jwt from request header
   verify signature
   decode payload
   extract user_id
   set user_id to to request context

*/

package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	
)


func AuthMiddleware() gin.HandlerFunc {

	return  func(c *gin.Context) {


		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401 , gin.H{"error" : "missing token"})
			return 
		}

		tokenstring := strings.TrimPrefix(authHeader , "Bearer")

		token , err := jwt.Parse(tokenstring , func(t *jwt.Token) (interface{} , error) {
			return  jwt_key , nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := int(claims["user_id"].(float64))

		// store user id in Gin context
		c.Set("user_id", userID)

		c.Next()
		
     
	}
	

	
}



// Step 2 — Use user_id in protected routes

// Now inside ANY protected route:

func hello() {

	r := gin.Default()
       
	r.GET("/profile", AuthMiddleware(), func(c *gin.Context) {

	userID, _ := c.Get("user_id")

	c.JSON(200, gin.H{
		"message": "authenticated",
		"user_id": userID,
	})
})

}