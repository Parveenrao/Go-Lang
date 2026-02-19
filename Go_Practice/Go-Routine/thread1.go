package main

/*
import ("fmt"
         "sync"
		 "time")

func downlaod(id int , wg *sync.WaitGroup) {        // Pointer to waitgroup --> all goroutine share same waitgroup
	defer wg.Done()             // tell waitgroup this goroutine is done

	time.Sleep(500 * time.Microsecond)

	fmt.Println("Downloaded file" , id)

}

func  main()  {

	var wg sync.WaitGroup

	for i:=1; i<=3; i+={
		go downlaod(i &wg)
	}

	wg.Wait()

	
	fmt.Println("All downloads completed")
	
}



*/