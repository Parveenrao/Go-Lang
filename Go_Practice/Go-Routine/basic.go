// Go - Routine is a light thread managed by go runtine , not by os 

package main

/* 

=> a function running concurrently with other functions 

   go somefunctions()

   Add go --> function become concurrent

=> why goroutines are special 
   
   1. Os thread 
      => heavy (MBs of stack)
	  => Slow creation 
	  => Limited creation (~ thousand)
   
   2. Goroutines
      => Start with ~ 2Kb stack
	  => Grow / shrink automatically
	  => Millions possible 
	  => Super fast to create 
	  
	  so go app can handle massive concurrency 
*--------------------------------------------------------------------------------

=> Internal Model 
   
   Go use GMP scheduler 
   G---> Goroutine 
   M---> Machine (Os thread)
   p---> Processor (logical core)

G (your function)
↓
scheduled on P
↓
executed by M

Go runtime maps many goroutines onto few OS threads.

This is why Go scales so well.



====> First program 

import "fmt"

func hello() {
     fmt.Println("Hello Go_Routine")
	 }

func main() {
     go hello()   // Output -> nothing because main() exist before goroutine runs
	 }           

   
   



*/