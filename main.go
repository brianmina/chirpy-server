package main

import (
    "log"
    "net/http"
)

func main() {
    // Create a new mux (router)
    mux := http.NewServeMux()
    

    // Add a file server handler for the root path
    fileServer := http.FileServer(http.Dir("."))
    mux.Handle("/", fileServer)

    // Start the server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
