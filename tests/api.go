package tests

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

func GetMessage(t *testing.T) (*http.Response, error) {

	t.Log(time.Now().Format(time.RFC3339Nano), " attempting to connect to API")
	for i := 0; i < 30; i++ {

		t.Log(time.Now().Format(time.RFC3339Nano), " connecting...")

		resp, err := http.Get("http://api")
		if err == nil {
			t.Log(time.Now().Format(time.RFC3339Nano), " connected!")
			return resp, nil
		}

		time.Sleep(5 * time.Second)
	}

	return nil, errors.New("could not connect to API via HTTP")
}
