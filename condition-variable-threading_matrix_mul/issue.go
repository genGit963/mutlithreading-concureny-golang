package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("stingy updated to : ", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendy() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		money -= 20
		fmt.Println("spendy reduced to : ", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("spendy Done")
}

func main() {

	go stingy()
	go spendy()

	time.Sleep(3 * time.Second)
	fmt.Println("Last: ", money)
}
