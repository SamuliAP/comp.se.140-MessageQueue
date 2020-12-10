package tests

import (
	"../api"
	"strings"
	"testing"
	"time"
)

func TestInitState(t *testing.T) {

	// put state in running in case the queue had been paused empty
	api.PutStateHandleError(t, api.STATE_RUNNING)
	time.Sleep(10 * time.Second)

	body := api.GetMessagesBodyHandleError(t)
	rows := strings.Split(body, "\n")
	if len(rows) == 0 {
		t.Error("Messages empty after 10 seconds of RUNNING")
	}

	// pause to change state from running, init again, messages should now be empty and state should be running
	api.PutStateHandleError(t, api.STATE_PAUSED)
	api.PutStateHandleError(t, api.CMD_INIT)
	body = api.GetMessagesBodyHandleError(t)
	rows = strings.Split(body, "\n")

	// as this is timing based integration testing, leniency for 1 msg
	if len(rows) > 1 {
		t.Error("Messages not emptied after CMD_INIT")
	}

	// state should be STATE_RUNNING
	state := api.GetStateHandleError(t)
	if state != api.STATE_RUNNING {
		t.Error("State not RUNNING after CMD_INIT")
	}
}

func TestPauseState(t *testing.T) {

	// first run init to make sure state is not paused, then pause
	api.PutStateHandleError(t, api.STATE_RUNNING)
	api.PutStateHandleError(t, api.STATE_PAUSED)
	state := api.GetStateHandleError(t)
	if state != api.STATE_PAUSED {
		t.Error("State not PAUSED")
	}

	// make sure all the messages have propagated through the queue
	time.Sleep(3 * time.Second)

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
	api.PutStateHandleError(t, api.CMD_INIT)
}

func TestRunningState(t *testing.T) {

	api.PutStateHandleError(t, api.STATE_PAUSED)
	api.PutStateHandleError(t, api.STATE_RUNNING)
	state := api.GetStateHandleError(t)
	if state != api.STATE_RUNNING {
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

	if len(originalBody) == len(newBody) {
		t.Error("PUT RUNNING did not start message generation")
	}

	if !strings.HasPrefix(newBody, originalBody) {
		t.Error("PUT RUNNING modified the original messages")
	}

	// reset state
	api.PutStateHandleError(t, api.CMD_INIT)
}

/*
func TestShutdownState(t *testing.T) {

	api.PutStateHandleError(t, api.CMD_SHUTDOWN)
	_, err := api.GetResponse("http://server/messages")
	if err == nil {
		t.Error("Services still running after shutdown")
	}
}
*/
