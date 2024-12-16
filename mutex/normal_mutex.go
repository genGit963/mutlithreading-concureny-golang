package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sharedResources int = 100
	lock                = sync.Mutex{}
)

// add to sharedResources
func stingy() {
	// executing 1000 different threads of stingy
	lock.Lock()
	for i := 0; i < 1000; i++ {
		sharedResources = sharedResources + 10
		time.Sleep(1 * time.Millisecond)

	}
	lock.Unlock()
	fmt.Println("done stingy !!")
}

// remove from sharedResources
func spendy() {
	// executing 1000 different threads of spendy
	lock.Lock()
	for i := 0; i < 1000; i++ {
		sharedResources = sharedResources - 10
		time.Sleep(1 * time.Millisecond)
	}
	lock.Unlock()
	fmt.Println("done spendy !!")
}

func main() {

	go stingy()
	go spendy()

	// sleep for 5 second
	time.Sleep(3 * time.Second)
	// then result
	fmt.Println(sharedResources, "done !")
}
