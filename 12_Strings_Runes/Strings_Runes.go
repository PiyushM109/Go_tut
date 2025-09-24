package main

import "fmt"

func main() {
	fmt.Println("Strings and Runes")

	str := "Piyush"
	fmt.Println(str)
	fmt.Println("length", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	fmt.Println()

	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
