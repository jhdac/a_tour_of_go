package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	fib := 3
	fib_minus_one := 2
	fib_minus_two := 1
	return func() int {
		if count > 2 {
			fib = fib_minus_one + fib_minus_two
			fib_minus_two = fib_minus_one
			fib_minus_one = fib
			return fib
		}
		if count == 2 {
			count++
			return 2
		}
		count++
		return 1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}