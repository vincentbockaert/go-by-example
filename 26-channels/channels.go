package main

import "fmt"

// channels are the pipes that connect concurrent goroutines.
// you can send values into channels from one goroutine and receive those values into another goroutine

func main() {
	messages := make(chan string) // channels are typed by the values they convey across

	go func() {
		messages <- "ping" // send value into channel, syntax: "channel <- value"
	}()

	msg := <-messages // receive value from the channel, syntax: "<-channel"
	// we don't need to add a timer to wait for the goroutines to finish, by default sends and receives are blocking until both are ready
	fmt.Println(msg)

}
