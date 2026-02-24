// Basic input using fmt.scanln 



package main

import  "fmt"

func hello() {

	var name string

	fmt.Println("Enter your name")
	fmt.Scanln(&name)              // address of the variable => Go puts user variable into that variable 

	fmt.Println("Hello" , name)

	// Taking integer input 


	var age int 

	fmt.Println("Enter your age")
	fmt.Scanln(&age)

	fmt.Println("Parveen age =>" , age)

	// Taking float input 

	var price float64

    fmt.Print("Enter price: ")
    fmt.Scanln(&price)

    fmt.Println(price)


	// Multiple input at once 

	var a int 
	var b int 

	fmt.Println("Enter two numbers :")
	fmt.Scanln(&a , &b)

	fmt.Println(a + b)


}                                         
