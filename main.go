// main.go
package main

import (
    "API/config"
    "API/router"
    "log"
    "net/http"
)

func main() {
    // Initialize database connection
    err := config.InitDB()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer config.DB.Close()

    // Create router
    r := router.NewRouter()
    
    // Start server
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
