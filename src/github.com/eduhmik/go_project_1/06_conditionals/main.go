package main

import (
	"fmt"
)

func main() {

	x := 6
	y := 7

	// if else

	if x%y == 0 {
		fmt.Printf("%d is a divisor of %d\n", y, x)
	} else {
		fmt.Printf("%d is not a divisor of %d\n", y, x)
	}

	// else if

	color := "red"

	if color == "red" {
		fmt.Println("There is danger")
	} else if color == "green" {
		fmt.Println("Go on")
	} else {
		fmt.Println("Invalid color")
	}

	// switch

	switch color {
	case "red":
		fmt.Println("There is danger. Stop!!!")
	case "green":
		fmt.Println("Drive on")
	default:
		fmt.Println("Does not recognize the value")
	}

	// do while

	var z = 6

	for {
		z--
		if z == 0 {
			break
		}
		fmt.Println("Counting...", z)
		if z == 1 {
			fmt.Println("Fireeeee!!!!")
		}
	}
}
