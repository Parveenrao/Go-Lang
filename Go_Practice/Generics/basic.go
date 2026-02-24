package main

/* 

Before generics , if you wanted the same logic for int, float64 , string 

we write 

func sumint(a []int) int{}

func sumfloat(a float64[]) float64  

=> thise cause    
    1. Code duplication 
	2. Hard to maintain
	3. Ugly api
*---------------------------------------------------------------------------

=> Generics ---> Generics allow types as parameter 

like func (a int , b int )

-> Generics allow you to write 

func add[t any](a T , b T)  -> T as a parameter   , T can be any type 

=> Basic syntax   
   
   func Print[T any](value T) {
    fmt.Println(value)
}







*/