package main

/*
Because it bares extra repeating 🫢:

In Go, an array is a numbered sequence of elements of a specific length.
In typical Go code, slices are much more common; arrays are useful in some special scenarios.
*/

import "fmt"

func main() {
	var a [5]int // array of length 5, all type int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5} // declare and init in one-line (thus not zero-valued)
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	fmt.Println(twoD)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
