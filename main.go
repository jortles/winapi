package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	user32DLL			= windows.NewLazyDLL("user32.dll")
	procSystemParamInfo	= user32DLL.NewProc("SystemParametersInfoW")
	procCreateFile		= user32DLL.NewProc("CreateFileW")
)

func main() {
	imagePath, _ := windows.UTF16PtrFromString(`C:\Users\hurtadoanthon\Pictures\image.jpg`)
	f, _ := windows.UTF16PtrFromString(`C:\Users\hurtadoanthon\Desktop\test.txt`)
	fmt.Println("[+] Changing now...")
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePath)), 0x001A)
}