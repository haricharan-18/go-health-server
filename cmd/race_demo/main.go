package main

import (
	"fmt"
	"sync"
	"time"
)

// Demo 1: BankAccount with race condition
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	// RACE VERSION: comment out the next 3 lines to see the race
	a.mu.Lock()
	a.balance += amount
	a.mu.Unlock()

	// UNCOMMENT THIS TO SEE RACE:
	// a.balance += amount
}

func (a *BankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	fmt.Println("=== Demo 1: Bank Account ===")
	fmt.Println("Initial balance:", account.Balance())

	// 100 goroutines each depositing 10
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(id int) {
			defer wg.Done()
			account.Deposit(10)
			fmt.Printf("Goroutine %d done. Balance: %d\n", id, account.Balance())
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d (expected: 2000)\n", account.Balance())

	if account.Balance() == 2000 {
		fmt.Println("✅ Correct — mutex is working")
	} else {
		fmt.Println("❌ Race condition! Balance is wrong")
	}

	// Demo 2: Read-Modify-Write race
	fmt.Println("\n=== Demo 2: Read-Modify-Write ===")
	var counter int
	var wg2 sync.WaitGroup

	wg2.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg2.Done()
			// This is 3 operations: READ, ADD, WRITE
			current := counter
			time.Sleep(time.Microsecond) // Force context switch
			counter = current + 1
		}()
	}

	wg2.Wait()
	fmt.Printf("Counter: %d (expected: 1000)\n", counter)

	// Demo 3: Slice append race
	fmt.Println("\n=== Demo 3: Slice Append Race ===")
	var sharedSlice []int
	var wg3 sync.WaitGroup

	wg3.Add(2)
	go func() {
		defer wg3.Done()
		for i := 0; i < 100; i++ {
			sharedSlice = append(sharedSlice, i)
		}
	}()

	go func() {
		defer wg3.Done()
		for i := 0; i < 100; i++ {
			sharedSlice = append(sharedSlice, i+1000)
		}
	}()

	wg3.Wait()
	fmt.Printf("Slice length: %d (expected: 200)\n", len(sharedSlice))
}