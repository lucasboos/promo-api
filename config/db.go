package config

import (
    "log"
    "os"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error to load .env: %v", err)
    }

    connStr := "host=" + os.Getenv("DB_HOST") + 
        " port=" + os.Getenv("DB_PORT") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " sslmode=disable"

    DB, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        log.Fatalf("Error to connect database: %v", err)
    }
    log.Println("Database connected.")
}
