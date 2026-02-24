// write  a go proram that ask my name  , age ,salary and print all 

package main

import "fmt"

func main() {

	var name string
	var age int 
	var salary float64


	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Enter your age :")
	fmt.Scanln(&age)

	fmt.Println("Enter your salary:")
	fmt.Scanln(&salary)

	fmt.Println("Hello Mr" , name , "Age is => " , age   , "Salary is=>" ,salary)
}
