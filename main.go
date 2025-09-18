// package main

// import (
// 	"fmt"
// 	"math"
// )

// func add(a int, b int) int {
// 	return a + b
// }

// func swap(x string, y string) (string, string) {
// 	return y, x
// }

// func split(num int) (x, y int) {
// 	x = num / 2
// 	y = num * 2
// 	return
// }

// func main() {
// 	fmt.Println("Go tour", math.Pi)
// 	fmt.Println(add(4, 5))
// 	first, last := swap("Piyush", "More")
// 	fmt.Println(first, last)
// 	x, y := split(8)
// 	fmt.Println(x, y)
// }

package main

import "fmt"

func main() {
	i, j := 22, 12
	fmt.Println(i, j)
	p := &i
	fmt.Println(p, *p, &p)
}
