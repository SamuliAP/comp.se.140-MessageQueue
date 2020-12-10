package tests

import (
	"testing"
)

func TestGetMessageBodyFormat(t *testing.T) {

	body := GetMessagesBody(t)
	validateGetMessages(body, t)
}
