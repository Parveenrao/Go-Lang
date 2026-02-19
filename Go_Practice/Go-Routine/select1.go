package main

/*  Write a prgram 
    => Two channels 
	=> One sends every 400ms 
	=> Second send every 700ms
	=> Print 5 msh in total 


import ("fmt"
        "time"
	)


func main(){

	ch1:= make(chan string)
	ch2:= make(chan string)


	//Producer (400ms)
	go func() {
		for {
			time.Sleep(400 * time.Millisecond)
			ch1 <- "From channel 1 (400ms)"
		}
		
	}()

	// Producer 2 ( 700ms)

	go func()  {

		for {
			time.Sleep(700 * time.Microsecond)
			ch2 <-  "From channel 2 (700s)"
		
		}

		    
	}()


	// Recieve only 5 msg total 

	for i:=0; i<5; i++{ 
		select{

	    case msg1 := <- ch1:
			fmt.Println(msg1)
			
	    case msg2 := <- ch2:
			fmt.Println(msg2)
			
	   
		}
	
}	
    fmt.Println("Done")

}

*/