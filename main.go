package main

import (
    "net/http"
    "os"
)

func main() {
    http.ListenAndServe(":8080", nil)
    os.Exit(0)
}
