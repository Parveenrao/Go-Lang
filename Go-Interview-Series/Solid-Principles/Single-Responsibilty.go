// A struct or module should have one reason to change 

// One Responsibilty = One struct 


/* 

type OrderService struct{}

func (o OrderService) PlaceOrder(userID int, amount float64) {
	// 1. Validate order
	if amount <= 0 {
		panic("invalid amount")
	}

	// 2. Save to database
	fmt.Println("Saving order to DB")

	// 3. Process payment
	fmt.Println("Processing payment")

	// 4. Send email
	fmt.Println("Sending confirmation email")
}

*/

package main


import "fmt"


// Order validator (Validation only)

type OrderValidator struct{}

func (o OrderValidator) Validate(amount float64) error {
	if amount < 0 {
		return  fmt.Errorf("Invalid amount")
	}

	return nil
}

// Order Repository (database only)

type OrderRepository struct{}

func (r OrderRepository) Save(UserID int , amount float64) {
	fmt.Println("Order Saved In Db")
}


// Payment Service (Payment Service)

type PaymentService struct{}

func (p PaymentService) Charge(amount float64) {
	fmt.Println("Payment Processed")
}

// Email Service (Email Only)

type EmailConfirmation struct{}

func (e EmailConfirmation) SendConfirmation(UserID int) {
	fmt.Println("Confirmation Email Sent")
}

// Order Service 

type OrderService struct{

	validator  OrderValidator
	repo       OrderRepository
	payment    PaymentService
	email      EmailConfirmation
}

func (o OrderService) PlaceOrder(UserID int , amount float64) error {

	if err := o.validator.Validate(amount); err != nil {
		return  err
	}

	o.payment.Charge(amount)
	o.repo.Save(UserID, amount)
	o.email.SendConfirmation(UserID)

	return  nil
}

func main() {

	service := OrderService{

		validator: OrderValidator{},
		repo:      OrderRepository{},
		payment:   PaymentService{},
		email:     EmailConfirmation{},
	}

	service.PlaceOrder(1 , 1000)

	
}