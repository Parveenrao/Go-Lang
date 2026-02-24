// Array in go => Array are the collections of the fixed elements 
// All elements are of same data type 
// stored in continuous memory 


package main 



/* [10 , 20 ,30 , 40]

   index : 0  1  2  3  4
   value : 10 20 30 40 50 

=> Array  lets your store multiple values under one variable 
   a:= 10
   b:= 20
   c:= 30

   nums:= [2]{10 ,20 30}


=> Fixed size
   nums:= [3]int{1 ,2, 3}

   if you dont intialized 

   nums:= [3]int => {0 , 0 ,0 }


=> Memory types 
   Array are value types
   Copied when passed to function 
   Copied when assigned

   a:= [3]int{10 ,20 ,30}
   b:=a

   a and b are different array 

   changing b doest not effect a

//----------------------------------------------------------

=> Array are stored on the stack - but sometime on the heap 

=> Stack 
   1. Fast
   2. Automatic cleanup
   3. Used for small and short lived values 
   4. Function - local variable goes to the stack

=> Heap 
   1. Slow
   2. Manage by garbage collector 
   3. Used when data must live longer or be shared
   
   

   Large array go to heap 



   
 
   */ 