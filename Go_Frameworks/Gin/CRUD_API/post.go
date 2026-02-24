package main 

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)


type User struct {

	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required"`
}


func main() {

	// Load env 

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading env")
	}

	// Build connection string from .env
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)


	db ,err := sql.Open("postgres" , connStr) 

	if err != nil {

		log.Fatal(err)
	}

	defer db.Close()

		if err := db.Ping(); err != nil {
		log.Fatal("DB not connected:", err)
	}

	log.Println("Postgres connected ✅")



	 r := gin.Default()


	 // Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})


	// post users

	// POST /users  -> insert into PostgreSQL
	r.POST("/users", func(c *gin.Context) {

		var user User

		// Bind + validate JSON
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	query := `
			INSERT INTO users (name, email, age)
			VALUES ($1, $2, $3)
			RETURNING id
		`

		err := db.QueryRow(query, user.Name, user.Email, user.Age).Scan(&user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user created",
			"data":    user,
		})
	})

	r.Run(":8080")
}
