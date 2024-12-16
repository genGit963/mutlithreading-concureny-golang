package main

import (
	"fmt"
	"time"
)

var sharedResources int = 100

// add to sharedResources
func stingy() {
	sharedResources = sharedResources + 10
	time.Sleep(1 * time.Millisecond)

	fmt.Println("done stingy !!")
}

// remove from sharedResources
func spendy() {
	sharedResources = sharedResources + 10
	time.Sleep(1 * time.Millisecond)
	fmt.Println("done spendy !!")
}

func main() {
	// executing 100 different threads of stingy and spendy
	for i := 0; i < 100; i++ {
		go stingy()
		go spendy()
	}
	fmt.Println("done !")
}
