package main

import (
	"fmt"
	"sync"
	"time"
)

// CV: condition variable

var (
	moneyBank = 100
	mutexLock = sync.Mutex{}

	// condition variable sync
	moneyDeposited = sync.NewCond(&mutexLock)
)

func stingyWithCV() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		moneyBank += 10
		fmt.Println("stingy updated to : ", money)

		// broadcast
		moneyDeposited.Broadcast()

		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendyWithCV() {
	for i := 0; i < 1000; i++ {

		// 1. mutex lock
		lock.Lock()
		// 2. wait-broadcast-signal
		for moneyBank-20 < 0 {
			// wait for 20$ in money
			moneyDeposited.Wait()
		}
		moneyBank -= 20
		fmt.Println("spendy reduced to : ", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("spendy Done")
}

func main() {

	go stingyWithCV()
	go spendyWithCV()

	time.Sleep(3 * time.Second)
	fmt.Println("Last: ", money)
}
