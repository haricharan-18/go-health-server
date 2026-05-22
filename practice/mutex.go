package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("Final count:", count)
}
