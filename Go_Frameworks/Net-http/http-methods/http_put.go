package main 

/*
import ("fmt"
        "encoding/json"
	    "net/http")




type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]User{

	1 : {ID : 1 , Name : "Parveen"},

}


func updateuser(w http.ResponseWriter , r *http.Request) {

	if r.Method != http.MethodPut {

		http.Error(w , "Only Put Allowed" , http.StatusMethodNotAllowed)
		return
	}


	// Read Json body

	var u User

	json.NewDecoder(r.Body).Decode(&u)


	// Check user exist 

	if _, ok := users[u.ID]; !ok {
		http.Error(w, "User not found" , http.StatusNotFound)
		return 
	}


	// update map 

	users[u.ID] = u

	fmt.Fprintln(w , "User Updated")
}




func main() {

	http.HandleFunc("/users", updateuser)

	http.ListenAndServe(":8080", nil)
}

*/