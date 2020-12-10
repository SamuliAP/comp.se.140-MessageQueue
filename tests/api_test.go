package tests

import (
	"testing"
)

func TestInitialGetMessageBodyFormat(t *testing.T) {

	body := GetMessagesBody(t)
	validateGetMessages(body, t)
}
