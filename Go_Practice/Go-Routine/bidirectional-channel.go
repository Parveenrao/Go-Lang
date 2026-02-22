/*
  Every channel is bidirectional
  You can send and receive both on same channel

*/

package main
/*
import (
	"fmt"
	
	"time"
)


func pinpong(ch chan string) {     // receive bidirection channel 

	msg := <- ch

	fmt.Println("Got :" , msg)

	ch <- "pong"

}

func main(){

	ch := make(chan string , 1) 

	ch <- "pong"

	go pinpong(ch)

	time.Sleep(100 * time.Millisecond)

	reply := <-ch

	fmt.Println("Reply :" , reply)


}

*/