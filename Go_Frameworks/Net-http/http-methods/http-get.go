// GET /users
// GET /users?id=1
package main



/*


import ("fmt"
        "net/http"
	     "strconv")

func main() {
	http.HandleFunc("/hello" , func(w http.ResponseWriter , r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w , "Only get allowed"  , http.StatusMethodNotAllowed)

			return 
		}

		fmt.Fprintln(w, "Hello from get request")
	})

	http.ListenAndServe(":8080" , nil)
}

*/

//------------------------------------------------------------------------------------------------
// Get with query parameter     
// /user? name = parveen&age = 22


/*


func helloHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Query parameters
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// Response
	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

*/


//--------------------------------------------------------------------------------------------

// Get Returing json 

/*

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getuser(w http.ResponseWriter , r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w , "Only Get method allowed" , http.StatusMethodNotAllowed)
		return 
	}

	users := []User {

		{1 , "Parveen"},
		{2 , "Yadav"},
	}

	w.Header().Set("COntent-Type" , "applicaton/json")            // tells client that i am sending json data
	json.NewEncoder(w).Encode(users)                              // convert user slice into json , Result sent to the client
}


func main() {

	http.HandleFunc("/users" , getuser)
	http.ListenAndServe(":8080" , nil)
}



//----------------------------------------------------------------------------------------------


// Get with map in memory database

var user = map[int]string {
	1: "Parveen",
	2: "Yadav",
}


func getusr(w http.ResponseWriter , r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w , "Method NOt allowed"  , http.StatusMethodNotAllowed)
		return

	}

	idStr := r.URL.Query().Get("id")

	if idStr == ""{
		http.Error(w , "Id Required" , http.StatusBadRequest)
		return
	}


	id, err := strconv.Atoi(idStr)
    if err != nil {
	    http.Error(w, "Invalid Id", http.StatusBadRequest)
	    return
    }




           // convert string into int


	name, ok := user[id]

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "User: %s", name)

}


func main() {
	http.HandleFunc("/users" , getusr)
	http.ListenAndServe(":8080" , nil)
}

*/