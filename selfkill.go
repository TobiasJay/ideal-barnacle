package main

import (
	"fmt"
	"syscall"
)

func main() {
	handle, err := syscall.GetCurrentProcess()

	if err != nil {
		fmt.Println(err)
	}
	err = syscall.TerminateProcess(handle, 1)
	fmt.Printf("Process: %d died unsuccessfully\n", handle)
	// We should not see process 1 die unsuccessfully
	// that would mean our process didn't terminate
}
