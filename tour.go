package main

import (
    "fmt"
    //"time"
    //"math/rand"
)

func main() {
    //rand.Seed(32) //not give random every time
    //fmt.Printf("%d number is random\n", rand.Intn(100))
    //r := rand.New(rand.NewSource(2))
    //fmt.Println("Hello, time is", time.Now(),'\n')
    //fmt.Printf("%f number is random\n", r.Float32())
    go foo()
    go bar()
    
    
}

func foo() {

    for i := 0; i < 5; i++ {
        fmt.Println("Foo", i)
    }
}
func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("Bar:", i)
	}
}

