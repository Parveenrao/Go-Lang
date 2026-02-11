/*
  => Panic means something happened that the program  does not know  how to contiue from
  
    when panic happens  --> current function stops 
	program crashes
*/


// Example panic at app startup / db cofig 

package main 

/*
import "log"


func connectDB() error {
	return  nil                //imagine db failed here
}


func main () {
	err := connectDB()

	if err != nil{
		panic("database connection failed")
	}

	log.Println("App started")
}

*/