package main


/* 
  => Rmutex is a special mutex that allows  -- many reader at the same time 


  => Imagine 

    100 goroutines reading data
	1 goroutines update data


	with normal mutex everyone waits one by one


	with Rmutex 

	all readers run together   -- writer wait until reader finish




import ("fmt"
        "sync"
	    "time")


var count = 0

var rw sync.RWMutex


func reader(id int) {

	rw.RLock()  // Read lock 

	fmt.Println("Reader" , id , "counts" , count)

	time.Sleep(time.Second)

	rw.Unlock()


}

func writer(id int ){

	rw.Lock()              // Write lock

	fmt.Println("Writer" , id , "writing")

	time.Sleep(time.Second)

	rw.Unlock()
}


func main() {

	go reader(1)
	go reader(2)
	go reader(3)


	time.Sleep(500 * time.Millisecond)

	go writer(1)

	time.Sleep(5 * time.Second)
}

*/