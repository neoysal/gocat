package privdetect

import (
	"syscall"
	"unsafe"
)

func Privlevel() string {
	var token syscall.Token
	token, err := syscall.OpenCurrentProcessToken()
	if err != nil {
		return "User"
	}
	//token := Token(^uintptr(4 - 1))
	var isElevated uint32
	var outLen uint32
	syscall.GetTokenInformation(token, syscall.TokenElevation, (*byte)(unsafe.Pointer(&isElevated)), uint32(unsafe.Sizeof(isElevated)), &outLen)
	if isElevated != 0 {
		return "Elevated"
	} else {
		return "User"
	}
}
