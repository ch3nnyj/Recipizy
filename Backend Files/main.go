package main

import (
	"github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "github.com/joho/godotenv"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "github.com/dgrijalva/jwt-go"

    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "easy-recipeasy-backend/server"
    "easy-recipeasy-backend/auth"
)

var db *pgx.Conn

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    dbURL := os.Getenv("DATABASE_URL")

	db, err = pgx.Connect(context.Background(), dbURL)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer db.Close(context.Background())
    r := gin.Default()

    // Example Route
    r.GET("/api/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "OK"})
    })

	// Route to check database connection
	r.GET("/api/db-check", func(c *gin.Context) {
		var result string
		err := db.QueryRow(context.Background(), "SELECT 'DB Connection Successful'").Scan(&result)
		if err != nil {
			c.JSON(500, gin.H{"error": "Database connection failed"})
			return
		}
		c.JSON(200, gin.H{"message": result})
	})

    // Start server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v\n", err)
    }
}
