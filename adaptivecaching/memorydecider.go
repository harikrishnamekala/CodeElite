package main

import (
	"fmt"
	"time"

	"github.com/capnm/sysinfo"
)

func main() {
	for {
		SI := sysinfo.Get()
		fmt.Println(SI)
		time.Sleep(time.Second * 5)
	}
}
