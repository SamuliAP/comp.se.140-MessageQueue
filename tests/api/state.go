package api

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

const (
	INIT     string = "INIT"
	PAUSED   string = "PAUSED"
	RUNNING  string = "RUNNNING"
	SHUTDOWN string = "SHUTDOWN"
)

func GetState() (string, error) {

	return GetResponseBody(GetResponse, "http://api/state")

}

func PutState(state string) error {

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://api/state", bytes.NewBufferString(state))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("API responded with status: " + res.Status)
	}

	return nil
}

func PutStateHandleError(t *testing.T, state string) {
	err := PutState(state)
	if err != nil {
		t.Error(err.Error())
	}
}

func GetStateHandleError(t *testing.T) string {
	body, err := GetState()
	if err != nil {
		t.Error(err.Error())
	}
	return body
}
