/* 
   Only one Goroutine is allowed inside the code at the time 


   => Imagine two go rountine both do 
      count ++ 
	  
	  --> Result 
	     unpredicatable 
		 Race condition
		 worong number   
      


*/


// With mutex 


package main 

/*
import ("fmt"
        
	     "time")


var count = 0

func increment() {
	for i := 0; i<1000;  i++ {
		count ++
	}
}


func main() {

	go increment()
	go increment()

	time.Sleep(time.Second)

	fmt.Println("Final count:", count)
}



// Now with mutex 


import ("fmt"
        "sync"
        "time")

var count = 0 

var mu sync.Mutex

func increment() {

	for i := 0; i <1000; i++ {

		mu.Lock()

		count++

		mu.Unlock()
	}
}


func main() {

	go increment()
	go increment()

	time.Sleep(time.Second)

	fmt.Println("Final count:", count)
}

*/