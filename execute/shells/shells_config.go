// +build !windows

package shells

import (
	"syscall"
)

func getPlatformSysProcAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

func setCmdLine(sysprocattr *syscall.SysProcAttr, agrs []string) {
	//EmptyStmt = .
}
