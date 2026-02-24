// Logging Middleware --> Log Every Request 


package main
/*
import ("log"
        "net/http")


func loggingmiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func (w http.ResponseWriter , r *http.Request) {

		log.Println(r.Method , r.URL.Path)

		next.ServeHTTP(w ,r)
	})
}		


func main () {
	http.HandleFunc("/users" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Println("Server started on port : 8080")

	http.ListenAndServe(":8080" , loggingmiddleware(http.DefaultServeMux))
}

*/