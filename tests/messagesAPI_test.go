package tests

import (
	"../api"
	"testing"
	"time"
)

func TestInitialGetMessageBodyFormat(t *testing.T) {

	time.Sleep(10 * time.Second)
	body := api.GetMessagesBodyHandleError(t)
	validateGetMessages(body, t)
}
