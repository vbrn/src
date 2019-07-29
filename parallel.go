package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(3)
	go bar()
	go crab()
	go foo()
	wg.Wait()
}

func foo() {
	for i := 0; i < 4; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 10; i < 14; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(4 * time.Millisecond)
	}
	wg.Done()
}

func crab() {
	for i := 15; i < 20; i++ {
		fmt.Println("Crab:", i)
		time.Sleep(6 * time.Millisecond)
	}
	wg.Done()
}
