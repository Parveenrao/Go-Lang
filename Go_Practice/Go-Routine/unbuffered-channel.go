/* An unbufferd channel has zereo capacity 
   Send and receive happen at the same time 
*/

package main

/*
import ("fmt"
        "time")


func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("Sender : waiting to send..." )

		ch <- 100 // block  here until someone receivess

		fmt.Println("Sender  value was taken  , continuing")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Receiver : about to receive..")
	val := <-ch

	fmt.Println("Receiver : got" , val)
	time.Sleep( 100 * time.Millisecond)
}

// Send and receiving in same go routine will cause deadlock

*/