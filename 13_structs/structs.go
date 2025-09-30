package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println("Structs")

	fmt.Println(person{"Piyush", 23})

	fmt.Println(person{name: "Piyu", age: 24})

	fmt.Println(newPerson("Krishna"))

	s := person{name: "Virat", age: 18}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rolex",
		true,
	}

	fmt.Println(dog)
}
