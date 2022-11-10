package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds base, looks like a field without a name
type container struct {
	base // seems similar to using super in java
	str  string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)   // same as co.num but makes it more readable, clearer I suppose :)
	fmt.Println("describe:", co.describe()) // note how this is actually coming from "base"

	// note how the below interface is already implemented in the base-struct
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
