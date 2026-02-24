package main

import (
	"log"
	"net/http"
	"time"
	"context"
)


func handler(w http.ResponseWriter , r *http.Request) {

	ctx := r.Context()

	select {

	case <-ctx.Done():
		log.Println("Client left")
		return

    default:
		log.Println("Client still connected")		


	}

	
	w.Write([]byte("Hello"))
}



// Timeout 

func handler1(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		w.Write([]byte("Finished"))
	case <-ctx.Done():
		http.Error(w, "Timeout", 408)
	}
}


// carrry data / pass value from middleware 

func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "user", "parveen")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
