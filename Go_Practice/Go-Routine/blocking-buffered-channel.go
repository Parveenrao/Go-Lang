package main
/*
import ("fmt"
       "time")


func main(){

	ch:= make(chan int , 2)

	go func() {

		// Prodcuer - send fast  buffer absorb first two

		for i := 1; i <= 5; i++{
			fmt.Printf("Sending %d...\n", i)
            ch <- i  // blocks on 3rd, 4th, 5th until space available
            fmt.Printf("Sent %d\n", i)
		}
		close(ch)
	}()

	// consumer -- read slowly 

	time.Sleep(1 * time.Second)        // let consumer fill buffer first

	for val := range ch{
		fmt.Println("Received" , val)

		time.Sleep(100 * time.Millisecond)
	}
}	   

*/