package main

import "fmt"

func main() {
	i := 0
	for i <= 7 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
		i++
	}
}
