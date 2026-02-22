/*  A buffered channel has capacity N 
    Sender doesnt block until buffer is full 
	Reciever does not block until buffer is empty 
	They are decoupled

*/


package main

/*
import ("fmt"
        )


func main() {

	ch := make(chan int , 3)

	// send without receiver , doesn't block buffer absorb

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Sending 3 values no reciever neended")
	fmt.Println("len  :" , len(ch) , "cap :" , cap(ch))



	// receive - FIFO order 

	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	fmt.Println("len" , len(ch))
}		


*/