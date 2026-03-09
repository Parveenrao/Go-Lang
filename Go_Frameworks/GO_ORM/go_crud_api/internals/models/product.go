package models

import (
        "gorm.io/gorm"
	     "time")



// Product represt the product table in the database 

type Product struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null"        json:"name"`
	Description string         `gorm:"type:text"                          json:"description"`
	Price       float64        `gorm:"type:decimal(10,2);not null"        json:"price"`
	Stock       int            `gorm:"default:0"                          json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // soft delete
}


// CreateProductInput is used for binding POST request body.
type CreateProductInput struct {
	Name        string  `json:"name"        binding:"required,min=2,max=255"`
	Description string  `json:"description"`
	Price       float64 `json:"price"       binding:"required,gt=0"`
	Stock       int     `json:"stock"       binding:"min=0"`
}


// UpdateProductInput is used for binding PUT/PATCH request body.
type UpdateProductInput struct {
	Name        *string  `json:"name"        binding:"omitempty,min=2,max=255"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"       binding:"omitempty,gt=0"`
	Stock       *int     `json:"stock"       binding:"omitempty,min=0"`

}