package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	// edrs love to watch this specific call
	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	// this is a giant signature for scanners
	getModuleHandle := kernel32.NewProc("GetModuleHandleW")

	// this triggers the runtime hooks I talked about
	handle, _, _ := getModuleHandle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("kernel32.dll"))))

	if handle == 0 {
		fmt.Println("[-] failed to get handle.")
		return
	}

	fmt.Printf("[!] kernel32.dll found at 0x%X\n", handle)
}