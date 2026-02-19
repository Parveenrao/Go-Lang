// A channel let goroutine talk to each other
package main
/*
import "fmt"



=> Channel

   A channel is a typed pipe that lets goroutine send data to each other safely

   Goroutine --> workers
   Channel --> conyer belt

   *** Communicate by share memory

   // create channel

   ch :=make(chan int) ->channel that carries int

   ch <- 10 (send)

   x:= <- ch(Recieve)



func worker(id int , ch chan int) {
	ch <- id *2
}

func main()  {

	ch := make(chan int)      // channel created sender wait for recieve and reciever wait for sender (unbuffered)

 
	for i:=1; i<=3; i++{
		go worker(i , ch)
	}

	for i:=1; i<=3; i++{
		fmt.Println(<-ch)
	}
	
}   */