package main

import (
	"fmt"
	"syscall"
)

func main() {
	var rusage *syscall.Rusage

	if err := syscall.Getrusage(1, rusage); err != nil {
		panic(err)
	}
	fmt.Println(rusage)
}

