// Select let one go routine listen on multiple channel simultaneously -- whoever send first wins 

package main 

import ("fmt"
         "time")


func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)


	go func() {
		time.Sleep(1 * time.Second)

		ch1 <- "msg from 1"
	}()

	go func() {

		time.Sleep((500 * time.Millisecond))

		ch2 <- "msg from 2"
	}()







	select{

	case msg1 := <- ch1:
		fmt.Println(msg1)

	case msg2 := <- ch2:
		fmt.Println(msg2)	
	}



}		 




/* 
 
select pick randomly , to make channel priority use select with priority trick

func main() {
    highPriority := make(chan string, 5)
    lowPriority  := make(chan string, 5)

    highPriority <- "urgent 1"
    highPriority <- "urgent 2"
    lowPriority  <- "normal 1"

    for {
        // drain high priority FIRST before touching low priority
        select {
        case msg := <-highPriority:
            fmt.Println("HIGH:", msg)
            continue
        default:
        }

        // now handle low priority
        select {
        case msg := <-highPriority:
            fmt.Println("HIGH:", msg)
        case msg := <-lowPriority:
            fmt.Println("LOW:", msg)
        }
    }
}


*/