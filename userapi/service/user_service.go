package service

import (
	"errors"
	"unicode"
	"userapi/models"
	"userapi/repository"

	"github.com/parveenrao/userapi/models"
	repository "github.com/parveenrao/userapi/repo"
)

// -------------------------------------------------------
// SERVICE LAYER
//
// The service holds BUSINESS LOGIC — rules about what's allowed.
// It knows nothing about HTTP (no gin here) or SQL (no gorm here).
//
// Examples of business logic:
//   - "You can't register with an email that's already taken"
//   - "Only admins can change a user's role"
//   - "Age must be between 0 and 120"
//
// This separation means you could swap Gin for gRPC tomorrow
// and keep all your business rules untouched.
// -------------------------------------------------------

// Sentinel errors — named errors your handlers can check against.
// Much cleaner than comparing error message strings.
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyTaken = errors.New("email already taken")
	ErrInvalidAge        = errors.New("age must be between 0 and 120")
)



type UserService interface {
	CreateUser(name , email string , age int) (*models.User , error)
	GetAllUsers() ([]models.User , error)
	GetByUserId(id uint) (*models.User , error)
	UpdateUser(id uint , name , email string , age int) (*models.User , error)
	Delete(id uint) error 
}


// concrete implementation 

type userService struct {
	repo repository.UserRepository 

}

// NewUserService is the constructor — accepts a repository, returns a service.

func NewUserService(repo repository.UseRepository) UserService {
	return  &userService{repo : repo}
}


// CreateUser validates input, checks for duplicate email, then creates the user.

func(s *userService) CreateUser(name ,email string , age int) (*models.User , error) {

	// validate age 

	if age  < 0 || age > 120 {
		return  nil , ErrInvalidAge
	}

	// email must not be taken already 

	existing , err := s.repo.FindByEmail(email)

	if err != nil {
		return  nill , err
	}

	if existing != nil {
		 return  nil , ErrEmailAlreadyTaken
	}


		user := &models.User{
		Name:  name,
		Email: email,
		Age:   age,
		Role:  "user", // default role — only admins can change this
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser applies changes to an existing user.
func (s *userService) UpdateUser(id uint, name, email string, age int) (*models.User, error) {
	// Business rule: validate age
	if age < 0 || age > 120 {
		return nil, ErrInvalidAge
	}

	// Make sure user exists
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// Business rule: if email is changing, make sure new email isn't taken
	if email != user.Email {
		existing, err := s.repo.FindByEmail(email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, ErrEmailAlreadyTaken
		}
	}

	// Apply updates
	user.Name = name
	user.Email = email
	user.Age = age

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}


// DeleteUser soft-deletes a user by ID.
func (s *userService) DeleteUser(id uint) error {
	// Make sure user exists before deleting
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return s.repo.Delete(id)

}