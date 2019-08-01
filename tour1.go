package main

import (
    "fmt"
    //"time"
    //"math/rand"
    "math/cmplx"
)
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	Zeta complex128 = cmplx.Sqrt(-5 + 12i)
)
var python, java bool

func main() {
    //rand.Seed(32) //not give random every time
    //fmt.Printf("%d number is random\n", rand.Intn(100))
    //r := rand.New(rand.NewSource(2))
    //fmt.Println("Hello, time is", time.Now(),'\n')
    //fmt.Printf("%f number is random\n", r.Float32())
    fmt.Println("add", add(2,3))   
    b, a := swap("hello", "world")
    fmt.Println("swap", b, a)
    fmt.Println(split(21))
    var c int
    var d float32
    x, y, z := 22, 33.1234567890123456789, "razum"
    fmt.Println("vars", c, d, python, java, x, y, z)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Zeta, Zeta)
}

func add(a int, b int)  int{
    return a+b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum float32) (x float32, y float32) {
    x = sum * 2.5 / 9.2
    y = sum - x
    return
}
