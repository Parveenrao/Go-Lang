package main 

/* Select => Select block until one channel is ready 
   
   => If both are ready then Go pick one randomly
   
    => think select like this wait until any channel is ready , then execute that case 
	




import (
	"fmt"
    "time"
)


func main() {

	ch1:= make(chan string)
	ch2:= make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		ch1 <- "message from channel 1"
		
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from channel 2"
		
	}()

	// multiple

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)

	
	case msg2:= <- ch2:
		fmt.Println(msg2)
	
    default:
		fmt.Println("Nothing is ready")		
	}


	} */