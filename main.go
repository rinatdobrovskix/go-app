package main

import (
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    
    if port == "" {
        port = "8000"
    }
    
    http.ListenAndServe(":" + port, nil)
    os.Exit(0)
}
