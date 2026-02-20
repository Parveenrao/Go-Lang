// In Liskov - Subsitution Principles

// If your function accept an  interface then any struct that implement interface should work correctly there , Not panic, NO Crash

package main

import "fmt"

/*
import "fmt"

type Payment interface {
	Pay(amount float64)
}


type CardPayment struct {}

func (c CardPayment) Pay(amount float64) {
	fmt.Println("Paid" , amount)
}


type FreePayment struct{}

func (f FreePayment) Pay(amount float64) {

	panic("I don't support freepayment")
}

*/

// Good Design

type Payment interface{
	Pay(amount float64) error
}


type Card struct{}

func (c Card) Pay(amount float64) error {

	fmt.Println("Paid :" , amount)

	return  nil
}


type Free struct{}

func (f Free) Pay(amount float64) error {
	if amount != 0 {
		return fmt.Errorf("Free support only zero")
	}

	fmt.Println("Free Order")

	return  nil
}

func Process(p Payment) {
	if err := p.Pay(100); err != nil {
		fmt.Println(err)
	}
}
