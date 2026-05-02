package main

import (
	"fmt"
	"unsafe"
)

func getPEB() uintptr

func main() {
	// get the peb address directly from the cpu
	peb := getPEB()
	fmt.Printf("[*] peb address: 0x%X\n", peb)

	// locate loader data (Ldr) at offset 0x18
	// this structure tracks all dlls loaded into the process
	ldr := *(*uintptr)(unsafe.Pointer(peb + 0x18))

	// access InMemoryOrderModuleList at offset 0x20
	// this is the head of a doubly linked list of modules
	listHead := ldr + 0x20
	
	// walk the list
	// entry 1: The .exe itself
	// entry 2: ntdll.dll
	// entry 3: kernel32.dll
	currEntry := *(*uintptr)(unsafe.Pointer(listHead)) 
	nextEntry := *(*uintptr)(unsafe.Pointer(currEntry)) 
	kernel32Entry := *(*uintptr)(unsafe.Pointer(nextEntry))

	// extract BaseAddress at offset 0x20 of the entry
	// in the InMemoryOrder list the BaseAddress is at 0x20
	kernel32Base := *(*uintptr)(unsafe.Pointer(kernel32Entry + 0x20))

	fmt.Printf("[+] kernel32.dll base address: 0x%X\n", kernel32Base)
}