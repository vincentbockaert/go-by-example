package main

import "fmt"

type rectangle struct {
	width, height int
}

// note how we have the "(struct)" in front of the func-name, not the case for regular "funcs", thus clearly a method
func (r *rectangle) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver. --> I presume pointer is usually more performanant (less memory copying?)
func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectangle{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
