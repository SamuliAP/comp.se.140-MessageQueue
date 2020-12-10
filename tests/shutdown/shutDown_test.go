package shutdown

import (
	"../../api"
	"testing"
)

func TestShutdownState(t *testing.T) {

	_, err := api.GetState()
	if err != nil {
		t.Error("Service already shut down")
		return
	}

	api.PutStateHandleError(t, api.CMD_SHUTDOWN)
	_, messagesErr := api.GetResponse("http://server/messages")
	_, stateErr := api.GetResponse("http://server/state")
	if messagesErr == nil || stateErr == nil {
		t.Error("Services still running after shutdown")
	}
}
