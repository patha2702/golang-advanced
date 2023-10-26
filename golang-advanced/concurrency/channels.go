package main

import (
	"fmt"
	"sync"
)


// Channels:
// - In golang, channels are a means through which different go routines communicate.
// - Do not communicate by sharing memory, instead share memory by communicating.
// - Note: memory in the form of data.
// - The communication is bi-directional by default, that means you can send and receive
//   values through same channel.
// - By default, channels send and receive until the other side is ready
// - This allows go routines to synchronize without explicit locks or condition variables.
// - Channels are passed by reference.

// Each channel can hold a particular data type


// Creating a new channel:

// var c chan string 
// where,
// chan => keyword for channel
// string => data type that a channel can hold


// Alternative way using make() function:

// ch := make(chan string) // For declaring and initializing a channel.


// Channel operations:

// 1. Sending a value:
// ch <- value //Format: channel variable <- value to be send(must be of type that channel can hold)

// 2. Receiving a value:
// v := <- ch

// 3. Closing a channel:
// close(ch) // close => built-in function, ch => argument must be of channel.

// 4. Querying buffer/ capacity of a channel
// cap(ch)

// 5.  To find the current total elements residing in the channel.
// len(ch)


// sends data to the channel
func sell(ch chan string, wg *sync.WaitGroup){
	ch <- "Hello, want to buy a pen?" // Execution of this go routine get blocked here, till there
	// is a go routine accepting this data.
	fmt.Println("Sent data to the channel")
	defer wg.Done()
}

// receives data from the channel
func buy(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("..Waiting for the data")
	value := <- ch // Execution of the go routine gets blocked here, till there is a go routine
	// sending the data, so that data can be accepted here.
	fmt.Printf("Received data: %s", value)
	defer wg.Done()
}


func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
	fmt.Println("Communication successful")

}
