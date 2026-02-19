/* 
   => When  your program is stopping ( ctrl + c / server kill )
      you dont exist immediately 

	You:
	1. Stop accepting new work
	2. finish current job 
	3. close resources 
	4. exit cleanly
	
=> Why this matters in backend

Imagine:

users are sending requests

workers processing jobs

DB connections open

You press Ctrl+C.

Without graceful shutdown → everything dies instantly.

With graceful shutdown → Go politely says:

“No new work. Finish current. Then exit.”

*/

// 1. catch os signal (ctrl + c)

package main

/*
import ("fmt"
        "os"
	    "os/signal")


func main() {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop , os.Interrupt)

	
	fmt.Println("App running... press Ctrl+C")

	<-stop // WAIT here

	fmt.Println("Shutting down gracefully...")
}

		

*/