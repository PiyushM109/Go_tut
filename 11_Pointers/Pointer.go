package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println("Pointers")

	val := 1

	fmt.Println(val)

	zeroval(val)
	fmt.Println(val)

	zeroptr(&val)
	fmt.Println(val)
}
