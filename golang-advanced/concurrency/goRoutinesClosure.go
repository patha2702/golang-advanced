package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=1; i<=10; i++ {
		go func() {
			fmt.Println(i)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Successful completion")
}
