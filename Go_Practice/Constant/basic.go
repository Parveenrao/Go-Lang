/* Const  => A constant in go language is a value that cannot be changed after declare 
          => Exist at compile time 
		  => has no memorya address
*/

package main

import "fmt"

func main(){

// simple declaration 
const pi = 3.14

// type const

const a int = 10

//a = 20 error

// Untype const

const b = 40

// Mulitple constant 


const ( 
	Appname = "Go"
	version  = 1.0
)

const (
	Admin = iota
	Member
	Guest
)


fmt.Println(Appname , version)
fmt.Println(Admin , Member , Guest)
}