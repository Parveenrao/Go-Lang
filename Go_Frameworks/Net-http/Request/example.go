package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method:", r.Method)
	log.Println("Path:", r.URL.Path)
	log.Println("Name:", r.URL.Query().Get("name"))
	log.Println("Header (User-Agent):", r.Header.Get("User-Agent"))

	w.Write([]byte("Check terminal logs"))
}

func main() {

	http.HandleFunc("/hello", handler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
