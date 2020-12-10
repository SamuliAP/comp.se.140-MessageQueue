package tests

import (
	"./api"
	"strings"
	"testing"
	"time"
)

func TestInitState(t *testing.T) {

	// put state in running in case the queue had been paused empty
	api.PutStateHandleError(t, api.RUNNING)
	time.Sleep(10 * time.Second)

	body := api.GetMessagesBodyHandleError(t)
	rows := strings.Split(body, "\n")
	if len(rows) == 0 {
		t.Error("Messages empty after 10 seconds of RUNNING")
	}

	// pause to change state from running, init again, messages should now be empty and state should be running
	api.PutStateHandleError(t, api.PAUSED)
	api.PutStateHandleError(t, api.INIT)
	body = api.GetMessagesBodyHandleError(t)
	rows = strings.Split(body, "\n")
	if len(rows) != 0 {
		t.Error("Messages not empty after INIT")
	}

	// state should be RUNNING
	state := api.GetStateHandleError(t)
	if state != api.RUNNING {
		t.Error("State not RUNNING after INIT")
	}

	// make sure messages are populated after init
	time.Sleep(10 * time.Second)
	body = api.GetMessagesBodyHandleError(t)
	rows = strings.Split(body, "\n")
	if len(rows) == 0 {
		t.Error("Messages empty after 10 seconds of INIT")
	}
}

func TestPauseState(t *testing.T) {

	// first run init to make sure state is not paused, then pause
	api.PutStateHandleError(t, api.RUNNING)
	api.PutStateHandleError(t, api.PAUSED)
	state := api.GetStateHandleError(t)
	if state != api.PAUSED {
		t.Error("State not PAUSED")
	}

	// we'll get all current messages and wait 10 seconds
	// to verify ORIG hasn't pushed any more messages to the queue
	originalBody := api.GetMessagesBodyHandleError(t)
	validateGetMessages(originalBody, t)

	time.Sleep(10 * time.Second)

	newBody := api.GetMessagesBodyHandleError(t)
	validateGetMessages(newBody, t)

	if originalBody != newBody {
		t.Error("PUT PAUSE did not stop message generation")
	}

	// reset state
	api.PutStateHandleError(t, api.INIT)
}

func TestRunningState(t *testing.T) {

	api.PutStateHandleError(t, api.PAUSED)
	api.PutStateHandleError(t, api.RUNNING)
	state := api.GetStateHandleError(t)
	if state != api.RUNNING {
		t.Error("State not RUNNING")
	}

	// we'll get all current messages and wait 10 seconds
	// to verify ORIG has pushed more messages to the queue, and that the original messages
	// are still at the top
	originalBody := api.GetMessagesBodyHandleError(t)
	validateGetMessages(originalBody, t)

	time.Sleep(10 * time.Second)

	newBody := api.GetMessagesBodyHandleError(t)
	validateGetMessages(newBody, t)

	if len(originalBody) <= len(newBody) {
		t.Error("PUT RUNNING did not start message generation")
	}

	if !strings.HasPrefix(newBody, originalBody) {
		t.Error("PUT RUNNING modified the original messages")
	}

	// reset state
	api.PutStateHandleError(t, api.INIT)
}

func TestShutdownState(t *testing.T) {

	api.PutStateHandleError(t, api.SHUTDOWN)
	_, err := api.GetResponse("http://api/messages")
	if err == nil {
		t.Error("Services still running after shutdown")
	}
}
