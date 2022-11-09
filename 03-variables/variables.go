package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued, thus output will be 0
	var e int
	fmt.Println(e)

	// with := syntax shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)
}
