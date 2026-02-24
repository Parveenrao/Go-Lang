// high level module should not depend on low level modules 
// Both are depend on abstraction 


package main 

import "fmt"


// Bad Example 

/*
   type MySQLDB struct{}

func (m MySQLDB) Save(data string) {
	fmt.Println("Saved to MySQL:", data)
}

type OrderService struct {
	db MySQLDB
}

func (o OrderService) CreateOrder() {
	o.db.Save("order")
}


tommorrof if you have postgress u have to change orderservice


*/


// GOOD Example 


// Abstraction 



type Database interface{
	Save (data string)
}

// MySQL Implementation 

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Save to Mysql :" , data)
}


// Postgress Implementation 


type PostgresDb struct{}

func (p PostgresDb) Save(data string) {
	fmt.Println("Saved to Postgress" , data)
}


// Order Service depend on interface not on concrete db 

type OrderService struct {

	db Database
}


func (o OrderService) CreateOrder() {

	o.db.Save("order")
}