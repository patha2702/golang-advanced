package main

import (
	"fmt"
	"sync"
)


// Iterating over the channel using range:
// - When using range with channels, it iterates overs the values 
// - sent on the channel until the channel is closed.

// - You can get error, if you are iterating over the channel 
// - that is not closed.


func send(ch chan int, wg *sync.WaitGroup) {
	for i:=0; i<10; i++ {
		ch <- i
	}
	close(ch)
	defer wg.Done()
}

func receive(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println("Received value: ", val)
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	go send(ch, &wg)
	go receive(ch, &wg)
	wg.Wait()
	fmt.Println("Successful communication.")
}
