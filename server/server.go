package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.HandleMessages)
	http.HandleFunc("/state", handlers.HandleState)
	http.ListenAndServe(":80", nil)
	log.Println("Server listening")
}
