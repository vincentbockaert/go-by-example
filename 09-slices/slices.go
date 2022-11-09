package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "hondje"
	s[1] = "komt"
	s[2] = "zo"

	fmt.Println("set:", s)
	fmt.Println("get string at index", 2, ":", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "dadelijk") // may return a new slice value
	s = append(s, "wel", "af")
	fmt.Println("apd:", s)

	c := make([]string, len(s))

	copy(c, s) // copy into c from s
	fmt.Println("cpy:", c)

	l := s[2:5] // from 2 up to but excluding 5, slice[low:high]
	fmt.Println("sl1:", l)

	l = s[:5] // same as s[0:5]
	fmt.Println("sl2:", l)

	l = s[0:5]
	fmt.Println("sl2:", l)

	l = s[2:] // start from 2 (inclusive) and go to the end
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // declare and init a slice in a single line
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) // length of inner slices can vary in multi-dimensional data structure unlike with arrays, but the outer is set by the declared size
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
