package main

import (
	"fmt"
	"sync"
)

// Go routine leak:
// - Whenever you launch a go routine function, you must make sure it exits eventually.
// - A go routine that would never terminate, forever occupies the memory it has reserved.
// - This is called as go routine leak.
// - Go routines leak if they end up either blocked forever on I/O like channel communication
//   or fall into infinite loops.


func leak(wg *sync.WaitGroup) {
	ch := make(chan string, 5)
	go func(){
		val := <- ch
		fmt.Println("Received value:", val)
		wg.Done()
	}()
	fmt.Println("Leak method exiting")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}
