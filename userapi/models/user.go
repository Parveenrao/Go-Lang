package models

import "gorm.io/gorm"

// User is the GORM model — maps directly to the "users" table in Postgres.
//
// gorm.Model gives us for free:
//   - ID          uint           (primary key, auto increment)
//   - CreatedAt   time.Time      (set automatically on Create)
//   - UpdatedAt   time.Time      (set automatically on Update)
//   - DeletedAt   gorm.DeletedAt (soft delete — sets timestamp instead of removing row)
type User struct {
	gorm.Model

	Name     string `gorm:"not null"                json:"name"`
	Email    string `gorm:"uniqueIndex;not null"    json:"email"`
	Age      int    `gorm:"default:0"               json:"age"`
	Role     string `gorm:"default:'user';size:50"  json:"role"`
}