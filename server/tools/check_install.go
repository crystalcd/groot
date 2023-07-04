package tools

import (
	"os/exec"
	"os/user"
	"runtime"

	"github.com/crystal/groot/global"
)

var MAC = "darwin"

func GoCheck() bool {
	cmd := exec.Command("go", "-version")
	err := cmd.Run()
	if (err != nil) {
		global.G_LOG.Warn("[!] Go is not installed, installing now")
		return false
	}
	global.G_LOG.Warn("[+] Go is installed")
	return true
}

func InstallGo() {
	osVersion := runtime.GOOS
	global.G_LOG.Info(osVersion)
	if (MAC == osVersion) {
		global.G_LOG.Info("Mac")
		cmd := exec.Command("brew", "install", "go")
		err :=cmd.Run()
		if (err != nil) {
			global.G_LOG.Warn("-----", err)
		}
	}
}

func CheckSubfinder() {
	
}


func getHomeDir() string{
	homeDir, err := user.Current()
	if err != nil {
		global.G_LOG.Error("无法获取Home目录:", err)
	}
	return homeDir.HomeDir
}