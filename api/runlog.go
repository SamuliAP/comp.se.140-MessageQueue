package api

import "testing"

func GetRunlog() (string, error) {

	return GetResponseBody(GetResponse, "http://server/run-log")
}

func GetRunLogHandleError(t *testing.T) string {
	log, err := GetRunlog()
	if err != nil {
		t.Error(err.Error())
	}
	return log
}
