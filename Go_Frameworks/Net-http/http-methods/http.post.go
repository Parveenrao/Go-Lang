// Post Method 

// Basic Post Handler 

package main 

/*


import ("fmt"
        "net/http"
	     "encoding/json")


func main () {

	http.HandleFunc("/create" , func(w http.ResponseWriter , r *http.Request)  {

		if r.Method != http.MethodPost {
			http.Error(w , "Only Post Allowed" , http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintln(w ,"Post Recieved")
	})

	http.ListenAndServe(":8080" , nil)
}



//-----------------------------------------------------------------------------------------

// Post with json body 


type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func createuser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w , "Only Post Allowed" , http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w , "Invalid json" , http.StatusBadGateway)

		return 

	}

	fmt.Println("Received:", user)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created")

}


func main() {
	http.HandleFunc("/users", createuser)
	http.ListenAndServe(":8080", nil)
}

*/