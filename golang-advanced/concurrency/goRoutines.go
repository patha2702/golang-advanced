package main

import (
	"fmt"
	"sync"
)


// Go Routine:
// - Light weight thread that has separate independent execution.
// - Can execute concurrently with other go routines.
// - All go routines are managed by Go runtime.
// - They are not deterministic, since all are running concurrently.

// Starting a Goroutine:
// - To start a new goroutine, use the 'go' keyword followed by a function call, e.g., 'go function()'.
// - The function specified in the 'go' statement is executed independently by a new goroutine.
// - The main program, including the instructions inside 'main()', runs in the main goroutine.
// - Calling a function with the 'go' keyword leads to the creation of a new goroutine dedicated to executing that function.
// - Each goroutine, whether it's the main goroutine or a new one, has its own stack.



func calculate(num int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(num * num)
}


func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go calculate(i, &wg)
	}

	wg.Wait()
	fmt.Println("All action completed.")
}
