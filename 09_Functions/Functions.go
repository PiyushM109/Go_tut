package main

import "fmt"

func main() {
	fmt.Println("Function tutorial")
	fmt.Println(sum(9, 8))
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
