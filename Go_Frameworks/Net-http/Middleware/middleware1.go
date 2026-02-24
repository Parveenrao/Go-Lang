// Smallest possible middleware

package main

import "net/http"



func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// BEFORE handler
		println("Middleware before")

		next.ServeHTTP(w, r) // call real handler

		// AFTER handler
		println("Middleware after")
	})
}


// Takes a handler (next)  and return new handler

// next.SrveHttp ( It forwards the request to the actual endpoint handler.  If you REMOVE this line → your API will never reach the real handler.)

