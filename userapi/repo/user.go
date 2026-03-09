package repository

import (
	"errors"
	"userapi/models"

	"github.com/parveenrao/userapi/models"
	"gorm.io/gorm"
)

// -------------------------------------------------------
// REPOSITORY PATTERN
//
// The repository is the ONLY layer that talks to the database.
// Service layer calls repository methods — it never touches db directly.
//
// Benefits:
//   - Swap Postgres for MySQL? Change only this file.
//   - Want to test without a real DB? Mock this interface.
//   - All DB logic is in one place — easy to find and change.
// -------------------------------------------------------

// UserRepository defines what operations are possible on users.
// It's an interface — the service layer depends on this contract,
// not on the concrete implementation below.


type UserRepository interface {

	Create(user *models.User) error
	FindAll([]models.User , error)
	FindByID(id uint) (*models.User , error)
	FindByEmail(email string) (*models.User , error)
	Update(user *models.User) error
	Delete(id uint) error
}


// userRepository is the concrete struct that implements UserRepository.
// It holds a *gorm.DB — passed in via NewUserRepository (not a global!).
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is the constructor.
// main.go calls this, passing the *gorm.DB it already created.
// Returns the interface, not the concrete type — callers only see the contract.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


// Create inserts a new user row into the database.
// GORM fills in user.ID, user.CreatedAt, user.UpdatedAt after this call.

func (r *userRepository) Create(user *models.User) error {

	result := r.db.Create(user)

	return  result.Error

}

// FindALll  

func (r * userRepository) FindAll([]models.User , error ) {
	var users models.User

	result := r.db.Find(&users)

	return result , result.Error
}


// FIndby iD 

func(r *userRepository) FindBYID(ID uint )(*models.User , error) {

	var user models.User 

	result := r.db.First(&user , ID)

	if errors.Is(result.Error  , gorm.ErrRecordNotFound) {
		return nil ,  nil
	}

	return result , result.Error
}

// Find by email 

func (r *userRepository) FindByEmail(email string ) (*models.User , error) {

	var user models.User

	result := r.db.where("email = ?" , email).First(&user)

	if errors.Is(result.Error , gorm.ErrRecordNotFound) {

		return  nil , nil 
		
	}

	return  result , result.Eror
}


// Update saves changes to an existing user.
// We use Updates with a map so zero-value fields (like Age=0) are also saved.
// If we used Updates(user struct), GORM would skip zero values — a common bug!
func (r *userRepository) Update(user *models.User) error {
	result := r.db.Model(user).Updates(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"age":   user.Age,
		"role":  user.Role,
	})
	return result.Error
}


// Delete performs a SOFT DELETE — sets deleted_at timestamp.
// The row stays in Postgres but GORM hides it from all future queries.
// Use db.Unscoped().Delete() if you ever need a hard (permanent) delete.
func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
