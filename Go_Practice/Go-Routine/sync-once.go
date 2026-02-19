// There is one method simpler than Go-Routine 

// How i make sure that something will run only one time , even with many go routines

// Sync.once --> gurantees it  (Code runs exactly once . Never twice , even if 100 100 go routines call it)

package main


/*
import ("fmt"
         "sync")


var once sync.Once


func initapp(){
	fmt.Println("App intilization ")
}

func worker(id int) {

	fmt.Println("Worker calling" , id , "calling init")
	once.Do(initapp)

}

func main () {

	for i := 0; i <100; i++{
		go worker(i)
	}

	
}

*/