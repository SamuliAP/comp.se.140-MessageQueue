package api

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

const (
	CMD_INIT      string = "INIT"
	CMD_SHUTDOWN  string = "SHUTDOWN"
	STATE_PAUSED  string = "PAUSED"
	STATE_RUNNING string = "RUNNING"
)

func ValidState(state string) bool {
	switch state {
	case STATE_PAUSED, STATE_RUNNING:
		return true
	default:
		return false

	}
}

func GetState() (string, error) {

	return GetResponseBody(GetResponse, "http://server/state")

}

func GetCMD() (string, error) {

	return GetResponseBody(GetResponse, "http://server/cmd")

}

func PutState(state string) error {

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://server/state", bytes.NewBufferString(state))
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
