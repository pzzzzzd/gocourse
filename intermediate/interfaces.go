package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	wigth, height float64
}

type circle struct {
	radius float64
}

// type rect1 struct {
// 	wigth, height float64
// }

func (r rect) area() float64 {
	return r.height * r.wigth
}
func (r rect) perim() float64 {
	return 2 * (r.height + r.wigth)
}

// func (r rect1) area() float64 {
// 	return r.height * r.wigth
// }

func (c circle) area() float64 {
	return math.Pi * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect{wigth: 3, height: 4}
	c := circle{radius: 5}
	// r1 := rect1{wigth: 5, height: 7}

	measure(r)
	measure(c)
	// measure(r1)

	myPrinter(1, "Jhon", 4.5654, false)

	printType(345)
	printType(" ")
	printType(false)

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	case bool:
		fmt.Println("Type: Bool")
	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

// func (r rect1) perim() float64 {
// 	return 2 * (r.height + r.wigth)
// }
