// Full example of pointers 

package main 

import "fmt"

type User struct {
	Name string
}


func update (a *int ){
	*a = 100
}

func updateName (u *User){
	u.Name = "Hello"
}
func main() {


	// Basic Pointer
	x := 10
	p := &x

	fmt.Print("Value " , x)
	fmt.Println("Adress " , p)
	fmt.Println("Deference " , *p )

	// Modify via Pointer 

	*p = 100

	fmt.Println("After modification " , x)

	// function pointer

	update(&x)
	fmt.Println("After function modification " , x)

	// struct pointer

	user := User{"Parveen"}
	updateName(&user)
	fmt.Println(user.Name)

}