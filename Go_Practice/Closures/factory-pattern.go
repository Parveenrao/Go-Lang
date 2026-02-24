// Clsoures as factory pattern 

package main 

import "fmt"

func prefixer( prefix string) func(string) string  {

	return func (s string) string  { 
		return  prefix + s
		
	}
}

func main()  {

	info := prefixer("INFO: ")
	errorlog := prefixer("ERROR :")

	fmt.Println(info("Server started"))
	fmt.Println(errorlog("connection failed"))
	
}