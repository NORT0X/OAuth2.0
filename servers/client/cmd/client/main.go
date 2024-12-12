package main

import (
	"client/internal/routes"
	"log"
	"net/http"
)

func main() {
    r := routes.NewRouter()
    log.Println("Starting server on: 8081")
    if err := http.ListenAndServe(":8081", r); err != nil {
        log.Fatal(err)
    }
}
