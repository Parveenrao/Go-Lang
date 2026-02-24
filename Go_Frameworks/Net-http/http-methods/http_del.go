/* 
   Http - Delete Method 

   1. Get id from url
   2. Convert id into int
   3. check exist 
   4. delete (map , id) 
*/

package main

import ("fmt"
       "strconv"
	   "net/http")

type User struct {
	ID   int
	Name string
}

var users = map[int]User{
	1: {ID: 1, Name: "Parveen"},
	2: {ID: 2, Name: "Rahul"},
}


func deleteuser(w http.ResponseWriter , r * http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w , "Only Delete Allowed" , http.StatusMethodNotAllowed)
		return
	}

	//Read id from query

	idstr := r.URL.Query().Get("id")

	if idstr == "" {
		http.Error(w , "id Required" , http.StatusBadRequest)
		return 
	}

	// Convert id into int 

	id, err := strconv.Atoi(idstr)

	if err != nil {
		http.Error(w , "Invalid id" , http.StatusBadRequest)
		return 
	}


		// 4. Check user exists
	if _, ok := users[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// 5. Delete from map
	delete(users, id)

	fmt.Fprintln(w, "User deleted")
}

func main() {

	http.HandleFunc("/users", deleteuser)

	http.ListenAndServe(":8080", nil)
}


