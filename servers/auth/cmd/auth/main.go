package main

import (
	"log"
	"net/http"

	"auth/internal/routes"
)

func main() {
	r := routes.NewRouter()
	log.Println("Starting server on: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
