package main
 
import "fmt"

func main ()  {

	// 1 var with type 
	var a  int = 10

	// var without type 
	var b = 20 

	// short declaration 
	c:= 30 

	// multiple variable 

	var d ,  e int = 40 ,50 
	
	fmt.Println(a , b, c, d,e)
	
}


// Interview question  => if we dont use variable it will show compilation error


func main1() { 

	// x:= 20    

	// compilation error beacuse it is not used
}


// String Variable 



func main2() {

    name := "Parveen"
    language := "Go"

    fmt.Println(name)
    fmt.Println(language)
}


// Float variable 


func main3() {

    price := 99.50      // Go automatically makes this float64
    pi := 3.14159

    fmt.Println(price)
    fmt.Println(pi)
}
