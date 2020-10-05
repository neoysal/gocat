package shells

import (
	"strings"
	"syscall"
)

func getPlatformSysProcAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: true}
}

func setCmdLine(sysprocattr *syscall.SysProcAttr, agrs []string) {
	sysprocattr.CmdLine = strings.Join(agrs, " ")
}
