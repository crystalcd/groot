package global

import "os/exec"

func Execute(inputcmd *exec.Cmd) {
	cmd := inputcmd
	// cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		G_LOG.Info("Command execution failed: ", err)
		return
	}
	G_LOG.Info(string(output))

	G_LOG.Info("done Dosubfinder")
}
