package main

import (
	"fmt"
	"sync"
)

var (
	// variable made to act like a bank account
	balance int = 100
)

// Deposit function registers an input to the balance variable
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	// The WaitGroup's method to indicate when the all the routines has finished
	defer wg.Done()
	// it blocks, so it is a variable that it being accessing by the others goroutines
	lock.Lock()
	// I created a copy from balance's value
	b := balance
	// I added the value of amount param to the copy and I assigned to the original balance variable
	balance = b + amount
	// then I unlock writes from the other goroutines to the balance variable thus it avoids the race condition
	lock.Unlock()
}

// Balance function is used to read the balance's value
func Balance(lock *sync.RWMutex) int {
	// as the previous functionality there I used a lock but lock's type is Read
	// so, I lock all the reads from the other goroutines
	lock.RLock()
	// Again I copy the balance's value to return
	b := balance
	// I unlock the read and finally I return the value
	lock.RUnlock()
	return b
}
func main() {
	// WaitGroup variable that is used to control the goroutines' flows
	var wg sync.WaitGroup
	// var lock sync.Mutex used to lock and unlock reads and writes
	var lock sync.RWMutex
	//  I simulated 5 inputs to the balance
	for i := 0; i <= 5; i++ {
		// It is required add an amount of goroutines to wait for them
		wg.Add(1)
		// I throw the Deposit function in a goroutine
		go Deposit(i*100, &wg, &lock)
	}
	// finally, I wait all the goroutines come to an end
	wg.Wait()
	// then I show the final balance's value
	fmt.Println("El balance es : ", Balance(&lock))
}

// command to know if it has race condition
// go build --race
