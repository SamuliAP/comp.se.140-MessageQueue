package main

import (
	"./handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	os.Create("/app/data/run-log.txt")
	http.HandleFunc("/messages", handlers.HandleMessages)
	http.HandleFunc("/state", handlers.HandleState)
	http.HandleFunc("/run-log", handlers.HandleRunLog)
	http.ListenAndServe(":80", nil)
	log.Println("Server listening")
}
