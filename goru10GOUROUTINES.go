package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 1; i < 100; i++ {
		amt := time.Duration(rand.Intn(250))
		fmt.Println(n, " : ", i, " - ",amt)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
for i:= 1; i < 100; i++ {
	go f(i)
	}
	
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
