package tests

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestHTTPConnection(t *testing.T) {

	resp, err := GetHTTPSERVResponse(t)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Error("HTTPSERV did not respond with HTTP status 200")
	}
}

func TestHTTPSERVBody(t *testing.T) {
	resp, err := GetHTTPSERVResponse(t)
	if err != nil {
		t.Error(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("could not read response body")
	}

	validateBody(string(body), t)
}

// does allow whitespace in end of rows
func validateBody(body string, t *testing.T) {
	rows := strings.Split(body, "\n")
	for _, row := range rows {
		parts := strings.Split(row, " ")
		for i, part := range parts {
			if part == "" {
				continue
			}

			switch i {
			case 0:
				ValidateTimestamp(part, t)
			case 1:
				ValidateTopic(part, t)
			case 2:
				ValidateTopicName(part, t)
			default:
				ValidateIsWhitespace(part, t)
			}
		}
	}
}
func ValidateTimestamp(s string, t *testing.T) {
	_, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t.Error("Invalid date format in response: ", s)
	}
}

func ValidateTopic(s string, t *testing.T) {
	if s != "Topic" {
		t.Error("Invalid Topic format in response: ", s)
	}
}

func ValidateTopicName(s string, t *testing.T) {
	if s != "my.i" && s != "my.o" {
		t.Error("Invalid Topic name in response: ", s)
	}
}

func ValidateIsWhitespace(s string, t *testing.T) {
	_, err := regexp.Match(`\d`, []byte(s))
	if err != nil {
		t.Error("Invalid row format, extra ouput found in the end of row")
	}
}

func GetHTTPSERVResponse(t *testing.T) (*http.Response, error) {

	t.Log(time.Now().Format(time.RFC3339Nano), " attempting to connect to HTTPSERV")
	for i := 0; i < 30; i++ {

		t.Log(time.Now().Format(time.RFC3339Nano), " connecting...")

		resp, err := http.Get("http://localhost:8081")
		if err == nil {
			t.Log(time.Now().Format(time.RFC3339Nano), " connected!")
			return resp, nil
		}

		time.Sleep(5 * time.Second)
	}

	return nil, errors.New("could not connect to HTTPSERV via HTTP")
}
