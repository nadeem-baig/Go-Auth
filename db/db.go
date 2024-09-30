package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nadeem-baig/go-auth/utils/logger"
)


func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
func ConnectDB() (*sql.DB ,error){
        // Load environment variables from .env file
    // Retrieve environment variables
    // Create the connection string with allowNativePasswords=true
    // Log the connection attempt
    // Open the connection to MySQL
    // Ping to verify connection
    
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
        return nil, err
	}

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "Go_Auth")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", user, password, host, port, dbname)

	log.Printf("Attempting to connect to DB at %s:%s", host, port)

	db, err := sql.Open("mysql", connectionString)
	logger.Fatal("Failed to open database connection", err)
    if err!=nil {
        return nil,err
    }
	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to ping database", err)
        return nil, err
	}

	log.Println("Successfully connected to the database")
	return db,nil
}
