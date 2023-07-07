package tools

import (
	"os/exec"
	"os/user"
	"runtime"

	"github.com/crystal/groot/logging"
)

var MAC = "darwin"

func GoCheck() bool {
	cmd := exec.Command("go", "-version")
	err := cmd.Run()
	if err != nil {
		logging.RuntimeLog.Warn("[!] Go is not installed, installing now")
		return false
	}
	logging.RuntimeLog.Warn("[+] Go is installed")
	return true
}

func InstallGo() {
	osVersion := runtime.GOOS
	logging.RuntimeLog.Info(osVersion)
	if MAC == osVersion {
		logging.RuntimeLog.Info("Mac")
		cmd := exec.Command("brew", "install", "go")
		err := cmd.Run()
		if err != nil {
			logging.RuntimeLog.Warn("-----", err)
		}
	}
}

func CheckSubfinder() {

}

func getHomeDir() string {
	homeDir, err := user.Current()
	if err != nil {
		logging.RuntimeLog.Error("无法获取Home目录:", err)
	}
	return homeDir.HomeDir
}
