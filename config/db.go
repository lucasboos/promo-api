package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sqlx.DB
	once       sync.Once
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbInstance, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	dbInstance.SetMaxOpenConns(25)
	dbInstance.SetMaxIdleConns(25)
	dbInstance.SetConnMaxLifetime(5 * 60)

	if err := dbInstance.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Printf("Connecting to database at %s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
}

func GetDB() *sqlx.DB {
	once.Do(func() {
		ConnectDB()
	})
	return dbInstance
}

func CloseDB() {
	if dbInstance != nil {
		if err := dbInstance.Close(); err != nil {
			log.Printf("Error closing the database: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
