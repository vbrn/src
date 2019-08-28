package main

import (
	"fmt"
	"time"
)

//TO SHOW SELECT IN CHANNELS LIKE A SWITCH

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)
    c4 := make(chan string)
    c5 := make(chan string)

    go func() {for {c1 <- "from 1"; time.Sleep(time.Second * 1)}}()
    go func() {for {c1 <- "from 2"; time.Sleep(time.Second * 2)}}()
    go func() {for {c1 <- "from 3"; time.Sleep(time.Second * 3)}}()
    go func() {for {c1 <- "from 4"; time.Sleep(time.Second * 4)}}()
    go func() {for {c1 <- "from 5"; time.Sleep(time.Second * 5)}}()

    go func() { for {
            select {
            case msg := <- c1: fmt.Println(msg)
            case msg := <- c2: fmt.Println(msg)
            case msg := <- c3: fmt.Println(msg)
            case msg := <- c4: fmt.Println(msg)
            case msg := <- c5: fmt.Println(msg)
            }}}()

    var input string
    fmt.Scanln(&input)
}
