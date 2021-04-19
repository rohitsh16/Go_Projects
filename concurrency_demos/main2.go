package main

import (
	"fmt"
	"time"
	__"sync"
)

//	 An unbuffered channel is a channel that needs a receiver as soon as a message is emitted to the channel.
//	If the capacity is zero or absent, the channel is unbuffered and communication succeeds
//	 only when both a sender and receiver are ready.


func main() {
	go count("First")	// Runs this in background and continue to next line immediately i.e. goroutine
	count("Second")

	// if go routine finishes then program terminates


}
 
func count (s string) {
	for i := 1; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Second)
	}
}