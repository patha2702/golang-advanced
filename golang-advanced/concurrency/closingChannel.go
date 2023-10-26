package main

import (
	"fmt"
)


// Closing a channel:
// - Meaning no more data can be sent to that channel
// - It's generally done when there is no more data to be sent.
// - Using close(channel), in-built function.

// Checking whether a channel is closed or not:
// - value, ok := <-ch
// - if ok => true, then ch channel is open, else closed.


// When you close a channel and there are still values in the channel's buffer,
// those values can be received by the remaining receivers.
// However, once the channel is closed and all values have been received,
// further attempts to receive from the channel will return the zero value for
// the channel's element type and false for the second return value (indicating that there are no more values to receive).


func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 10
	close(ch)// Channel closed meaning cannot send data, but if data is present 
	// in the channel we can receive it.
	val, ok := <- ch
	fmt.Println(val, ok)
	val, ok = <- ch
	fmt.Println(val, ok)
	val, ok = <- ch
	fmt.Println(val, ok)
}
