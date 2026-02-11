// Balance error --> function  withdraw amount --> if amount > balance custom error --> catch using error.As

package main

/*
import (
	"errors"
	"fmt"
)

type BalanceError struct {
	Balance int
}

// Implement error interface
func (e BalanceError) Error() string {
	return fmt.Sprintf("insufficient balance: %d", e.Balance)
}

var balance = 100

func withdraw(amount int) error {

	if amount > balance {
		return BalanceError{Balance: balance}
	}

	balance -= amount
	return nil
}

func main() {

	err := withdraw(150)

	if err != nil {

		var balErr BalanceError

		if errors.As(err, &balErr) {
			fmt.Println("Custom Error Caught!")
			fmt.Println("Available balance:", balErr.Balance)
			return
		}

		fmt.Println(err)
		return
	}

	fmt.Println("Withdrawal successful")
	fmt.Println("Remaining balance:", balance)
}
*/