package main

import (
	"fmt"
	"sei-ratelimiter/pkg/counter"
)

func main() {
	result := counter.ConcurrentCounter(1000)
	fmt.Println("Final count:", result)
}