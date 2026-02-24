package main 

import ("encoding/json"
       "log"
	    "net/http")


type User struct {

	Name string `json:"name"`
	Age  int    `json:"age"`
}


func handler(w http.ResponseWriter , r *http.Request) {

	// Header 
	w.Header().Set("content-Type" , "application/json")

	//json 
	w.WriteHeader(http.StatusOK)

	//Body 

	user := User{
		Name : "Parveen",
		Age : 22,
	}

	json.NewEncoder(w).Encode(user)

}


func main() {

	http.HandleFunc("/user", handler)

	log.Println("Server running :8080")
	http.ListenAndServe(":8080", nil)
}