package tests

import (
	"./api"
	"testing"
)

func TestInitialGetMessageBodyFormat(t *testing.T) {

	body := api.GetMessagesBodyHandleError(t)
	validateGetMessages(body, t)
}
