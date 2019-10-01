package main

import "fmt"

func main(){
	fmt.Println("Area of a Circle")
	// Area of a circle formula
	// pi*r*r
	// Define pie
	var pi float64 = 3.142
	var r int = 7
	var Area = pi * float64(r) * float64(r)

	fmt.Println(Area)
}