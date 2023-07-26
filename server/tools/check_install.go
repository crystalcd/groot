package tools

import (
	"os/exec"
	"os/user"
	"runtime"
)

var MAC = "darwin"

func GoCheck() bool {
	cmd := exec.Command("go", "-version")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func InstallGo() {
	osVersion := runtime.GOOS
	if MAC == osVersion {
		cmd := exec.Command("brew", "install", "go")
		err := cmd.Run()
		if err != nil {
		}
	}
}

func CheckSubfinder() {

}

func getHomeDir() string {
	homeDir, err := user.Current()
	if err != nil {
	}
	return homeDir.HomeDir
}
