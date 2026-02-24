// Recovery middleware prevent your Go Server crashing when a panic happens 


// Without Recovery --> Server crash , ALl client disconnected , Production outrage

// With Recovery --> Call a defer function  -- call recover  --- log panic --> give error 500 --keep server alive 


package main

import ("log"
         "net/http"
		 "time")


func loggingmiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request" , r.Method , r.URL.Path)

		next.ServeHTTP(w ,r)
	})
}		 

func timingmiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w , r)

		log.Println("Time:", time.Since(start))
	})
}


func recoverymiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func ()  {

			if err := recover(); err != nil {

				log.Println("Panic :" , err)
				
				http.Error(w, "Internal Server Error", 500)

			}
			
		}()

		next.ServeHTTP(w , r)
	})
}


func main() {
	http.HandleFunc("/hello" , func(w http.ResponseWriter , r *http.Request) {
		
		panic("boom")
	})
	
	handler := recoverymiddleware(loggingmiddleware(timingmiddleware(http.DefaultServeMux),
	), 
    )


	log.Println("Server Running : 8080")

	http.ListenAndServe(":8080" , handler)
}
