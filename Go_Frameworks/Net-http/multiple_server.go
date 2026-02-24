package main 

import ("fmt"
        "net/http")


func home(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w  , "Home Page")
}		

func about(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintln(w , "About page")
}


func main() {
	http.HandleFunc("/" , home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080" , nil)
}
