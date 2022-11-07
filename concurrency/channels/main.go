package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//ch <- 353 // send

	// Example 1
	fmt.Println("Example 1")
	fmt.Println("_________")

	go func() {
		ch <- 353
	}()

	val := <-ch // receive
	fmt.Printf("got %d\n", val)

	// Example 2
	fmt.Println(" ")
	fmt.Println("Example 2")
	fmt.Println("_________")

	const count = 3

	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	// Example 3
	fmt.Println(" ")
	fmt.Println("Example 3")
	fmt.Println("_________")

	/*
		By closing the Go routine you can now range through the channels and return the values
	*/

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	// This is a buffered channel
	// The buffered value allows you to specify how many values you can send through that channel without blocking

	// Create channel with buffer value of 1
	ch2 := make(chan int, 2)
	// Send value 19 to ch2
	ch2 <- 19
	// Send value 20 to ch2
	ch2 <- 20
	// receive the ch2 value of 19 and assign it to val2
	val2 := <-ch2
	// receive the ch2 value of 20 and assign it to val3
	val3 := <-ch2
	// print value to terminal
	fmt.Println(val2, val3)
}
