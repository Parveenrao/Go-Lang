/*
  => Authentication ---> Means who are u
  => Authorization --> what are you allowed to do


     We already learn authentication , now we learn how to add persmission


🎯 Goal

We will build:

✅ Role stored in DB
✅ Role added to JWT
✅ Middleware checks role
✅ Admin-only route

No theory — practical.
*/

package main
/* 
import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generatetoken(userID int , role string ) (string , error) {

	claims := jwt.MapClaims{
		"user_id" : userID,
		"role" : role,
		"expire" : time.Now().Add(24 * time.Hour).Unix(),
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)

	return token.SigningString(jwt_key)

	// Now jwt contain role, userid and expire 

	// Fetch role during login

	var role sting 

	err := db.QueryRow("select id , password , role from users where email = $1" , input.Email).Scan(&userID , &hashedpassword ,&role)


	token, _ := generateToken(userID, role)




   claims := token.Claims.(jwt.MapClaims)

   userID := int(claims["user_id"].(float64))
   role := claims["role"].(string)

   c.Set("user_id", userID)
   c.Set("role", role)


}


   func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, _ := c.Get("role")

		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "admin only"})
			return
		}

		c.Next()
	}
}


*/