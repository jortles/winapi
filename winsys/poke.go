package winsys

import (
	"fmt"
	"github.com/getlantern/errors"
	"unsafe"
)

func OpenProcessHandle(p *Poke) error {
	var rights uint32 = PROCESS_CREATE_THREAD |
		PROCESS_QUERY_INFORMATION |
		PROCESS_VM_OPERATION |
		PROCESS_VM_WRITE |
		PROCESS_VM_READ

	var inheritHandle uint32 = 0
	var processID uint32 = p.Pid
	remoteProcHandle, _, lastErr := ProcOpenProcess.Call(
		uintptr(rights),
		uintptr(inheritHandle),
		uintptr(processID))
	if remoteProcHandle == 0 {
		return errors.Wrap(lastErr)
	}
	p.RemoteProcHandle = remoteProcHandle
	fmt.Printf("[-] Input PID: %v\n", p.Pid)
	fmt.Printf("[-] Input DLL %v\n", p.DllPath)
	fmt.Printf("[-] Process handle: %v\n", unsafe.Pointer(p.RemoteProcHandle))
	return nil
}
