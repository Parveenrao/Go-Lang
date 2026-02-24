package main 

import "fmt"

type Users struct {
	ID int
	Name string
	Age int 
}


func (u *Users) Birthday() {
	u.Age++
}

func  main()  {

	// create struct

	user := Users{
		ID: 1 ,
		Name : "Parveen",
		Age : 22,
	}
	
	fmt.Println("Before :" ,user)

	// call method 

	user.Birthday()

	fmt.Println("After :" , user)

	// pointer to struct

	p := &user

	p.Name = "Developer"

	fmt.Println("Updated" , user)
}