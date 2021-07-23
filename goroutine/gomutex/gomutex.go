package gomutex

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	Mutex *sync.Mutex
}

func (a * Account) WithDraw(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= value
	fmt.Printf("Balance withdraw successful: %v\n", value)
	a.Mutex.Unlock()
	wg.Done()
}

func (a *Account) Deposit(value int, wg *sync.WaitGroup)  {
	a.Mutex.Lock()
	a.Balance += value
	fmt.Printf("Balance add successful: %v\n", value)
	a.Mutex.Unlock()
	wg.Done()
}