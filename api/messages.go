package api

import (
	"testing"
)

func GetMessages() (string, error) {

	return GetResponseBody(GetResponse, "http://server/messages")
}

func GetMessagesBodyHandleError(t *testing.T) string {
	body, err := GetMessages()
	if err != nil {
		t.Error(err.Error())
	}
	return body
}
