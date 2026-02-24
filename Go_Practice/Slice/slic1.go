package main

import "fmt"

func main()  {

	// 1. Create slice 
	nums := []int{10 ,20 ,30, 40, 50}
	fmt.Println("Initial Array" , nums)

	// 2. Lenght and Capacity 
	fmt.Println("Length  Array" , len(nums))
    fmt.Println("Capacity Array" , cap(nums))

	// 3. Access ELement 
	fmt.Println("First element of Array" , nums[0])

	// 4. Modify element 
	nums[0] = 15
	fmt.Println("After modification array" , nums)

	// 5. Append ELement 

	nums = append(nums, 60)
	nums = append(nums, 70 ,80)
	fmt.Println("After append Array :" , nums)

	// 6. Remove element at index 2
    index := 2
    nums = append(nums[:index], nums[index+1:]...)
    fmt.Println("After remove index 2:", nums)

	// 7. Loop Slice 
	fmt.Println("Loopoing")
	for i , v := range nums{
		  fmt.Println(i , v)
	}

	// 8. Create slice with make     make([]int, length, capacity)  , length -> how mmanu elements i will start with , capacityhow much memory

	arr := make([]int, 1, 2)
	fmt.Println("Array using make" , arr)
	fmt.Println("Len:" , len(arr) , "Cap:" , cap(arr))

	// 9. Append to make slice
	arr = append(arr, 100)
	fmt.Println("Array using arr1:" , arr)

	// 10. Slice from slice
    part := nums[1:3]
    fmt.Println("Sub slice:", part)

	// 11. Memory sharing demo
    part[0] = 999
    fmt.Println("Nums after sub-slice change:", nums)

    // 12. Sum all values
    sum := 0
    for _, v := range nums {
        sum += v
    }
    fmt.Println("Sum:", sum)
}





	





	
