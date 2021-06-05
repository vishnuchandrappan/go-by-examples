package main

import "fmt"

func main() {

	var friends [5]int
	fmt.Println("emp:", friends)

	friends[4] = 100
	fmt.Println("set:", friends)
	fmt.Println("get:", friends[4])

	fmt.Println("len:", len(friends))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
