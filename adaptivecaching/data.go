package main

import (
	"fmt"
)

func bigBytes() *[]byte {
        s := make([]byte, 100000000)
        return &s
}



func main() {
	fmt.Println("Hello, playground")
	var wg sync.WaitGroup
	var mem runtime.MemStats
        runtime.ReadMemStats(&mem)
        fmt.Println(mem.Alloc)
        fmt.Println(mem.TotalAlloc)
        fmt.Println(mem.HeapAlloc)
        fmt.Println(mem.HeapSys)

        for i := 0; i < 10; i++ {
                s := bigBytes()
                if s == nil {
                        fmt.Println("oh noes")
                }
        }

        runtime.ReadMemStats(&mem)
        fmt.Println(mem.Alloc)
        fmt.Println(mem.TotalAlloc)
        fmt.Println(mem.HeapAlloc)
        fmt.Println(mem.HeapSys)
}
