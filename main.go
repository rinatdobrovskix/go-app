package main

import (
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    
    http.ListenAndServe(":" + port, nil)
    os.Exit(0)
}
