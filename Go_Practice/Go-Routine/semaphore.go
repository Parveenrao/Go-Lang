package main

/* 
  Semaphore means only N go routines are allowwed at once 
  Basically semaphore is used to limit concurrency 


  => Example ---Parking slot 

    3 slots 
	10 cars 
	only 3 inside 
	3 can wait 

  => Withour goroutine 
  
  for i :=0; i <1000; i++{
      
      go work(i)}

	  100 Go routine create instantly ---> server dies 





// We use buffered channel as Semaphore 



import ("fmt"
        "time")


func main() {

	sem := make(chan struct{} , 3)  // only three allowed


	for i := 0; i < 5; i++ {

		sem <- struct{}{}  // entry slot (struct{} --> struct intilzation )  struct{}{} --> emtpy struct

		go func(id int) {

			fmt.Println("Start" , id)

			time.Sleep(2 *time.Second)

			fmt.Println("End" , id)

			<- sem     // exit slot 
		}(i)
	}

	time.Sleep(8 * time.Second)
}	

*/

