package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	_ "os/exec"
	"strings"
)

var nopath_error = errors.New("Path isn't Provided")

const ShellToUse = "bash"

// executing commands
func exec(input string) error {
	new_input := strings.TrimSuffix(input, "\n")
	args := strings.Split(new_input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return nopath_error
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(ShellToUse, args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// for readding the commands continously
	for {
		fmt.Println("-> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = exec(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
