package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(3)) // 3 * fact(3-1) => 3 * 2 * fact(2-1) => 3 * 2 * 1 * fact(1-1) => 3 * 2 * 1 * 1 = 3 * 2 * 1 = 3 * 2 = 6

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
