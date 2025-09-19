package main

import "fmt"

func main() {
	fmt.Println("Slices")

	a := make([]int, 4)

	fmt.Println("a", a, "length", len(a), "cap", cap(a))

	b := make([]int, 0, 8)

	fmt.Println("a", b, "length", len(b), "cap", cap(b))

	a[0] = 1

	a = append(a, 1, 2, 3)

	fmt.Println(a)

	b = make([]int, len(a))
	copy(b, a)
	fmt.Println(b)
	fmt.Println(a)

	l := a[2:5]
	fmt.Println("sl:", l)

	l1 := a[:5]
	fmt.Println("sl1:", l1)
	
	l2 := a[2:]
	fmt.Println("sl2:", l2)
}
