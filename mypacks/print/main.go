package print

import "fmt"

var Foo string
var Too int

//simple printing
func Print(a ...interface{}) {
	fmt.Println(a...)
	}


func init() {
	Foo = "Zuka"
	}

//SHOULD BE make 
//go install 
//in the folder mypack


