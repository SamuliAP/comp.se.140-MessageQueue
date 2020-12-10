package handlers

import (
	"../../api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

	runLog := GetRunLogFile(w)
	defer runLog.Close()

	if stringBody == api.CMD_INIT {
		WriteRunLog(runLog, api.CMD_INIT)
		state = api.CMD_INIT
		// empty datastore
		os.Create("/app/data/data.txt")
		fmt.Fprint(w, stringBody)
		WriteRunLog(runLog, api.STATE_RUNNING)
		state = api.STATE_RUNNING
		return
	}

	if stringBody == api.CMD_SHUTDOWN {

		WriteRunLog(runLog, api.CMD_SHUTDOWN)
		state = api.CMD_SHUTDOWN

		// wait for httpserv to exit, this is the end of the queue so
		// we can be sure the queue has been shut down when this server shuts down
		for true {
			_, err := http.Get("http://httpserv")
			if err != nil {
				os.Exit(0)
			}
			time.Sleep(5 * time.Second)
		}
	}

	if api.ValidState(stringBody) {
		log.Println("200 PUT", req.URL, stringBody)
		fmt.Fprint(w, stringBody)
		WriteRunLog(runLog, stringBody)
		state = stringBody
		return
	}

	log.Println("400 PUT", req.URL, stringBody)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Invalid request body")
}

func WriteRunLog(runLog *os.File, newState string) {
	if state == newState {
		return
	}
	_, err := runLog.WriteString(time.Now().Format(time.RFC3339Nano) + ": " + newState + "\n")
	if err != nil {
		log.Println("Couldn't write runlog")
	}
}

func GetRunLogFile(w http.ResponseWriter) *os.File {
	runLog, err := os.OpenFile("/app/data/run-log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Couldn't open data store")
	}
	return runLog
}
