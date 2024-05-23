package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

func setDLLDirectory() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setDllDirectoryProc := kernel32.NewProc("SetDllDirectoryW")

	dirUTF16, err := syscall.UTF16PtrFromString(exeDir)
	if err != nil {
		return err
	}
	r1, _, err := setDllDirectoryProc.Call(uintptr(unsafe.Pointer(dirUTF16)))
	if r1 == 0 {
		return fmt.Errorf("failed to set DLL directory: %v", err)
	}
	return nil
