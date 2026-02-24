// Variadic Function => a variadic function can take 1 , 2 or many arugments of same type 

// Basic Syntax    func myFunc(nums .. int){}   --> nums is  slice [] int

package main 

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _ , n:= range nums {
		total += n
	}
	return  total
}

func main(){
	fmt.Println(sum(1 ,2))
	fmt.Println(sum(1,3,4,5))
	fmt.Println(sum(1,2,3,4,5))
}