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

	putStateErr := api.PutState(api.CMD_SHUTDOWN)
	_, messagesErr := api.GetResponse("http://server/messages")
	_, stateErr := api.GetResponse("http://server/state")
	if putStateErr == nil || messagesErr == nil || stateErr == nil {
		t.Error("Services still running after shutdown")
	}
}
