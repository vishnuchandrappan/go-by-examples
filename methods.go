package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func main() {
	fmt.Println("fmt")

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
