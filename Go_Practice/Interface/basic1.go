package main

/*
import "fmt"

// 1. Interface 

type Payment interface {
	Pay(amouunt float64) string

}

// 2. Implementations

type UPI struct{}

func (u UPI) Pay(amount float64) string {
	return  fmt.Sprintf("Pay %.2f using UPI" , amount)
}

type CARD struct{}

func (c CARD) Pay(amount float64) string {
	return  fmt.Sprintf("Pay %.2f using CARD" , amount)
}

// 3. Use Interface 

func ProcessPayment(p Payment , amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {

	ProcessPayment(UPI{}, 500)
	ProcessPayment(CARD{}, 1200)

}

*/