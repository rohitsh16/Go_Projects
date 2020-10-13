package main

import (
	"fmt"

	_ "github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("Port Scanner")
	port_open := port.scanport("tcp", "localhost", 8080)
	fmt.Println("Port Open : %t\n", port_open)
}
