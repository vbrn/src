package main

import (
    "fmt"
    "math"
    "runtime"
    
)





   //IF WITH A STATEMENT
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {return v} else {fmt.Printf("%g >= %g\n", v, lim)
	}	// can't use v here, though
	    return lim}
	    
func Sqrt(x float64) float64 {
	for z := float64(int(x/2)); z > 0; z-- {
 		formula := (z*z - x)/(2*z)
 			if formula == 0 {
 				return z } } 
		return 0.0 }







func main() {
    //FOR
    for i := 0; i <=5; i++ {
    fmt.Printf("%T %v\n", i, i)
    }
    //WHILE but FOR
    sum := 1
	for sum < 100 {
	    fmt.Println(sum)   
		sum += sum	}
	fmt.Println("FINAL", sum)   
    //WHILE forever
    sum = 55
    for {
        if sum < 0 { break}
        fmt.Println(sum)
        sum = sum - 5 }
        
    //GOTO
    var x3 = 15
    LOOP: if x3> 0 {
    fmt.Println("goto loop", x3)
    x3--
    goto LOOP }
    
    //IF WITH A STATEMENT
 	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	x1 := 3
	y1 := 2
	if x1 == 3 && y1 > 0 { fmt.Println("even")} 
    fmt.Println("selfsqrt", Sqrt(9))

//SWITCH
fmt.Println("Го ранз он")
switch os := runtime.GOOS; os {
case "darwin": fmt.Println("OS X")
case "linux": fmt.Printf("Linux %s", os)
default: //freebsd openbsd plan9 windows
    fmt.Printf("%s, \n", os) }
	}
