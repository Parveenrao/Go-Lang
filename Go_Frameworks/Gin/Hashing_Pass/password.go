package main 

import ("fmt"
        "golang.org/x/crypto/bcrypt")

func main() {

	password := "12345"

	hash , _ := bcrypt.GenerateFromPassword([]byte(password) , 10)

	fmt.Println("Hash :" , string(hash))


	err := bcrypt.CompareHashAndPassword(hash, []byte("12345"))


	if err != nil {
		fmt.Println("Wrong Password")
	}else {

		fmt.Println("Correct password")
	}
}
