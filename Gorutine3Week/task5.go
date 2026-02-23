package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	mu      sync.RWMutex
	balance int
}

func NewAccount(initialBalance int) *Account {
	return &Account{balance: initialBalance}
}

func (a *Account) Balance() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func main() {
	acc := NewAccount(1000)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				bal := acc.Balance()
				fmt.Printf("Читатель %d: баланс = %d\n", id, bal)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			acc.Deposit(500)
			fmt.Printf("Писатель: пополнение, новый баланс = %d\n", acc.Balance())
			time.Sleep(50 * time.Millisecond)

			if err := acc.Withdraw(200); err != nil {
				fmt.Printf("Ошибка снятия: %v\n", err)
			} else {
				fmt.Printf("Писатель: снятие, новый баланс = %d\n", acc.Balance())
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Printf("Итоговый баланс: %d\n", acc.Balance())
}
