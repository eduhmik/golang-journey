package main

import "fmt"

func main()  {

	// using vars
	var name  = "edwin"

	var age int32 = 18
	var isCool = false
	isCool = true

	// shorthand assignment that must be done within a function
	number := +254718433329

	fmt.Println(name, age, isCool, number)
	fmt.Printf("%T\n", isCool)
}