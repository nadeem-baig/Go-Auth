package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/joho/godotenv"
	"github.com/nadeem-baig/go-auth/utils/logger"
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Retrieve environment variables
    host := getEnv("DB_HOST", "localhost")
    port := getEnv("DB_PORT", "3306")
    user := getEnv("DB_USER", "root")
    password := getEnv("DB_PASSWORD", "")
    dbname := getEnv("DB_NAME", "Go_Auth")

    // Create the connection string with allowNativePasswords=true
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", user, password, host, port, dbname)

    // Log the connection attempt
    log.Printf("Attempting to connect to DB at %s:%s", host, port)

    // Open the connection to MySQL
    db, err := sql.Open("mysql", connectionString)
    logger.Fatal("Failed to open database connection", err)

    // Ping to verify connection
    err = db.Ping()
    logger.Fatal("Failed to ping database", err)

    log.Println("Successfully connected to the database")
}
