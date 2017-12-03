package adaptivecaching

import (
	"fmt"
	"sync"
	"time"
)

func Sleeper(dec chan bool, wg sync.WaitGroup) {
	fmt.Println("Sleeper")
	time.Sleep(time.Millisecond * 10)
	wg.Done()
	dec <- true
}

func Wakers(dec chan bool, wg sync.WaitGroup) {
	fmt.Println("Waker")
	time.Sleep(time.Second * 2)
	wg.Done()
	dec <- false
}

func Quit(qui chan int) {
	qui <- 1
}

func main() {
	fmt.Println("Hello, playground")
	var wg sync.WaitGroup
	c := make(chan bool)
	q := make(chan int)
	ch := make(chan bool)
	go func() {
		for {
			select {
			case deci := <-c:
				if deci == true {
					fmt.Println("Executed Sleeper")
				}
			case mega := <-ch:
				if mega == false {
					fmt.Println("Executed Waker")
				}
			case quutting := <-q:
				fmt.Print(quutting)
				return

			}
		}
	}()
	for i := 0; i < 10; i++ {
		c <- true
		go Sleeper(c, wg)
		wg.Add(1)
		ch <- false
		go Wakers(ch, wg)
		wg.Add(1)

	}
	Quit(q)
	wg.Wait()
}
