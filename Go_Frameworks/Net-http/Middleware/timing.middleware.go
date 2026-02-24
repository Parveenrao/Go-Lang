/*
👉 when the request starts
👉 when the handler finishes
👉 logs the total duration


*/

package main 

import ("log"
        "net/http"
	     "time")

func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now() // start time

		next.ServeHTTP(w, r) // real handler

		duration := time.Since(start) // total time

		log.Printf("%s %s took %v\n", r.Method, r.URL.Path, duration)
	})
}



