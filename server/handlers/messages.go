package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleMessages(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid HTTP method")
		return
	}

	log.Println("200 GET", req.URL)

	resp, err := http.Get("http://httpserv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		fmt.Fprint(w, "Couldn't read data source")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		fmt.Fprint(w, "Couldn't read data source")
		return
	}

	fmt.Fprint(w, string(body))
}
