// Client should not forced to depend on methods they dont use 

// Dont create big interface , split them into small focused interfaces

/*
   => Imagine we design a example 
       
      type PaymentProcessor interface{
	       Pay(amount float64) error
		   Refund(amount) error 
		   GenerateInvoive() string 
		   }
        
		Now imagine card payment support refund payment 

		But UPI payment doesn't support refund payment 
		--> so UPI is forced to implement Refund and GenerateInvoice()
		     even if it doesnt need 

        This voilates the ISP			 

*/


// Good Design

// Small Interfaces

package main

import "fmt"

type Payer interface{
	Pay(amount float64) error
}


type Refunder interface{
	Refund(amount float64) error 
}


type Invoicer interface {
	GenerateInvoice() error
}


// Card implement all 

type CardPayment struct{}

func (c CardPayment) Pay(amount float64) error {

	fmt.Println("Paid By Card :" , amount)

	return nil
}

func (c CardPayment) Refund(amount float64) error {

	fmt.Println("Refunded :" , amount)

	return  nil
}

func ( c CardPayment) GenerateInvoice() string {

	return  "Card Invoice"
}



// UPI Struct implment only which is needed 


type UPIPayment struct{}

func (u UPIPayment) Pay(amount float64) error {

	fmt.Println("Paid using UPI" , amount)

	return nil
}