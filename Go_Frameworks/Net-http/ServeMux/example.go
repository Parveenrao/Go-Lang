package main

import ("log"
       "net/http")

func hello(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("hello"))
}	   

func upload(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("upload page"))
}


func main() {

	mux := http.NewServeMux()

	// Register Route

	mux.HandleFunc("/hello" , hello)

	mux.HandleFunc("/upload" , upload)

	log.Println("Server running :8080")

	// Use YOUR mux
	http.ListenAndServe(":8080", mux)
}