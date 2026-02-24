// Software enitites should be open for extension  but close for modification

// You shoudl be able to add new features
// Without modifying existing working code

/*
   => Bad Design

   type NotificationService struct{}

func (n NotificationService) Send(method string, message string) {

	if method == "email" {
		fmt.Println("Sending Email:", message)
	} else if method == "sms" {
		fmt.Println("Sending SMS:", message)
	}
}


*/

package main

import (
	"fmt"

)

type Notifier interface {
	Send(message string)
}

// Email Notifier Implementation 

type EmailNotifier struct{}

func(e EmailNotifier) Send(message string) {
	fmt.Println("Sending Email " , message)
}

// SMS Implementation 

type SMSNotifier struct{}

func(s SMSNotifier) Send(message string) {
	fmt.Println("Sending Email" , message)
}


// Service depend on NOtifier 

type NotificationService struct {
	notifier Notifier
}

func (n NotificationService) Notify(message string) {
	n.notifier.Send(message)
}


func main() {

	email := EmailNotifier{}

	service := NotificationService{
		notifier: email,
	}

	service.Notify("Welcome User")
}
