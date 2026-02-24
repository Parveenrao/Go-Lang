// If-else in go 

package main

import "fmt"

func main()  {

	var name string
	var marks int 
	
	fmt.Println("Enter your name : " , name)
	fmt.Scanln(&name)


	fmt.Println("Enter your grades :")
	fmt.Scanln(&marks)

	if marks>= 90{
		fmt.Println(name  , "Grade A")
	} else if marks >= 70{
		fmt.Println(name  , "Grade B")
    } else if marks >= 50{
		fmt.Println(name , "Grade C")
	} else {
		fmt.Println(name , "Failed")
	}
	
}