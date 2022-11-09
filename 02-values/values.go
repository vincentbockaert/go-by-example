package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // golang, string concat

	fmt.Println("1+1 =", 1+1) // similar to string concat but with a space/for values

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// outputs
	// golang
	// 1+1 = 2
	// 7.0/3.0 = 2.3333333333333335
	// false
	// true
	// false
}
