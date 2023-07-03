package tools

import (
	"os/exec"
	"strings"

	"github.com/crystal/groot/global"
	"github.com/crystal/groot/pool"
)

func Dosubfinder(jobParam pool.JobParam) {

	global.G_LOG.Info("doing Dosubfinder")
	dir := global.HOME + "/" + jobParam.Project + "/" + "subfinder-domains.txt"
	domainStr := strings.Join(jobParam.Domains, ",")
	global.Execute(&exec.Cmd{
		Path: "/Users/crystal/academy/tools/bin/subfinder",
		Args: []string{"subfinder", "-d", domainStr, "-o", dir},
	})

}
