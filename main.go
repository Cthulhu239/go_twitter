package main

import (
    "github.com/Cthulhu239/go_twitter/db"     
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    db.Connect()
    defer db.DB.Close()

    r := gin.Default()

    log.Println("server running on :8080")
    r.Run(":8080")
}