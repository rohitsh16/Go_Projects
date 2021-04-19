package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered and Unbuffered Channels
	// Unbuffered Channel
	done := make(chan bool)

	go func() {
		fmt.Println("First")
		time.sleep(2*time.Second)
		done <- true
	}()
	fmt.Println("Second")
	result := <-done
	time.sleep(2*time.Second)

}
