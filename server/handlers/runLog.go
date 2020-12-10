package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleRunLog(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid HTTP method")
		return
	}

	log.Println("200 GET", req.URL)

	content, err := ioutil.ReadFile("/app/data/run-log.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Couldn't read data source")
		log.Println("Couldn't read data source")
	}

	fmt.Fprint(w, string(content))
}
