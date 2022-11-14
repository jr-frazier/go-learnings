package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func noConcurrency() {
	loopA()
	loopB()
}

func withConcurrency() {
	wg.Add(2)
	go loopA()
	go loopB()
	wg.Wait()
}

func loopA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", 1)
	}
}

func loopB() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", 1)
	}
}

func main() {
	// noConcurrency()
	withConcurrency()
}
