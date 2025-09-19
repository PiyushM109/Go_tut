package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b", b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("c", c)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
