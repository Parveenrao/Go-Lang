// Custom error -- In real system you need error code , type  ,context  , machine-readable info

// Stage 1 Custom Error --> invalid age 

package main 

/*

import ("fmt"
        )

type AgeError struct {
	Age int 
}		


func (e AgeError) Error() string {

	return  fmt.Sprintf("Invalid age : %d" , e.Age)

}

func validateage(age int) error {
	if age < 18 {
		return  AgeError{Age: age}
	}

	return nil

}


func main() {

	err := validateage(15)

	if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Println("Valid age")
}


*/