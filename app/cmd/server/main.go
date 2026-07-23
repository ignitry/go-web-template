package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, routes()))
}
