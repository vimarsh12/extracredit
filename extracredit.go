package main

import (
	"fmt"
	"sync"
)

var (
	accountBalance int        // Shared bank account balance
	mu             sync.Mutex // Mutex for synchronizing access to the account balance
	once           sync.Once  // Used to initialize the account balance once
)

// initializeAccount initializing the bank account with a starting balance.
func initializeAccount(initialBalance int) {
	once.Do(func() {
		accountBalance = initialBalance
		fmt.Println("Account initialized with balance:", accountBalance)
	})
}

// deposit is simulating depositing money into the account.
func deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	accountBalance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, accountBalance)
}

// withdraw simulates withdrawing money from the account.
func withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	if accountBalance >= amount {
		accountBalance -= amount
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, accountBalance)
	} else {
		fmt.Printf("Failed to withdraw %d, insufficient balance. Current balance: %d\n", amount, accountBalance)
	}
}

func main() {
	var wg sync.WaitGroup

	// Initializing the account
	initializeAccount(1000)

	// Simulating customers
	customerActions := []func(int, *sync.WaitGroup){
		deposit,
		withdraw,
		deposit,
		withdraw,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go customerActions[i%len(customerActions)](100, &wg) // Alternate deposit and withdraw
	}

	wg.Wait()
	fmt.Println("Final balance:", accountBalance)
}
