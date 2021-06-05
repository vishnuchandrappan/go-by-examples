package main

import "fmt"

func zeroval(value int) {
	value = 0
}

func zeroptr(value *int) {
	*value = 0
}

func main() {

	i := 1
	fmt.Println("initial =>", i)

	zeroval(i)
	fmt.Println("zeroval =>", i)

	zeroptr(&i)
	fmt.Println("zeroptr =>", i)

	fmt.Println("pointer =>", &i)

}
