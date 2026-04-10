package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
    godotenv.Load()

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )
    
    database, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("failed to open db:", err)
    }

    if err = database.Ping(); err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    database.SetMaxOpenConns(25)
    database.SetMaxIdleConns(10)

    DB = database
	
    log.Println("connected to mysql")
	
}