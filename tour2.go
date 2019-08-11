package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
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
    //FOR and DEFER
    for i := 0; i <=5; i++ {
    defer fmt.Printf("%T %v\n", i, i)
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
    // FOR and RANGE
    // _range_ iterates over elements in a variety of data
    // structures. Let's see how to use `range` with some
    // of the data structures we've already learned.

    // Here we use `range` to sum the numbers in a slice.
    // Arrays work like this too.
    nums := []int{2, 3, 4}
    summ := 0
    for _, num := range nums {
        summ += num
    }
    fmt.Println("sum:", summ)

    // `range` on arrays and slices provides both the
    // index and value for each entry. Above we didn't
    // need the index, so we ignored it with the
    // blank identifier `_`. Sometimes we actually want
    // the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` on strings iterates over Unicode code
    // points. The first value is the starting byte index
    // of the `rune` and the second the `rune` itself.
    for i, c := range "goСУРЖИК" {
        fmt.Println(i, c)
    }


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
	
	
today := time.Now().Weekday()
fmt.Println("\nWhen is monday?")	
switch time.Saturday {
case today%7 + 0:
		fmt.Println("Today.")
	case today%7 + 1:
		fmt.Println("Tomorrow.")
	case today%7 + 2:
		fmt.Println("In two days.")	
	case today%7 + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}	
		t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
	// DEFER
		fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)}
	fmt.Println("done")
			
	}
	
	
