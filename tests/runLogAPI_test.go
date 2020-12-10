package tests

import (
	"../api"
	"strings"
	"testing"
)

func TestRunLog(t *testing.T) {

	api.PutStateHandleError(t, api.STATE_RUNNING)
	HandleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{api.STATE_PAUSED})
	HandleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{})
	HandleRunlogNewStateCheck(t, api.STATE_RUNNING, []string{api.STATE_RUNNING})
	HandleRunlogNewStateCheck(t, api.STATE_RUNNING, []string{})
	HandleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
	HandleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
	HandleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{api.STATE_PAUSED})
	HandleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
}

func HandleRunlogNewStateCheck(t *testing.T, state string, contains []string) {
	runLogSize := len(api.GetRunLogHandleError(t))

	api.PutStateHandleError(t, state)
	runLog := api.GetRunLogHandleError(t)

	newLog := runLog[runLogSize:]
	t.Log(newLog)
	if len(contains) == 0 && len(newLog) != 0 {
		t.Error("Run log logged duplicate event", state)
	}
	for _, c := range contains {
		if strings.Contains(newLog, c) {
			t.Error("Run log did not record state:", state)
		}
	}
}
