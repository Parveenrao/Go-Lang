// Go doest not do automatically type conversion  , we have to do manually type conversion 

package main

import "fmt"



/* 
a = 10 
b = 2.0

c:= a + b  error 

*/

func main(){

	a := 10 
	b := 2.5 

	c:= float64(a) + b

	fmt.Println(c)


	// float to int 

	pi:= 3.15

	n := int(pi)

	fmt.Println(n)

	// str := string(65)   // gives ASCII 'A' not allowed


}

