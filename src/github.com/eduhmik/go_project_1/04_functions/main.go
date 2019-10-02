package main

import "fmt"

func greetings(name string) string{

	return "Hello " + name
}

func getAreaOfaRect(num1 int, num2 int) int {
	return num1 * num2
}

func main()  {
	fmt.Println(greetings("Edwin"))
	fmt.Println(getAreaOfaRect(5, 3))
}