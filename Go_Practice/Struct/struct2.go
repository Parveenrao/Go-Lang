// Nested struct 

package main

type Address struct {
	City  string
	Country string
}

type Item struct {
	Name string
	price float64
}

type Order struct {
	ID int
	Items []Item
}

type User struct {
	ID int 
	Name string 
	Address Address
	orders []Order
}