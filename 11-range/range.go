package main

import "fmt"

/*
Because it bares repeating from: https://gobyexample.com/range

range iterates over elements in a variety of data structures.
Let’s see how to use range with some of the data structures we’ve already learned.
*/

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// 'range' on arrays and slices provides both the index and value for each entry, hence the blank identifier to assign the index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	kvpairs := map[string]string{"a": "apple", "b": "banana"}

	// 'range' on a map can iterate over key/value pairs
	for k, v := range kvpairs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// or 'range' can just iterate of only the keys
	for k := range kvpairs {
		fmt.Println("key:", k)
	}

	// 'range' on strings iterates over Unicode code points
	// the first value is the starting byte index of the rune
	// the second value is the rune itself
	// https://gobyexample.com/strings-and-runes
	// A Go string is a read-only slice of bytes.
	// The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
	// In other languages, strings are made of “characters”.
	//  In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
	for index, runeValue := range "go" {
		fmt.Println(index, runeValue)
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
