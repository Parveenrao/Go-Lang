package main

import "fmt"

func main() {

	// 1. Create array 

	numbers:= [5]int{10 ,20 ,30, 40 ,50}

	fmt.Println("Original Array :" , numbers)

	// 2. Access elements of array 

	fmt.Println("First Element :" , numbers[0])
	fmt.Println("Third Element :" , numbers[2])

	// 3. Modify ELement 

	numbers[0] = 15
	fmt.Println("After modification array  :" , numbers)

	// 4. Lenght of array 

	fmt.Println("Length of array :" , len(numbers))

	// 5. Loop through array 

	for i:=0; i < len(numbers); i++{
		fmt.Println(numbers[i])
	}

	// 6. Range loop
    fmt.Println("Using range:")
    for index, value := range numbers {
        fmt.Println(index, value)
    }


	  // 7. Calculate sum
    sum := 0
    for _, val := range numbers {
        sum += val
    }

    fmt.Println("Sum of elements:", sum)
}




