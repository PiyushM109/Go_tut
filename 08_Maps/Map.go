package main

import "fmt"

func main() {
	fmt.Println("Maps")

	students := make(map[string]int)

	students["Piyush"] = 2
	students["Raj"] = 2
	students["Swapnil"] = 2

	fmt.Println(students)

	fmt.Println(students["Krish"])

	fmt.Println(len(students))
	delete(students, "Raj")

	fmt.Println(students)

	val, ispres := students["Raj"]

	fmt.Println("val", val, "Present", ispres)

}
