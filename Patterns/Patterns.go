package main

import "fmt"

func RectanglePattern(row int, col int) {
	for range row {
		for range col {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func RightAngledStarPattern(row int) {
	for i := 1; i <= row; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func RightAngledNumberPattern(row int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func RightAngledNumberPattern2(row int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func InvertedRightAngle(row int) {
	for i := range row {
		for j := row - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func InvertedRightAngleNumber(row int) {
	for i := 0; i < row; i++ {
		for j := 1; j <= row-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {
	row := 5
	// col := 5

	// // Rectangle star pattern
	// RectanglePattern(row, col)

	//Right angled Star Triangle Pattern
	// RightAngledStarPattern(row)

	//Right Angled Number Pattern
	// RightAngledNumberPattern(row)
	// RightAngledNumberPattern2(row)

	//inverted Right Angle
	// InvertedRightAngle(row)
	InvertedRightAngleNumber(row)
}
