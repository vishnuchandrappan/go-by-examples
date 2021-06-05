package main

import "fmt"

func getName() (string, string) {
	return "John", "Doe"
}

func main() {

	firstName, secondName := getName()
	fmt.Printf("My name is %s %s\n", firstName, secondName)

}
