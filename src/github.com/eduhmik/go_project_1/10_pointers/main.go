package main

import "fmt"

func main()  {

	// pointer allows you to point to a memory address of a value
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read val from address
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println("new value:", a)
}