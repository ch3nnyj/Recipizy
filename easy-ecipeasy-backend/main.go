package main

import (
	"github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "context"
    "log"
)

var db *pgx.Conn

func main() {
    var err error
	db, err = pgx.Connect(context.Background(), "postgres://chenny:Jj8458535220!@localhost:5432/easyrecipeasydb")
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
    r.Run(":8080")
}