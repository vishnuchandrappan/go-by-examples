package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	fmt.Println("nums:", nums)
	sum := 0

	for index, num := range nums {
		fmt.Println(index)
		sum += num
	}

	fmt.Println("sum:", sum)

	days := map[string]string{"MON": "Monday", "TUE": "Tuesday"}
	for key, value := range days {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for k := range days {
		fmt.Println("key:", k)
	}

	// (index, unicodeVal)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
