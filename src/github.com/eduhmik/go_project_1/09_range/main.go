package main

import "fmt"

func main()  {
	ids := []int{1, 2, 3, 4, 5, 6}

	// loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	// use an underscore in the place of the index holder

	for  _, idw := range ids {
		fmt.Printf("ID: %d\n", idw)
	}

	// finding sum of all items in an array

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is:", sum)

	// range with maps

	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	for k,v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}
}