package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello from", name)
		time.Sleep(500 * time.Second)
	}
}

func main() {
	go hello("goroutine-1")
	go hello("goroutine-2")

	time.Sleep(3 * time.Second)
}
