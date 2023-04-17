package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func abort(funcname string, err int) {
	panic(funcname + " failed: " + syscall.Errno(err).Error())
}

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	user32, _          = syscall.LoadLibrary("user32.dll")
	messageBox, _      = syscall.GetProcAddress(user32, "MessageBoxW")
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300
)

func MessageBox(caption, text string, style uintptr) (result int) {
	// var hwnd HWND
	ret, _, callErr := syscall.Syscall6(uintptr(messageBox), 4,
		0, // HWND
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),    // Text
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))), // Caption
		style, // type
		0,
		0)
	if callErr != 0 {
		abort("Call MessageBox", int(callErr))
	}
	result = int(ret)
	return
}

func init() {
	fmt.Print("Starting Up\n")
}
