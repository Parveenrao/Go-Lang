
// Reading from map 

package main 

import "fmt"

func main() {
	m := make(map[string]int)

	m["go"] = 10
	m["C++"] = 100
	m["Scala"] = 40 


	v := m["go"]


	fmt.Println(v)

	// if key doesnt exist --> return 0 for missing keys ( int --> 0 , string --> "" , bool --> false)

	z := m["python"]

	fmt.Println(z)


	//checking if exist 

	val , ok := m["Rust"]

	if ok {
		fmt.Println("Found" , val)
	}else {
		fmt.Println("Not Found")
	}


	// len of map 

	fmt.Println(len(m))


	// Iterate over map 

	// order is random

	for k , v := range m {
		fmt.Println(k , v)
	}

}