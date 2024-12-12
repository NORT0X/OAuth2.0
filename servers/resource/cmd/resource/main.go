package main

import (
	"log"
	"net/http"
	"resource/internal/routes"
)

func main() {

	r := routes.NewRouter()

    log.Println("Starting resource server on: 8082")

    if err := http.ListenAndServe(":8082", r); err != nil {
        log.Fatal(err)
    }

}
