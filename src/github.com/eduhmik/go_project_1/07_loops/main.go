package main

import "fmt"

func main() {

	// for loop
	for i := 1; i <= 10; i++ {
		// fmt.Println(i)
	}

	// FizzBuzz
	for k := 1; k <= 100; k++ {
		if k % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if k % 3 == 0 {
			fmt.Println("Fizz")
		} else if k % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(k)
		}
	}
	// calling the fibnonacci function inside the main func

	fib(10)

}

// Fibonacci sequence

func fib(n int) {
    for i, j := 0, 1; j < n; i, j = i+j,i {
        fmt.Println(i)
    }
}