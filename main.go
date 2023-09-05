package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	// Se bloquea por lo que es una variable que estan siendo accedida por otras goroutines
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
	// Se desloquea para que las demas goroutines puedan acceder a la variable
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}
func main() {
	var wg sync.WaitGroup
	// var lock sync.Mutex
	var lock sync.RWMutex
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
		// fmt.Println(Balance())
	}
	wg.Wait()
	fmt.Println("El balance es : ", Balance(&lock))
}

// saber si hay condicion de carrera con comando
// go build --race
