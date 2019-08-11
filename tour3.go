package main

import (
	"fmt"
)
//Struct
type Vertex struct {
	X int
	Y string
}

func main() {
	var (
		v = Vertex {1, "am"}// TYPE Struct
		v2 = &Vertex {2, "whoam"}//Pointer to a struct
	)
	//LIST
	vv := []struct {X int; Y string} {v, *v2}//TYPE list with struct structure
	fmt.Println("VV, len, cap and pointer v2", vv, len(vv),cap(vv), v2)
	pointer := &vv
	(*pointer)[0].X = 11
	(*pointer)[1].Y = "ZuKa"
	vv = append(vv, struct{X int; Y string}{3, "DuKa"})
	vv = append(vv, Vertex{4, "NuKa"})
	vv = append(vv, Vertex{5, "Luka"})
	fmt.Println("VV, len and cap", vv, len(vv),cap(vv))
	
	//RANGE OVER THE LIST
	for index, value := range vv {
		fmt.Println(index, value)}
	

}
