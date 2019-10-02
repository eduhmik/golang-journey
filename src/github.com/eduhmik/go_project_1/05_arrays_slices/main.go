package main

import "fmt"

func main()  {
	// var fruitArr [2]string


	// arrays
	fruitArr := [2]string {"Apple", "Orange"}

	fmt.Println(fruitArr)


	// slices
	fruitSlice := []int {1, 2, 3, 5}

	// use len() to display count of items in an array
	/*choose the range of items to display in an array using [1:2]
	this will display items at index 1 and will not include the item in index 2 */

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice[1:]))
	fmt.Println(fruitSlice[1:])
}