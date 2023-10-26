package main

import (
	"fmt"
	"sync"
)


// Mutexes:
// - Mutex(Mutual Exclusion) is a synchronization primitive provided by the standard library to protect shared data from being
//   accessed concurrently by multiple goroutines.
// - Mutexes allow us to lock access to shared data. This ensures that we can control which go routines can access data
//   at which time.

// Read-Write mutexes (sync.RWmutex):
// - It has extra functionality than mutexes.
// - With Lock(), Unlock() it also has:
//   - RLock()
//   - RUnlock()
// - It is safe for multiple threads to read from the shared resource,
// - while it is dangerous for multiple threads to perform read - write operation.
// - To improve the performance of application, by granting many threads to read, we can use RWmutex.
// - 1 writer at one instance of time.
// - Multiple readers at one instance of time.


type safeData struct {
	mux sync.Mutex
	data map[string]int
}


func (d *safeData) insert(key string, val int) {
	d.mux.Lock()
	d.data[key] = val
	d.mux.Unlock()
}


func (d *safeData) read(key string) int {
	d.mux.Lock()
	val := d.data[key]
	d.mux.Unlock()
	return val
}


func main() {
	var wg sync.WaitGroup
	var marks safeData
	marks.data = make(map[string]int)
	marks.insert("raj", 23)
	marks.insert("man", 25)
	wg.Add(1)
	go func(obj *safeData, s *sync.WaitGroup) {
		defer s.Done()
		obj.insert("ram", 43)
		val := obj.read("raej")
		fmt.Println(val)
	}(&marks, &wg)
	wg.Wait()
	fmt.Println("Successful completion.")
}
