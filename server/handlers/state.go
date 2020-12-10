package handlers

import (
	"../../api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var state = api.STATE_RUNNING

func HandleState(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		GetState(w, req)
	case http.MethodPut:
		SetState(w, req)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid HTTP method")
	}
}

func GetState(w http.ResponseWriter, req *http.Request) {
	log.Println("200 GET", req.URL)
	fmt.Fprint(w, state)
}

func SetState(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stringBody := string(reqBody)

	if stringBody == api.CMD_INIT {
		// empty datastore
		os.Create("/app/data/data.txt")
		fmt.Fprint(w, stringBody)
		state = api.STATE_RUNNING
		return
	}

	if api.ValidState(stringBody) {
		log.Println("200 PUT", req.URL, stringBody)
		fmt.Fprint(w, stringBody)
		state = stringBody
		return
	}

	log.Println("400 PUT", req.URL, stringBody)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Invalid request body")
}
