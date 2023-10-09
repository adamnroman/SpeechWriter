package main

import (
	"log"
	"net/http"

	"github.com/adamnroman/SpeechWriter/pkg/handler"
)

func main() {

	http.HandleFunc("/ping", handler.Pong)
	log.Println("Listing for requests at http://localhost:8000/ping")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
