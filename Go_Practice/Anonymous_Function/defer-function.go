// Defer with anonymous function 


package main

import "fmt"

func main () {

	defer func() {
		fmt.Println("Deferred anonymous function")
	}()

	fmt.Println("Main running")
}