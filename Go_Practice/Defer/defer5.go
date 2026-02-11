// Defer With Closure 

// Closure will see the latest value 


package main

import "fmt"

func main() {
	x := 10

	defer func() {
		fmt.Println(x)
	}()


	x = 20
}



// Work even with panic
