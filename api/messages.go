package api

import (
	"testing"
)

func GetMessagesBody() (string, error) {

	return GetResponseBody(GetResponse, "http://server/messages")
}

func GetMessagesBodyHandleError(t *testing.T) string {
	body, err := GetMessagesBody()
	if err != nil {
		t.Error(err.Error())
	}
	return body
}
