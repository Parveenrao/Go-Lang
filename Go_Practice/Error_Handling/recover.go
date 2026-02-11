/* 
   => Recover --> catches a panic and lets the program to finish 
   => It works in the defer function 
*/
// Panic with recover 

package main
import "fmt"


/*



func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("boom")

	fmt.Println("This will NOT run")
}

*/



// Recover inside function 

func printBad() {
	panic("Something bad")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in main:", err)
		}
	}()

	printBad()

	fmt.Println("Program continues")
}
