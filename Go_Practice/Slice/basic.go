package main

/* 
=> Slice --> Dynamic Array (dynamic collection of same type of elements)
   
   Built on top of array 
   size can grow/shrink
   reference type

=> Internal Structure of the slice 

   1. Pointer to array 
   2. Lenght
   3. Capacity 

=> Creating Slice 
   
nums:= []int{1, 2, 3}

// using make    nums:= make([] 3, 5)


=> From array 
   
a:= [5]int{1, 2, 3, 4, 5}
b:= a[1:4]


=> Memory behaviour 

SLice share same memory --> so changing one affect others(unless reallocated)




*/