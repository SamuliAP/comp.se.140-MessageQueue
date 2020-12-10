package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMessages(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("http://httpserv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Couldn't read data source")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Couldn't read data source")
	}

	fmt.Fprint(w, string(body))
}

func main() {

	http.HandleFunc("/messages", GetMessages)
	http.ListenAndServe(":80", nil)
}
