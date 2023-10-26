package main 

import (
	"fmt"
	"sync"
)


// Unbuffered channel:
// - A channel that needs a receiver as soon as a message is emitted to the channel.
// - It's capacity is 0
// - We do not declare any capacity, hence cannot store any data.

// Buffered channel:
// - Has some capacity to hold data
// - Sending to the channel blocks the goroutine, only if the buffer is full.
// - Receiving from a channel blocks only when the channel is empty.
// - It's length is always <= capacity of the buffered channel.


func send(ch chan string, wg *sync.WaitGroup){
	defer wg.Done()
	ch <- "First message"
	ch <- "Second message"
	ch <- "Third message"
	fmt.Println("All messages sent.")
	go receive(ch, wg)
}

func receive(ch chan string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Message received: ", <- ch)
	fmt.Println("Message received: ", <- ch)
	fmt.Println("Message received: ", <- ch)
	fmt.Println("All messages received.")
}

func main() {
	var wg sync.WaitGroup
	// Creating a buffer channel
	ch := make(chan string, 3) // where 3 => capacity of the buffer
	wg.Add(2)
	go send(ch, &wg)
	wg.Wait()
	fmt.Println("All operations completed.")

}
