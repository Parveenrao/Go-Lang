package main 


/* 
  => A context is a shared stop signal + data carrier 


  lets say to many goroutines 

  Stop now 
  You have 2 second 
  This Requst is cancelled




import ("fmt"
       "time"
	   "context")


func main() {

	ctx , cancel := context.WithTimeout(context.Background() , 2 * time.Second)

	defer cancel()

	select {

		case <- time.After(5 * time.Second): {
			fmt.Println("Work finished")
		}

	case <-ctx.Done():
		fmt.Println("cancelled" , ctx.Err())
	}
}

*/