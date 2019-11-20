package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	fmt.Println("Start")
	runtime.GOMAXPROCS(1)
	x := 0
	go func() {
		for {
			x++
			runtime.Gosched()//указание явно на переключение контекста
		}
	}()
	time.Sleep(500*time.Millisecond)
	fmt.Println(x)
}
