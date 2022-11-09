package main

import "fmt"

func zeroValue(iValue int) {
	iValue = 0

}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroValue(i)
	fmt.Println("zeroValue:", i)

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i)

	fmt.Println("pointer:", &i)

	/*
		zeroValue doesn’t change the i in main, but zeroPointer does because it has a reference to the memory address for that variable.
	*/
}
