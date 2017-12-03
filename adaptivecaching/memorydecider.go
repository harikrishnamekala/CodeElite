package adaptivecaching

import (
	"fmt"
	"syscall"
)

func Getusage() {
	var rusage *syscall.Rusage

	if err := syscall.Getrusage(2, rusage); err != nil {
		panic(err)
	}
	fmt.Println(rusage)
}ls
