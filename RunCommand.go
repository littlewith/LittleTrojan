package main

import (
	"os/exec"
)

func RunCommand(commandName string, params ...string) string {
	cmd := exec.Command("cmd", "/C", commandName)
	output, err := cmd.CombinedOutput()
	isRunSuccess := chkErrorNoExit(err)
	if isRunSuccess {
		res := covert2Utf8(output)
		return res
	} else {
		return err.Error()
	}
}
