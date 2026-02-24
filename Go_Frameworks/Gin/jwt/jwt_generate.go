package main 

import("fmt" 
       "time"
	    "github.com/golang-jwt/jwt/v5")


var jwt_key = []byte("secretkey")

func generatejwttoken(userID int) (string , error) {
	
	claims := jwt.MapClaims{              // data stored inside jwt key -value pair 
 
		"user_id" : userID,              // this token belongs to user 1
		"expire" : time.Now().Add(24 * time.Hour).Unix(),   // unix -- convert to unix timestamp , jwt require expiration in unix timestam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)       // singing algo , creat jwt in memory 

	return  token.SignedString(jwt_key)                               // now signed and convert into string

}


func main () {

	token , err := generatejwttoken(1)

	if err!= nil {
		panic(err)
	}

		fmt.Println("JWT Token:")
	fmt.Println(token)
}
