package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("Hello", c)

	msg := <-c
	fmt.Println(msg)
}

func count(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- s
		time.Sleep(time.Second)
	}
}
