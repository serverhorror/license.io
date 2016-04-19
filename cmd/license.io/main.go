package main

import (
	"log"
	"net/http"

	"github.com/serverhorror/license.io/api"
)

func main() {
	log.Print("Hello, license.io!")

	http.HandleFunc("/api/", api.HandleLicense)

	log.Print("all Handlers registered!")
	err := http.ListenAndServe("[::1]:8080", nil)
	if err != nil {
		log.Printf("http.ListenAndServe: %#q", err)
	}
}
