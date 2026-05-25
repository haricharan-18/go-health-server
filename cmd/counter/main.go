package main

import (
	"fmt"
	"github.com/Zartex-the-art/sei-ratelimiter/pkg/counter"
	"sync"
)

func main() {
	c := counter.NewSafeCounter()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", c.Value())
}