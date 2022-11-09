package main

import "fmt"

/*
Because it bares repeating from: https://gobyexample.com/maps

Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

To create an empty map, use the builtin make: make(map[key-type]val-type).

Set key/value pairs using typical name[key] = val syntax.

Use the builtin "delete(map,key)" to remove kv-pairs from a map
*/

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	disambiguate := m["k2"] // will return 0 --> yet we know the keys exists
	fmt.Println("dsm:", disambiguate)

	// assign the value to to the blank identifier _,
	// the second return value when getting a value from a map indicates whether or not the key was present in the map,
	// assign this to 'prs'
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
