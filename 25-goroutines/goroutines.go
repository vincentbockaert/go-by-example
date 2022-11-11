package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// :) wait to finish ... if you don't believe this is running async, comment out the time.Sleep code and run it again.
	time.Sleep(time.Second)
	fmt.Println("done")
}
