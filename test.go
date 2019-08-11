package main

import (
	"fmt"
	)
type Vertigo struct {a int; b string}

func main() {
	a := []struct{a int; b string} {{1, "a"}, {2, "b"}}
	fmt.Printf("len %d,   cap %d  values %v\n",  len(a), cap(a), a)
	var s = Vertigo {3, "c"}
	a = append(a, s)
	fmt.Printf("len %d,   cap %d  values %v\n",  len(a), cap(a), a)
	
	a = append(a, struct{a int; b string}{4, "f"})
	a = append(a, struct{a int; b string}{12, "dsd"})
	
	fmt.Printf("len %d,   cap %d  values %v\n",  len(a), cap(a), a)
	}
	
