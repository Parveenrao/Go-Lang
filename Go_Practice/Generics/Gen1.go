package main 

import "fmt"


func Print[T any](v T) {
	fmt.Println(v)
}

// Generic function with return 

func Identity[T any](x T) T{
	return  x
}


// Constraint --> You cant do math On Generics 
/*
func addition[T any](a , b T) T{
	return  a + b
}  
	
*/

// Built in Constraint Comparable 

func Equal[T comparable] (a , b T) bool {
	return  a == b
}

// Generic Struct 

type Box[T any] struct {
	value T
}

func main(){
	Print(10)
	Print("Go")
	Print(3.14)

	a := Identity(10)
	b := Identity("Go Run Go")

	fmt.Println(a)
	fmt.Println(b)

	c := Equal(1, 2)
	fmt.Println(c)


	b1 := Box[int]{value: 10}
	b2 := Box[int]{value: 20}

	fmt.Println(b1)
	fmt.Println(b2)
}


