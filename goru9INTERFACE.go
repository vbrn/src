package main

import (
	"fmt"
)
//SIMPLE METHOD
type Person struct {name string; age int}
func (self *Person) talk() {
	fmt.Println("Simple method")
	if self.age > 3 {fmt.Println("Hello, my name is ", self.name, " I am ", self.age)} else {fmt.Println("I am to young to speak sorry, come in few years")}}


//VARIANT 1
func distance(x, y float64) (float64) {
	return y-x
}
func square_area( x, y float64) (float64) {
	square_storona := distance(x, y)
	return square_storona * square_storona
}
func circle_area(x, y float64) (float64) {
	r :=  distance(x, y)
	return 3.141592653589793 * r*r
}

//VARIANT 2
type Circle struct {x float64; y float64}

func circle_area2(c *Circle) (float64) {
	r :=  distance(c.x, c.y)
	return 3.141592653589793 * r*r
}


//VARIANT 3
type Circle3 struct {x float64; y float64}
func (self *Circle3) distance() (float64) {
	return self.y-self.x}
func (self *Circle3) area() float64 {
	return 3.141592653589793 * self.distance()*self.distance()}
func (self *Circle3) perimetr() float64 {
	return self.distance() * (2 * 3.141592653589793)}//self.distance() = radius kruga
func (self *Circle3) scale(f float64) {
	self.x *= f
	self.y *= f}

type Square3 struct {x float64; y float64}
func (self *Square3) distance() (float64) {
	return self.y - self.x}
func (self *Square3) area() (float64) {
	square_storona := self.distance()
	return square_storona * square_storona}
func (self *Square3) perimetr() float64 {
	return 4 * self.distance()}//self.distance() = storona kvadrata
func (self *Square3) scale(f float64) {
	self.x *= f
	self.y *= f}

//INTERFACE to calculate total area of square and Circle
type Total interface {
	area() float64
	perimetr() float64
	scale(f float64)}
func TotalArea(items ...Total) float64 {
	var total_area float64
	for _, item := range items {
		total_area += item.area()}
	return total_area}
func TotalPerimetr(items ...Total) float64 {
	var total_perimetr float64
	for _, item := range items {
		total_perimetr += item.perimetr()}
	return total_perimetr}
func TotalScale(f float64, items ...Total) {
	for _, item := range items {
		item.scale(f)
	}
}

//EMPTY INTERFACE can get any type
func Print_it(item ...interface{}) {
	fmt.Printf("\n %T type, %v value", item, item)
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//EMPTY INTERFACE AND SWITCHES
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


//FMT STRING() interface
//fmt package has type Stringer interface { String() string} that can be used following wayn
type People struct {name string; poll string; age int}
func (self *People) String() string {
	return fmt.Sprintf("This person is %v %v %v", self.name, self.age, self.poll)

}

func main() {
	//SIMPLE METHOD
	var p = Person{"Bob", 13}
	p.talk()

	//VAR1
	x := distance(5,18)
	fmt.Println("\nDistance m 5 i 18", x)
	fmt.Println("square area", square_area(5, 10))
	fmt.Println("circle area", circle_area(5, 10))
	
	//VAR2
	c := Circle{5, 10}
	fmt.Println("\ncircle2 area", circle_area2(&c))
	
	//VAR3
	c3 := Circle3{5, 10}//Initiliazation of Kruga
	fmt.Println("\ncircle3 area", c3.area())
	fmt.Println("circle3 perimetr", c3.perimetr())
	s3 := Square3{5, 10}//Initiliazation of Kvadrata
	fmt.Println("square3 area", s3.area())
	fmt.Println("square3 perimetr", s3.perimetr())
	//s3.scale(3)
	//fmt.Println("square3 area after scale", s3.area())
	//fmt.Println("square3 perimetr after scale", s3.perimetr())
	
	
	//INTERFACE
	fmt.Println("TotalArea of square and Circle is ", c3.x, TotalArea(&c3, &s3))
	fmt.Println("TotalPerimetr of square and Circle is ", TotalPerimetr(&c3, &s3))
	TotalScale(3, &c3, &s3)
	fmt.Println("TotalArea of square and Circle after 3d scale",  TotalArea(&c3, &s3))
	fmt.Println("TotalPerimetr of square and Circle after scale ", TotalPerimetr(&c3, &s3))
	
	fmt.Println("\nEMPTY INTERFACE")
	Print_it(2)
	Print_it(2.32)
	Print_it("THIS IS STRING")
	fmt.Println("\n")
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
	
	fmt.Println("\nEmtpy Interface and Switches")
	do(21)
	do("hello")
	do(true)

	//STRINGER BUILD IN TO fmt interface
	fmt.Println("\n Now stringer interface to fmt comes")
	var p1 = People{"Bob", "m", 23}
	var p2 = People{"Lana", "z", 19}
	fmt.Println(&p1, &p2)

}
