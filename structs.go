package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func otherPerson(name string) person {
	p := person{name: name}
	p.age = 42
	return p
}

func main() {
	fmt.Println(person{"Jane", 20})
	fmt.Println(person{name: "Alice", age: 31})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	person1 := newPerson("John")
	fmt.Println("newPerson => ", person1)

	person1.age = 32
	fmt.Println("newPerson 2 => ", person1)

	person2 := otherPerson("Jane")
	fmt.Println("otherPerson => ", person2)

	person2.name = "Patrick Jane"
	fmt.Println("otherPerson2 =>", person2)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
