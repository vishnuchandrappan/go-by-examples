package main

import "fmt"

func main() {

	m := make(map[string]int)

	fmt.Println("map => ", m)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map => ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	a, prs := m["k2"]
	fmt.Println("prs:", a, prs)

	delete(m, "k2")
	fmt.Println("map:", m)

	a2, prs2 := m["k2"]
	fmt.Println("prs:", a2, prs2)

	a3, prs3 := m["non-existing"]
	fmt.Println("test non existing", a3, prs3)

	n := map[string]string{"foo": "test", "bar": "test2"}
	fmt.Println("map:", n)
}
