package tests

import (
	"testing"
)

func TestHTTPConnection(t *testing.T) {

	resp, err := GetMessage(t)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Error("API did not respond with HTTP status 200")
	}
}
