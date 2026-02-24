// Banking management System 

package main

/* 

import "fmt"

// Deposite money 

func Deposite(balance float64 , amount float64) float64  {
	return  amount + balance
	
}

// Withdraw money 

func Withdraw(balance float64 , amount float64) float64  {

	if amount > balance{
		fmt.Println("Insufficient fund")
	}

	return balance - amount
	
}

// Show balance 

func  showbalance(balance float64) float64  {
	fmt.Println("Current balance" , balance)
	
}


func main(){
	var name string
	var balance float64
	var amount float64
	var choice float64


	fmt.Println("Enter your name :")
	fmt.Scanln(&name)

	fmt.Println("Enter staarting balance:")
	fmt.Scanln(&balance)

	for {

		fmt.Println("n---- Banking Menu----:")
		fmt.Println("1. Deposite:")
		fmt.Println("2. Withdraw:")
		fmt.Println("3. Check Balance:")
		fmt.Println("4. Exit:")


		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)


		switch choice {
			
		case 1:
			fmt.Print("Enter Deposite amount:")
			fmt.Scanln(&amount)
			balance = deposit(balance, amount)
            showBalance(balance)

        case 2:
            fmt.Print("Enter withdraw amount: ")
            fmt.Scanln(&amount)
            balance = withdraw(balance, amount)
            showBalance(balance)

        case 3:
            showBalance(balance)

        case 4:
            fmt.Println("Thank you,", name)
            return

        default:
            fmt.Println("Invalid option")
        }
    }
}
			
			
				/* */

				  
