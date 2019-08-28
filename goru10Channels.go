package main

import (
	"fmt"
	"time"
	"strings"
)

func pinger(c chan <- string) {for i:=0; ;i++ {c <- fmt.Sprintf("ping %v", i)}}
func ponger(c chan <- string, b <- chan string) {for i:=0; ;i++ {if strings.Contains(<-b, "0") {fmt.Println("Zero Ping", <-b)}; c <- fmt.Sprintf("pong %v", i)}}

/*func printer (c <- chan string) {for {msg := <-c; fmt.Println(msg); if strings.Contains(msg, "ping") {time.Sleep(time.Second*1)} else {time.Sleep(time.Second*6)}}}*/


func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c, c)
	//go printer(c)
	go func(){for {fmt.Println(<-c); time.Sleep(time.Second*1)}}()
	var input string
	fmt.Scanln(&input)
}
