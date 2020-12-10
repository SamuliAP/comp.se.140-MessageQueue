package tests

import (
	"../api"
	"strings"
	"testing"
)

func testRunLog(t *testing.T) {

	handleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{api.STATE_PAUSED})
	handleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{api.STATE_PAUSED})
	handleRunlogNewStateCheck(t, api.STATE_RUNNING, []string{api.STATE_RUNNING})
	handleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
	handleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
	handleRunlogNewStateCheck(t, api.STATE_PAUSED, []string{api.STATE_PAUSED})
	handleRunlogNewStateCheck(t, api.CMD_INIT, []string{api.CMD_INIT, api.STATE_RUNNING})
}

func handleRunlogNewStateCheck(t *testing.T, state string, contains []string) {
	runLogSize := len(api.GetRunLogHandleError(t))

	api.PutStateHandleError(t, state)
	runLog := api.GetRunLogHandleError(t)

	newLog := runLog[runLogSize:]
	for _, c := range contains {
		if strings.Contains(newLog, c) {
			t.Error("Run log did not record state:", state)
		}
	}
}
