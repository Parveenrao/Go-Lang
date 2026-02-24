package main
/*
import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


func main() {

	r := gin.Default()


	r.POST("/login" , func(c *gin.Context) {

		var input LoginInput

		var hashedpassword string 

		var userID int 


		// Read json 

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400 , gin.H{
				"error" : err.Error(),
				
			})
			return
		}

		// get id and hashedpassword from db 

		err := db.QueryRow("Select id , password FROM users WHERE email = $1" , input.Email).Scan(&userID , &hashedpassword)

		if err == sql.ErrNoRows {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	  }

	   if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	  }

	   // Compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "wrong password"})
		return
	}

	// SUCCESS
	c.JSON(200, gin.H{
		"message": "login successful",
		"user_id": userID,
	})
})





}


*/