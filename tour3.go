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
	
	//Map
	m := make(map[Vertex]int)
	m[Vertex{11, "mapYuka"}] = 111
	m[v] = 18
	// ranging over map
	for key, value := range m {
		fmt.Println(key, key.X*11, value * 33)
	}
	fmt.Println("before delete map", m)
	delete(m, v)
	fmt.Println("after delete map", m)
	
	var countryCapitalMap map[string]string
   /* create a map*/
   countryCapitalMap = make(map[string]string)
   
   /* insert key-value pairs in the map*/
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"

   /* print map using keys*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }

   /* test if entry is present in the map or not*/
   capital, ok := countryCapitalMap["United States"]

   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      fmt.Println("Capital of United States is", capital)  
   } else {
      fmt.Println("Capital of United States is not present") 
   }
}
