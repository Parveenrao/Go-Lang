package main 

/*
import ("log"
        "time"
	     "net/http")
		
		 
func loggingmiddleware1(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request :" , r.Method , r.URL.Path)

		next.ServeHTTP(w,r)
    })
} 


func timemiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)


		log.Printf("Time taken: %v\n", time.Since(start))

		
	})
}


func main() {

	http.HandleFunc("/home" , func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(500 * time.Millisecond) // simulate work
		        w.Write([]byte("Hello"))
	})

    // chain middleware 

	handler := loggingmiddleware1(timemiddleware(http.DefaultServeMux))

	http.ListenAndServe(":8080" , handler)

}

*/