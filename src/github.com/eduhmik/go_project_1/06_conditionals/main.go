package main

import "fmt"

func main()  {

	x := 6
	y := 7

	if x % y == 0 {
		fmt.Printf("%d is a divisor of %d\n", y, x)
	} else {
		fmt.Printf("%d is not a divisor of %d\n", y, x)
	}
}