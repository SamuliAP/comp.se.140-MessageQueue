package main

import (
	"../api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FileContents(w http.ResponseWriter, req *http.Request) {

	content, err := ioutil.ReadFile("/app/data/data.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Couldn't read data source")
		log.Println("Couldn't read data source")
	}

	if string(content) == api.CMD_SHUTDOWN {
		os.Exit(0)
	}
	log.Println("Responded with status 200")
	fmt.Fprint(w, string(content))
}

func main() {
	http.HandleFunc("/", FileContents)
	http.ListenAndServe(":80", nil)
	log.Println("Listening for http...")
}
