package main

import (
	"fmt"
)


// Panic:
// - It is a runtime exception.
// - It means an unexpected condition has occured, due to
//   which the execution of the program has terminated.


// Few scenarios for panic situations:
// 1. Sending to a channel after it has been closed
// 2. closing an already closed channel.

func main() {
	ch := make(chan string, 5)
	ch <- "Hello"
	ch <- "World"
	close(ch)
	close(ch) // closing the already closed channel
	ch <- "Error" // Sending to channel, which is closed.

	fmt.Println(<- ch)
}
